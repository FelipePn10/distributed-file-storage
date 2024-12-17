package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
}

// Transport is anything that handles the comunication
// between the nodes in the network.
// This can be of the form (TCP, UDP, websockets,...)
type Transport interface {
	listenAndAccept() error
}
