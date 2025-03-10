package core

import (
	"github.com/Mostbesep/Modularis/crypto"
	"github.com/Mostbesep/Modularis/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().Unix(),
	}
	tx := Transaction{
		Data: []byte("Foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestBlock_Sign(t *testing.T) {
	prvKey := crypto.GeneratePrivateKey()
	assert.NotNil(t, prvKey)
	b := randomBlock(0)
	err := b.Sign(prvKey)
	assert.NoError(t, err)
	assert.NotNil(t, b.Signature)
}

func TestBlock_Verify(t *testing.T) {
	prvKey := crypto.GeneratePrivateKey()
	assert.NotNil(t, prvKey)
	b := randomBlock(0)
	err := b.Verify()
	assert.ErrorIs(t, err, BlockNotSignedErr)
	err = b.Sign(prvKey)
	assert.NoError(t, err)
	assert.NotNil(t, b.Signature)
	assert.NoError(t, b.Verify())

	otherPrvKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrvKey.PublicKey()
	assert.ErrorIs(t, b.Verify(), InvalidBlockSignatureErr)

}
