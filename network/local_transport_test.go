package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LocalTransport_Connect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	err := tra.Connect(trb)
	assert.NoError(t, err)
	err = trb.Connect(tra)
	assert.NoError(t, err)
	//assert.Equal(t, tra.peers[trb.Addr()], trb)
	//assert.Equal(t, trb.peers[tra.Addr()], tra)
}

func Test_LocalTransport_SendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	tra.Connect(trb)
	trb.Connect(tra)

	message := []byte("hello world")
	err := tra.SendMessage(trb.Addr(), message)
	assert.NoError(t, err)

	receivedRPC := <-trb.Consume()
	assert.Equal(t, message, receivedRPC.payload)
	assert.Equal(t, tra.Addr(), receivedRPC.from)
}
