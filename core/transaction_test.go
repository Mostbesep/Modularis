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
	assert.NotNil(t, tx.signature)
	assert.NotNil(t, tx.publicKey)
}

func TestTransaction_Verify(t *testing.T) {
	data := []byte("foo")
	tx := &Transaction{
		Data: data,
	}
	verify, err := tx.Verify()
	assert.ErrorIs(t, err, TransactionNotSignedErr)
	assert.False(t, verify)

	prvKey := crypto.GeneratePrivateKey()
	err = tx.Sign(prvKey)
	assert.NoError(t, err)
	verify, err = tx.Verify()
	assert.NoError(t, err)
	assert.True(t, verify)

	otherPrvKey := crypto.GeneratePrivateKey()
	tx.publicKey = otherPrvKey.PublicKey()
	verify, err = tx.Verify()
	assert.ErrorIs(t, err, InvalidSignatureErr)
	assert.False(t, verify)

	tx.Data = []byte("foo")
	tx.Sign(prvKey)
	tx.Data = []byte("bar")
	verify, err = tx.Verify()
	assert.ErrorIs(t, err, InvalidSignatureErr)
	assert.False(t, verify)

}
