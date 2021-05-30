package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

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

func (msg *Message) SendResponse(conn *net.UDPConn, addr *net.UDPAddr, length string) {
	_, err := conn.WriteToUDP([]byte(length), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func (msg *Message) UDPServerListen() {
	envelope := make([]byte, 0)
	buf := make([]byte, 1024*10)

	addr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 8000,
	}
	serv, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		n, remoteaddr, err := serv.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("Some error %v", err)
			continue
		}
		log.Println("Received a Block from", remoteaddr)
		log.Println("Received size of Block data:", n)

		envelope = append(envelope, buf[:n]...)

		protoEnvelope := &proto.Envelope{}
		err = protoG.Unmarshal(envelope, protoEnvelope)
		if err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}
		length := string(n)

		go msg.SendResponse(serv, remoteaddr, length)
	}
}

func (msg *Message) UDPBlockSender() {
	conn, err := net.Dial("udp", "192.168.1.2:8000")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	defer conn.Close()

	envelope, err := protoG.Marshal(msg.Block)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshalled data size:", len(envelope))

	n, err := conn.Write(envelope)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wrote data size:", n)

	/*
		p := make([]byte, 2048)
		_, err = bufio.NewReader(conn).Read(p)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}
		log.Println("Received size of the block:", string(p))
	*/
}

func (msg *Message) BlockDataForUDP(ctx context.Context, envelope *udp.Envelope) (*udp.Status, error) {
	log.Println("Receive Block data from the Peer container")

	msg.Block.Payload = envelope.Payload
	msg.Block.Signature = envelope.Signature
	msg.Block.SecretEnvelope = envelope.SecretEnvelope

	go msg.UDPBlockSender()

	return &udp.Status{Code: udp.StatusCode_Ok}, nil
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
	udp.RegisterUDPServiceServer(grpcServer, msg)
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
	msg.Block = &udp.Envelope{
		Payload:        nil,
		Signature:      nil,
		SecretEnvelope: nil,
	}

	// waiting Peer containers' connection
	go msg.Peer2UDP()
	go msg.UDPServerListen()

	// waiting D2D containers' connection
	go msg.WaitPeerConnection()
	for {

	}
}

func main() {
	start()
}
