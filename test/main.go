package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"time"

	udp "github.com/docbull/inlab-fabric-udp-proto"
	protoG "github.com/golang/protobuf/proto"
	proto "github.com/hyperledger/fabric-protos-go/gossip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Message means received block data
type Message struct {
	Block           *udp.Envelope
	PeerContainerIP string
	Key             []byte
	IV              []byte
}

func (msg *Message) UDPServerListen() {
	addr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 8000,
	}
	serv, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	defer serv.Close()

	for {
		msg.handleUDPConnection(serv)
	}
}

func (msg *Message) handleUDPConnection(serv *net.UDPConn) {
	// buf for receiving RT symbols of block
	buf := make([]byte, 4096*10)

	n, remoteaddr, err := serv.ReadFromUDP(buf)
	fmt.Println("Block Received:", time.Now())
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	buf = buf[0:n]

	log.Println("Received a Block from", remoteaddr)
	log.Println("Received size of the Block data:", n)

	msg.UDPBlockHandler(serv, remoteaddr, n, buf)
}

func (msg *Message) UDPBlockHandler(serv *net.UDPConn, remoteaddr *net.UDPAddr, n int, buf []byte) {
	envelope := &proto.Envelope{}
	err := protoG.Unmarshal(buf, envelope)
	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}
	msg.Block.Payload = envelope.Payload
	msg.Block.Signature = envelope.Signature

	length := string(n)
	go msg.SendResponse(serv, remoteaddr, length)
	go msg.SendBlock2Peer()
}

func (msg *Message) SendResponse(conn *net.UDPConn, addr *net.UDPAddr, length string) {
	_, err := conn.WriteToUDP([]byte(length), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func (msg *Message) SendBlock2Peer() {
	peerIP := msg.PeerContainerIP + ":16220"
	conn, err := grpc.Dial(peerIP)
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

// Peer2UDP listens connections from the peer container for
// forwarding block data. This function is working based on gRPC.
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

	fmt.Println(msg.Block.Payload)
	go msg.UDPBlockSender()

	return &udp.Status{Code: udp.StatusCode_Ok}, nil
}

func (msg *Message) UDPBlockSender() {
	conn, err := net.Dial("udp", "192.168.1.7:8000")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	defer conn.Close()

	marshalledEnvelope, err := protoG.Marshal(msg.Block)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshalled data size:", len(marshalledEnvelope))

	_, err = conn.Write(marshalledEnvelope)
	if err != nil {
		log.Println(err)
	}
}

// WaitPeerConnection receives Peer Endpoint
// from the peer container, and then saves it.
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
		Key:             []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC},
		IV:              []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC},
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