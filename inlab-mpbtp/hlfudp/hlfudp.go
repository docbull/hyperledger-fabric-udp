package hlfudp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"strings"
	"time"

	des "github.com/docbull/hyperledger-fabric-udp/inlab-udp/des"
	udp "github.com/docbull/inlab-fabric-udp-proto"
	rtfountain "github.com/docbull/inlab-fabric-udp/inlab-udp/raptor"
	protoG "github.com/golang/protobuf/proto"
	proto "github.com/hyperledger/fabric-protos-go/gossip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Message means received block data
type Message struct {
	Block           *udp.Envelope
	PeerContainerIP string
	Fountain        *RaptorCodec
	Key             []byte
	IV              []byte
}

type RaptorCodec struct {
	RTSize int
	RTData []byte
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
}

func (msg *Message) UDPBlockHandler() {
	protoEnvelope := &proto.Envelope{}
	msg.Fountain.RTData = msg.Fountain.RTData[:msg.Fountain.RTSize]
	err := protoG.Unmarshal(msg.Fountain.RTData, protoEnvelope)
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
	// RT decoding, 128 symbols with 4 symbol size
	codec := rtfountain.NewRaptorCodec(128, 4)
	dec := codec.NewDecoder(128 * 4)
	var encSymbols []rtfountain.LTBlock

	// envelope for block unmarshalling
	// buf for receiving RT symbols of block
	buf := make([]byte, 4096*10)

	n, remoteaddr, err := serv.ReadFromUDP(buf)
	fmt.Println("Message RECEIVE", time.Now())
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	buf = buf[0:n]

	log.Println("Received a Block from", remoteaddr)
	log.Println("Received size of the Block data:", n)

	if n < 30 {
		// receiving block size
		err := json.Unmarshal(buf, &msg.Fountain)
		if err != nil {
			panic(err)
		}
		fmt.Println("Received Block Size:", msg.Fountain.RTSize)
	} else {
		desMsg, _ := des.DesDecryption(msg.Key, msg.IV, buf)
		fmt.Println("DES decoded data size:", len(desMsg))

		// receiving encoding symbol slices
		err = json.Unmarshal(desMsg, &encSymbols)
		if err != nil {
			panic(err)
		}

		// if success to docode, return the original symbols
		errCheck := rtfountain.decoder(encSymbols, codec, dec)
		if errCheck != nil {
			fmt.Println("Complete recovery!")

			res := "received"
			go msg.SendResponse(serv, remoteaddr, res)

			//UDPenvelope.Envelope = append(UDPenvelope.Envelope, errCheck[:len(errCheck)]...)
			msg.Fountain.RTData = append(msg.Fountain.RTData, errCheck[:]...)
			if len(msg.Fountain.RTData) >= msg.Fountain.RTSize {
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
	symbolSize := 4
	redundancySymbols := 7

	envelope, err := protoG.Marshal(msg.Block)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshalled data size:", len(envelope))

	// Raptor encoding
	codec := rtfountain.NewRaptorCodec(symbols, symbolSize)

	// sent block size first
	d, err := json.Marshal(&msg.Fountain)
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write(d)
	if err != nil {
		log.Println(err)
	}

	iter := int(math.Ceil(float64(len(envelope)) / (float64(symbols) * float64(symbolSize))))
	fmt.Println("iter:", iter)

	var start, end int
	var ltBlks []rtfountain.LTBlock

	index := make([]int64, (symbols*symbolSize)+redundancySymbols)
	for i := 0; i < (symbols*symbolSize)+redundancySymbols; i++ {
		index[i] = int64(i)
	}

	sum := 0

	// slices for sending RT symbols
	var slice []byte
	for i := 0; i < iter; i++ {
		if (sum + (symbols * symbolSize)) > len(envelope) {
			start = end
			end = len(envelope)
			slice = envelope[start:end]

			padding := make([]byte, (sum+(symbols*symbolSize))-len(envelope))
			slice = append(slice, padding...)
			fmt.Println("size of slice:", len(slice))

			sum += (end - start)
		} else {
			start = ((symbols * symbolSize) * i)
			end = ((symbols * symbolSize) * (i + 1))
			sum += (symbols * symbolSize)
			slice = envelope[start:end]
		}

		ltBlks = rtfountain.EncodeLTBlocks(slice, index, codec)
		fmt.Println("Size of RT Block:", len(ltBlks))
		ltBlks = append(ltBlks[:2], ltBlks[4:(symbols+redundancySymbols)-1]...)

		message, err := json.Marshal(ltBlks)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("length of marshalled msg:", len(message))

		desMsg, _ := des.DesEncryption(msg.Key, msg.IV, message)

		fmt.Println("Message SEND", time.Now())
		n, err := conn.Write(desMsg)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Wrote data size:", n)

		p := make([]byte, 1024)
		_, err = bufio.NewReader(conn).Read(p)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}
		resMsg := string(p)
		log.Println("Reponse message from UDP:", resMsg)
		fmt.Println()
	}
	fmt.Println()
}

// Peer2UDP listens connections from the peer container for
// forwarding block data. This function is working based on gRPC.
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
