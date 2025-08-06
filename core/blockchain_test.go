package core

import (
	"fmt"
	"github.com/Mostbesep/Modularis/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.NoError(t, err)
	return bc
}

func TestBlockchain_AddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 1; i < lenBlocks; i++ {

		err := bc.AddBlock(randomBlockWithSignature(t, uint32(i), getPreviousHash(t, bc, uint32(i))))
		assert.NoError(t, err)
		assert.Equal(t, uint32(i), bc.Height())
		assert.True(t, bc.HasBlockAt(uint32(i)))
		assert.False(t, bc.HasBlockAt(uint32(i+1)))
	}
	assert.Equal(t, uint32(lenBlocks-1), bc.Height())
	assert.Equal(t, len(bc.headers), lenBlocks)

	err := bc.AddBlock(randomBlock(89, types.Hash{}))
	fmt.Println(err)
	assert.Error(t, err)

}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotEmpty(t, bc.validator)
	assert.Equal(t, uint32(0), bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlockAt(0))
}

func TestAdd_Block_to_Height(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	err := bc.AddBlock(randomBlockWithSignature(t, 2, types.Hash{}))
	fmt.Println(err)
	assert.Error(t, err)
}

func TestBlockchain_GetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenBlocks := 1000

	for i := 1; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i), types.Hash{})
		err := bc.AddBlock(block)
		assert.NoError(t, err)
		header, err := bc.GetHeader(block.Height)
		assert.NoError(t, err)
		assert.Equal(t, block.Header, header)
		if err != nil {
			return
		}
	}
}

func getPreviousHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.NoError(t, err)
	return BlocHasher{}.Hash(prevHeader)
}
