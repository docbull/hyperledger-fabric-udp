package main

import (
	"fmt"
	"runtime"

	hlfudp "github.com/docbull/hyperledger-fabric-udp/inlab-udp/hlfudp"
	udp "github.com/docbull/inlab-fabric-udp-proto"
)

func start() {
	msg := hlfudp.Message{
		Block:           nil,
		PeerContainerIP: "",
		Fountain:        &hlfudp.RaptorCodec{RTSize: 0, RTData: make([]byte, 0)},
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
	go msg.UDPServerListen()

	// waiting D2D containers' connection
	msg.WaitPeerConnection()
}

func main() {
	runtime.GOMAXPROCS((runtime.NumCPU()))
	fmt.Println("Running CPU cores:", runtime.GOMAXPROCS(0))

	start()
}
