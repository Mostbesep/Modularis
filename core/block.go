package core

import (
	"github.com/Mostbesep/Modularis/crypto"
	"github.com/Mostbesep/Modularis/types"
	"io"
)

type Header struct {
	Version      uint32
	DataHash     types.Hash
	PrevBlocHash types.Hash
	Height       uint32
	Timestamp    uint64
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// Cached version of the header hash
	hash types.Hash
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{Header: h, Transactions: txx}
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) (*Block, error) {
	return dec.Decode(r)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}
	return b.hash
}
