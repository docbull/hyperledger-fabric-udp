package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	udp "github.com/docbull/inlab-fabric-udp-proto"
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

func (msg *Message) BlockDataForUDP(ctx context.Context, envelope *udp.Envelope) (*udp.Status, error) {
	log.Println("Receive Block data from the Peer container")

	msg.Block.Payload = envelope.Payload
	msg.Block.Signature = envelope.Signature
	msg.Block.SecretEnvelope = envelope.SecretEnvelope

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

	// waiting D2D containers' connection
	go msg.WaitPeerConnection()
	for {

	}
}

func main() {
	start()
}
