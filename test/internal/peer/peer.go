package peer

import (
	"context"
	"fmt"
	"log"
	"net"

	udp "github.com/docbull/inlab-fabric-udp-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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
