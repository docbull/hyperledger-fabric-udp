package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	fountain "gofountain"
	"log"
	"math"
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

type AttrMsg struct {
	Size int
	Cnt  int
}

type UDPEnvelope struct {
	Envelope []byte
}

var UDPenvelope UDPEnvelope
var attr AttrMsg

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

func (msg *Message) SendBlock2Peer() {
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
	fmt.Println("-----------------------------")

	recvBuf := make([]byte, 64)
	conn.Read(recvBuf)
	resMsg := string(recvBuf)

	fmt.Println("Received message from the Peer:", resMsg)

	UDPenvelope.Envelope = UDPenvelope.Envelope[:0]
}

func (msg *Message) UDPBlockHandler() {
	protoEnvelope := &proto.Envelope{}

	fmt.Println("--------- UDP Block Handler ---------")
	fmt.Println("Block size with padding:", len(UDPenvelope.Envelope))
	UDPenvelope.Envelope = UDPenvelope.Envelope[:attr.Size]
	fmt.Println("Block size without padding:", len(UDPenvelope.Envelope))
	fmt.Println("-------------------------------------")

	err := protoG.Unmarshal(UDPenvelope.Envelope, protoEnvelope)
	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}
	msg.Block.Payload = protoEnvelope.Payload
	msg.Block.Signature = protoEnvelope.Signature

	go msg.SendBlock2Peer()
}

func (msg *Message) SendResponse(conn *net.UDPConn, addr *net.UDPAddr, res string) {
	_, err := conn.WriteToUDP([]byte(res), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func (msg *Message) handleUDPConnection(serv *net.UDPConn) {
	// RT decoding, 64 symbols and 5 extra symbols
	codec := fountain.NewRaptorCodec(128, 7)
	dec := codec.NewDecoder(128)
	var encSymbols []fountain.LTBlock

	// envelope for block unmarshalling
	// buf for receiving RT symbols of block
	buf := make([]byte, 1024*10)

	n, remoteaddr, err := serv.ReadFromUDP(buf)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	buf = buf[0:n]

	log.Println("Received a Block from", remoteaddr)
	log.Println("Received size of the Block data:", n)

	if n < 30 {
		// receiving block size
		err := json.Unmarshal(buf, &attr)
		if err != nil {
			panic(err)
		}

		fmt.Println("Received Block Size:", attr.Size)
	} else {
		// receiving encoding symbol slices
		err = json.Unmarshal(buf, &encSymbols)
		if err != nil {
			panic(err)
		}

		// if success to docode, return the original symbols
		errCheck := decoder(encSymbols, codec, dec)
		if errCheck != nil {
			fmt.Println("Complete recovery!")
			attr.Cnt += len(errCheck)
			fmt.Println("attr.Cnt:", attr.Cnt)

			res := "received"
			go msg.SendResponse(serv, remoteaddr, res)

			UDPenvelope.Envelope = append(UDPenvelope.Envelope, errCheck[:len(errCheck)]...)
			if attr.Cnt >= attr.Size {

				fmt.Println("************")
				fmt.Println("END OF BLOCK")
				fmt.Println("************")

				msg.UDPBlockHandler()
			}
		}
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

func (msg *Message) UDPBlockSender() {
	conn, err := net.Dial("udp", "192.168.1.2:8000")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	defer conn.Close()

	symbols := 128
	extraSymbols := 7

	envelope, err := protoG.Marshal(msg.Block)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshalled data size:", len(envelope))

	// Raptor encoding
	codec := fountain.NewRaptorCodec(symbols, extraSymbols)

	// sent block size first
	attrMsg := AttrMsg{Size: len(envelope), Cnt: 0}
	d, err := json.Marshal(attrMsg)
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write(d)
	if err != nil {
		log.Println(err)
	}

	iter := int(math.Ceil(float64(len(envelope)) / float64(symbols)))

	var start, end int
	var ltBlks []fountain.LTBlock

	index := make([]int64, symbols+extraSymbols)
	for i := 0; i < symbols+extraSymbols; i++ {
		index[i] = int64(i)
	}

	sum := 0
	for i := 0; i < iter; i++ {
		if (sum + symbols) > len(envelope) {
			start = end
			end = len(envelope)
			sum += (end - start)
		} else {
			start = (symbols * i)
			end = (symbols * (i + 1))
			sum += symbols
		}
		slice := envelope[start:end]

		ltBlks = fountain.EncodeLTBlocks(slice, index, codec)
		ltBlks = append(ltBlks[:46], ltBlks[50:134]...)

		message, err := json.Marshal(ltBlks)
		if err != nil {
			log.Fatal(err)
		}
		_, err = conn.Write(message)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Wrote data size:", sum)

		p := make([]byte, 1024)
		_, err = bufio.NewReader(conn).Read(p)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}
		resMsg := string(p)
		log.Println("Received size of the block:", resMsg)
	}
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

func decoder(encSymbols []fountain.LTBlock, codec fountain.Codec, dec fountain.Decoder) []byte {
	fmt.Println(dec.AddBlocks(encSymbols))
	return dec.Decode()
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

	UDPenvelope.Envelope = make([]byte, 0)

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
