package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	d2d "github.com/docbull/inlab-fabric-modules/inlab-d2d/protos"
	protoG "github.com/golang/protobuf/proto"
	proto "github.com/hyperledger/fabric-protos-go/gossip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Message means received block data
type Message struct {
	Block           *d2d.Envelope
	PeerContainerIP string
}

// D2DSend2Peer sends a block data to peer container
func (msg *Message) D2DSend2Peer() {
	// dial to peer container
	// need to set the IP address
	peer2D2DBlockIP := msg.PeerContainerIP + ":16220"
	conn, err := net.Dial("tcp", peer2D2DBlockIP)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	protoEnvelope := proto.Envelope{
		Payload:   msg.Block.Payload,
		Signature: msg.Block.Signature,
	}

	envelope, err := protoG.Marshal(&protoEnvelope)
	if err != nil {
		log.Println(err)
		return
	}
	n, err := conn.Write(envelope)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("-----------------------------")
	fmt.Println("Size of Block Data:", n)
	msg.SentDataSize += n
	fmt.Println("Total Block Data sizes sent: ", msg.SentDataSize)
	fmt.Println("-----------------------------")

	recvBuf := make([]byte, 64)
	conn.Read(recvBuf)
	resMsg := string(recvBuf)

	fmt.Println("Received message from the Peer:", resMsg)
	if strings.Contains(resMsg, "error") {
		fmt.Println("Re-send the block")

		go msg.D2DSend2Peer()
	}
}

// D2DtoD2D waits D2D connection
func (msg *Message) D2DtoD2D() {
	// For D2D connection, open :12800 port
	conn, err := net.Listen("tcp", ":12800")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("D2D waiting ...")
	grpcServer := grpc.NewServer()
	d2d.RegisterD2DServiceServer(grpcServer, msg)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}

func (msg *Message) D2DBlockStream(ctx context.Context, envelope *d2d.Envelope) (*d2d.Status, error) {
	fmt.Println("D2D Block Stream gRPC")

	msg.Block.Payload = envelope.Payload
	msg.Block.Signature = envelope.Signature
	msg.Block.SecretEnvelope = envelope.SecretEnvelope

	go msg.D2DSend2Peer()

	return &d2d.Status{Code: d2d.StatusCode_Ok}, nil
}

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
		defer sock.Close()

		remoteAddr := sock.RemoteAddr().String()
		peerIP := strings.Split(remoteAddr, ":")
		msg.PeerContainerIP = peerIP[0]

		fmt.Println(msg.PeerContainerIP)
	}
}

// D2DReceiveFromPeer receives data block from peer container and save it
func (msg *Message) D2DReceiveFromPeer(conn net.Conn) {
	envelope := make([]byte, 0)
	buf := make([]byte, 1024)
	length := 0

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		envelope = append(envelope, buf[:n]...)

		length += n
	}

	fmt.Println("size of msg from Peer:", length)

	protoEnvelope := &proto.Envelope{}
	err := protoG.Unmarshal(envelope[:length], protoEnvelope)
	if err != nil {
		log.Println(err)
		conn.Write([]byte("error"))
		return
	}

	msg.Block = &d2d.Envelope{
		Payload:   protoEnvelope.Payload,
		Signature: protoEnvelope.Signature,
	}

	go msg.D2DSend()
}

// D2DSend sends block data to another D2D container
func (msg *Message) D2DSend() {
	pairIP := os.Getenv("D2D_PAIR") + ":12800"

	conn, err := grpc.Dial(pairIP, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(time.Now(), "D2D Send")

	client := d2d.NewD2DServiceClient(conn)

	res, err := client.D2DBlockStream(context.Background(), msg.Block)
	if err != nil {
		log.Println("D2D connection loss")
		return
	}
	if res.Code != d2d.StatusCode_Ok {
		log.Println(res.Code)
		log.Println("It might be not connected with Peer container. Please check that is the peer container running.")
		return
	} else {
		fmt.Println("Block has sent to the D2D container")
	}
	fmt.Println()
}

// Peer2UDP waits Peer container connection
func (msg *Message) Peer2UDP() {
	lis, err := net.Listen("tcp", ":11800")
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	fmt.Println("Peer connection waiting...")

	grpcServer := grpc.NewServer()
	info.RegisterSDNServiceServer(grpcServer, overlay)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func start() {
	msg := Message{
		Block:           nil,
		PeerContainerIP: "",
	}
	msg.Block = &d2d.Envelope{
		Payload:        nil,
		Signature:      nil,
		SecretEnvelope: nil,
	}

	// waiting Peer containers' connection
	go msg.Peer2D2D()

	// waiting D2D containers' connection
	go msg.D2DtoD2D()
	go msg.WaitPeerConnection()
	for {

	}
}

func main() {
	start()
}
