package core

import (
	"fmt"
	"github.com/Mostbesep/Modularis/types"
	"testing"
	"time"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:      1,
		PrevBlocHash: types.RandomHash(),
		Height:       height,
		Timestamp:    uint64(time.Now().Unix()),
	}
	tx := Transaction{
		Data: []byte("Foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestBlock_Hash(t *testing.T) {
	b := randomBlock(0)
	fmt.Println(b.Hash(BlocHasher{}))
}
