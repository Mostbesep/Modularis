package core

import (
	"github.com/Mostbesep/Modularis/crypto"
	"github.com/Mostbesep/Modularis/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().Unix(),
	}
	tx := Transaction{
		Data: []byte("Foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	b := randomBlock(height, prevBlockHash)
	prvKey := crypto.GeneratePrivateKey()
	err := b.Sign(prvKey)
	assert.NoError(t, err)
	assert.NotNil(t, b.Signature)
	return b
}

func TestBlock_Sign(t *testing.T) {
	prvKey := crypto.GeneratePrivateKey()
	assert.NotNil(t, prvKey)
	b := randomBlock(0, types.RandomHash())
	err := b.Sign(prvKey)
	assert.NoError(t, err)
	assert.NotNil(t, b.Signature)
}

func TestBlock_Verify(t *testing.T) {
	prvKey := crypto.GeneratePrivateKey()
	assert.NotNil(t, prvKey)
	b := randomBlock(0, types.RandomHash())
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
