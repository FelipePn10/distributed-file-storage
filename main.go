package main

import (
	"log"

	"github.com/FelipePn10/distributed-file-storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("TCP listen error: %s", err)
	}
	select {}
}
