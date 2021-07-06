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
	"runtime"
	"strings"
	"time"

	//"quic-go"
	des "github.com/docbull/hyperledger-fabric-udp/inlab-udp/des"
	udp "github.com/docbull/inlab-fabric-udp-proto"
	protoG "github.com/golang/protobuf/proto"
	proto "github.com/hyperledger/fabric-protos-go/gossip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	//quic "github.com/lucas-clemente/quic-go"
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
	//UDPenvelope.Envelope = UDPenvelope.Envelope[:0]
}

func (msg *Message) UDPBlockHandler() {
	protoEnvelope := &proto.Envelope{}

	fmt.Println("--------- UDP Block Handler ---------")
	fmt.Println("Block size with padding:", len(msg.Fountain.RTData))
	msg.Fountain.RTData = msg.Fountain.RTData[:msg.Fountain.RTSize]
	fmt.Println("Block size without padding:", len(msg.Fountain.RTData))
	fmt.Println("-------------------------------------")

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
	codec := fountain.NewRaptorCodec(128, 4)
	dec := codec.NewDecoder(128 * 4)
	var encSymbols []fountain.LTBlock

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
		errCheck := decoder(encSymbols, codec, dec)
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
	codec := fountain.NewRaptorCodec(symbols, symbolSize)

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
	var ltBlks []fountain.LTBlock

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
		fmt.Println("start:", start)
		fmt.Println("end:", end)

		ltBlks = fountain.EncodeLTBlocks(slice, index, codec)
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

func (msg *Message) BlockDataForUDP(ctx context.Context, envelope *udp.Envelope) (*udp.Status, error) {
	log.Println("Receive Block data from the Peer container")

	msg.Block.Payload = envelope.Payload
	msg.Block.Signature = envelope.Signature
	msg.Block.SecretEnvelope = envelope.SecretEnvelope

	go msg.UDPBlockSender()

	return &udp.Status{Code: udp.StatusCode_Ok}, nil
}

func decoder(encSymbols []fountain.LTBlock, codec fountain.Codec, dec fountain.Decoder) []byte {
	return dec.Decode()
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

/*
const addr = "192.168.1.10:4242"

var quicStream quic.Stream

// QUIC
type BlockQUIC struct {
	receiver   string
	quicStream quic.Stream
}

func (msg *Message) StartQUIC() {
	func() { log.Fatal(msg.QuicServer()) }()
}

// A wrapper for io.Writer for storing the block data
// and response a message to the QUIC sender.
type loggingWriter struct {
	Writer io.Writer
	// quicBlock stores a Block for sending to another
	// peer using QUIC protocol.
	quicBlock *udp.Envelope
}

// QuicServer receives block data and store it into own block message.
func (msg *Message) QuicServer() error {
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		return err
	}
	sess, err := listener.Accept(context.Background())
	if err != nil {
		return err
	}
	quicStream, err = sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	w := loggingWriter{Writer: quicStream}
	if _, err = io.Copy(w, quicStream); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (w loggingWriter) Write(b []byte) (int, error) {
	// Store the received block into own QUIC Block message.
	w.quicBlock.Payload = append(w.quicBlock.Payload, b[:len(b)]...)
	return w.Writer.Write(b)
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

func (msg *Message) QuicClient() error {
	// tlsConf references TLS information
	tlsConf := msg.QuicTLS()
	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		return err
	}
	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}

	// Send the block data using QUIC protocol.
	_, err = stream.Write([]byte(msg.Block.Payload))
	if err != nil {
		return err
	}

	// Receive a response message from the receiver.
	buf := make([]byte, len(msg.Block.Payload))
	_, err = io.ReadFull(stream, buf)
	if err != nil {
		return err
	}

	return nil
}

func (msg *Message) QuicTLS() *tls.Config {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"QUIC-HLF"},
	}

	return tlsConf
}
*/

func start() {
	msg := Message{
		Block:           nil,
		PeerContainerIP: "",
		Fountain:        &RaptorCodec{RTSize: 0, RTData: make([]byte, 0)},
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
	//go msg.StartQUIC()

	// waiting D2D containers' connection
	msg.WaitPeerConnection()
}

func main() {
	runtime.GOMAXPROCS((runtime.NumCPU()))
	fmt.Println("Running CPU cores:", runtime.GOMAXPROCS(0))

	start()
}
