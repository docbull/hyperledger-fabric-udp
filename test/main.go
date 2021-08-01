package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"runtime"
	"strings"

	protoG "github.com/golang/protobuf/proto"
	proto "github.com/hyperledger/fabric-protos-go/gossip"
	quic "github.com/lucas-clemente/quic-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	udp "github.com/docbull/inlab-fabric-udp-proto"
)

type loggingWriter struct{ io.Writer }

// Message means received block data
type Message struct {
	Block           *udp.Envelope
	PeerContainerIP string
	writer          *loggingWriter
}

// --------------------------Peer to MPBTP (gRPC)------------------------

// SendBlock2Peer forwards the received block that from another
// MPBTP container to connected Peer container.
func (msg *Message) SendBlock2Peer() {
	peerIP := msg.PeerContainerIP + ":16220"
	conn, err := grpc.Dial(peerIP, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	mpbtpClient := udp.NewUDPServiceClient(conn)
	res, err := mpbtpClient.BlockDataForUDP(context.Background(), msg.Block)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.Code != udp.StatusCode_Ok {
		log.Println("Not OK for MPBTP block transmission:", res.Code)
		return
	} else {
		log.Println("REceived message from MPBTP:", res)
	}
}

// Peer2UDP waits connections from the peer container for
// forwarding block data to other MPBTP containers over UDP.
// This function is working based on gRPC.
func (msg *Message) Peer2UDP() {
	lis, err := net.Listen("tcp", ":11800")
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	fmt.Println("waiting Peer connection...")

	grpcServer := grpc.NewServer()
	udp.RegisterUDPServiceServer(grpcServer, msg)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

// BlockDataForUDP forwards the block data over UDP protocols.
// It stores the block that received through gRPC call, then
// forward the block data to UDP transmission functions.
func (msg *Message) BlockDataForUDP(ctx context.Context, envelope *udp.Envelope) (*udp.Status, error) {
	log.Println("Receive Block data from the Peer container")

	msg.Block.Payload = envelope.Payload
	msg.Block.Signature = envelope.Signature
	msg.Block.SecretEnvelope = envelope.SecretEnvelope

	go msg.UDPBlockSender()

	return &udp.Status{Code: udp.StatusCode_Ok}, nil
}

// -------------------------MPBTP to MPBTP (QUIC)--------------------------

func (msg *Message) UDPServerListen() {
	listener, err := quic.ListenAddr(":8000", generateTLSConfig(), nil)
	if err != nil {
		fmt.Println(err)
	}
	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			fmt.Println(err)
		} else {
			go msg.handleUDPConnection(sess)
		}
	}
}

func (msg *Message) handleUDPConnection(sess quic.Session) {
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	} else {
		for {
			msg.writer = &loggingWriter{stream}
			_, err = io.Copy(msg, stream)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (msg *Message) Write(buf []byte) (int, error) {
	envelope := &proto.Envelope{}
	err := protoG.Unmarshal(buf, envelope)
	if err != nil {
		log.Println("Unmarshal error:", err)
		return 0, nil
	}
	msg.Block.Payload = envelope.Payload
	msg.Block.Signature = envelope.Signature

	go msg.SendBlock2Peer()

	return msg.writer.Writer.Write(buf)
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}

func (msg *Message) SendResponse(conn *net.UDPConn, addr *net.UDPAddr, length string) {
	_, err := conn.WriteToUDP([]byte(length), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func (msg *Message) UDPBlockSender() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr("203.247.240.234:8000", tlsConf, nil)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Received Block data from the Peer container")

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	marshalledEnvelope, err := protoG.Marshal(msg.Block)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshalled data size:", len(marshalledEnvelope))

	_, err = stream.Write([]byte(marshalledEnvelope))
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, len(marshalledEnvelope))
	_, err = io.ReadFull(stream, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(buf))
}

// ------------------------Peer to MPBTP (TCP)----------------------------

// WaitPeerConnection receives Peer Endpoint
// from the peer container, and then saves it
// for communicating to docker virtual IP.
func (msg *Message) WaitPeerConnection() {
	conn, err := net.Listen("tcp", ":20000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		sock, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
		}

		remoteAddr := sock.RemoteAddr().String()
		peerIP := strings.Split(remoteAddr, ":")
		msg.PeerContainerIP = peerIP[0]

		fmt.Println(msg.PeerContainerIP)
	}
}

func start() {
	msg := &Message{
		Block:           nil,
		PeerContainerIP: "",
		writer:          &loggingWriter{},
	}
	msg.Block = &udp.Envelope{
		Payload:        nil,
		Signature:      nil,
		SecretEnvelope: nil,
	}

	// waiting Peer containers' connection
	go msg.Peer2UDP()
	// waiting MPBTP containers' connection
	go msg.UDPServerListen()

	// waiting peer container's connection
	msg.WaitPeerConnection()
}

func main() {
	runtime.GOMAXPROCS((runtime.NumCPU()))
	fmt.Println("Running CPU cores:", runtime.GOMAXPROCS(0))

	start()
}
