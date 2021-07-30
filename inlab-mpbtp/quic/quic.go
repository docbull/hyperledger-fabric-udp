package quic

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"math/big"
	"quic-go"

	proto "github.com/hyperledger/fabric-protos-go/gossip"
)

const addr = "192.168.1.7:4242"

var quicStream quic.Stream

// QUIC
type BlockQUIC struct {
	receiver   string
	quicStream quic.Stream
}

func StartQUIC() {
	func() { log.Fatal(QuicServer()) }()
}

// A wrapper for io.Writer for storing the block data
// and response a message to the QUIC sender.
type loggingWriter struct {
	Writer io.Writer
	// quicBlock stores a Block for sending to another
	// peer using QUIC protocol.
	quicBlock *proto.Envelope
}

func QuicBlockSender(block *proto.Envelope) error {
	// tlsConf references TLS information
	tlsConf := QuicTLS()
	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		return err
	}
	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}
	// Send the block data using QUIC protocol.
	_, err = stream.Write([]byte(block.Payload))
	if err != nil {
		return err
	}
	// Receive a response message from the receiver.
	buf := make([]byte, len(block.Payload))
	_, err = io.ReadFull(stream, buf)
	if err != nil {
		return err
	}
	return nil
}

// QuicServer receives block data and store it into own block message.
func QuicServer() error {
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

func QuicTLS() *tls.Config {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"QUIC-HLF"},
	}
	return tlsConf
}
