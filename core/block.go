package core

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/Mostbesep/Modularis/crypto"
	"github.com/Mostbesep/Modularis/types"
	"io"
)

var (
	BlockNotSignedErr        = errors.New("verify block error :block not signed yet")
	InvalidBlockSignatureErr = errors.New("verify block error :signature or data is invalid")
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Height        uint32
	Timestamp     int64
}

func (h Header) Bytes() []byte {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(h)
	if err != nil {
		panic(err) // should never happen
	}
	return buf.Bytes()
}

func (h Header) String() string {
	return fmt.Sprintf("Header:{version:%d ,Height: %d, time: %d, DataHash: %+v, PrevBlockHash: %+v}",
		h.Version, h.Height, h.Timestamp, h.DataHash, h.PrevBlockHash)
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// Cached version of the header hash
	hash types.Hash
}

func (b *Block) Sign(prvKey crypto.PrivateKey) error {
	sign, err := prvKey.Sign(b.Header.Bytes())
	if err != nil {
		return err
	}
	b.Validator = prvKey.PublicKey()
	b.Signature = sign
	return nil

}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return BlockNotSignedErr
	}
	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return InvalidBlockSignatureErr
	}
	return nil
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{Header: h, Transactions: txx}
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}
func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}
	return b.hash
}
