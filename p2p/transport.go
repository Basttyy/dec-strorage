package p2p

// Peer is an interface that represents the remote nodes.
type Peer interface {
}

// Transport is anything that handles the communication
// Between the nodes in the network. This can be of the
// form TCP, UDP, Websockets ....
type Transport interface {
	Listen() error
}