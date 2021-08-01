package mpbtp

import (
	"fmt"
	"log"
	"net"
	"time"

	protoG "github.com/golang/protobuf/proto"
	proto "github.com/hyperledger/fabric-protos-go/gossip"
)

func (msg *Message) UDPBlockSender() {
	conn, err := net.Dial("udp", "203.247.240.234:8000")
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
