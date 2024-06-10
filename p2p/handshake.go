package p2p

type HandshakeFunc func(Peer) error

func NoopHandshakeFunc (Peer) error {return nil}