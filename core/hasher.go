package core

import (
	"crypto/sha256"
	"github.com/Mostbesep/Modularis/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlocHasher struct{}

func (BlocHasher) Hash(h *Header) types.Hash {
	hash := sha256.Sum256(h.Bytes())
	return types.Hash(hash)
}
