package core

import (
	"crypto/sha256"
	"github.com/Mostbesep/Modularis/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlocHasher struct{}

func (BlocHasher) Hash(b *Block) types.Hash {
	h := sha256.Sum256(b.HeaderData())
	return types.Hash(h)
}
