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
	assert.NotNil(t, tx.From)
}

func TestTransaction_Verify(t *testing.T) {
	data := []byte("foo")
	tx := &Transaction{
		Data: data,
	}
	err := tx.Verify()
	assert.ErrorIs(t, err, TransactionNotSignedErr)

	prvKey := crypto.GeneratePrivateKey()
	err = tx.Sign(prvKey)
	assert.NoError(t, err)
	err = tx.Verify()
	assert.NoError(t, err)

	otherPrvKey := crypto.GeneratePrivateKey()
	tx.From = otherPrvKey.PublicKey()
	err = tx.Verify()
	assert.ErrorIs(t, err, InvalidTransactionSignatureErr)

	tx.Data = []byte("foo")
	tx.Sign(prvKey)
	tx.Data = []byte("bar")
	err = tx.Verify()
	assert.ErrorIs(t, err, InvalidTransactionSignatureErr)

}
