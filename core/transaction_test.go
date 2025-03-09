package core

import (
	"github.com/Mostbesep/Modularis/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_Sign(t *testing.T) {
	data := []byte("foo")
	tx := &Transaction{
		Data: data,
	}
	prvKey := crypto.GeneratePrivateKey()
	err := tx.Sign(prvKey)
	assert.NoError(t, err)
	assert.True(t, tx.signature.Verify(tx.publicKey, data))
}
