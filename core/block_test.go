package core

import (
	"bytes"
	"fmt"
	"github.com/Mostbesep/Modularis/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Header_Encode_Decode_Binary(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBloc:  types.RandomHash(),
		Timestamp: uint64(time.Now().Unix()),
		Height:    10,
		Nonce:     989394,
	}

	buf := &bytes.Buffer{}
	err := h.EncodeBinary(buf)
	assert.NoError(t, err)
	receivedHeader := &Header{}
	err = receivedHeader.DecodeBinary(buf)
	assert.NoError(t, err)
	assert.Equal(t, h, receivedHeader)
}

func Test_Bloc_Encode_Decode_Binary(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBloc:  types.RandomHash(),
			Timestamp: uint64(time.Now().Unix()),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	err := b.EncodeBinary(buf)
	assert.NoError(t, err)
	receivedBlock := &Block{}
	err = receivedBlock.DecodeBinary(buf)
	assert.NoError(t, err)
	assert.Equal(t, b, receivedBlock)
	fmt.Println(b, receivedBlock)
}

func Test_Bloc_Hash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBloc:  types.RandomHash(),
			Timestamp: uint64(time.Now().Unix()),
			Height:    10,
		},
		Transactions: nil,
	}

	h := b.Hash()
	assert.False(t, h.IsZero())
	fmt.Println(h)
}
