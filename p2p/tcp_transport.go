package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

//TcpPeer represents a remotet node over a TCP established connection
type TcpPeer struct {
	conn net.Conn
	//if this peer made inbound or outbound connection
	outbound bool
}

type TcpTransport struct {
	listenAddress string
	listener net.Listener
	shakeHands HandshakeFunc
	decoder Decoder
	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTcpPeer(conn net.Conn, outbound bool) *TcpPeer {
	return &TcpPeer{
		conn: conn,
		outbound: outbound,
	}
}

func NewTcpTransport (listenAddr string) *TcpTransport {
	return &TcpTransport{
		shakeHands: NoopHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TcpTransport) Listen() error {
	ln, err := net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}
	t.listener = ln
	
	go t.startAcceptPoll()

	return nil
}

func (t *TcpTransport) startAcceptPoll() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct {}

func (t *TcpTransport) handleConn(conn net.Conn) {
	peer := NewTcpPeer(conn, true)
	if err := t.shakeHands(peer); err != nil {

	}

	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error %s\n", err)
			continue
		}
	}
	fmt.Printf("New incoming connection %+v\n", peer)
}