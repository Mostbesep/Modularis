package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/Mostbesep/Modularis/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlocHasher struct{}

func (BlocHasher) Hash(b *Block) types.Hash {

	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(b.Header); err != nil {
		panic(err)
	}

	h := sha256.Sum256(buf.Bytes())

	return h
}
