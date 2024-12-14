package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listeenAddr := ":4000"
	tr := NewTCPTransport(listeenAddr)
	assert.Equal(t, tr.listenAddress, listeenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}
