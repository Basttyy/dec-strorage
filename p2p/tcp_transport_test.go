package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport (t *testing.T) {
	listenAddr := ":4000"
	tr := NewTcpTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)

	assert.Nil(t, tr.Listen())

	select {}
}