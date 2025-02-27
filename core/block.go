package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"github.com/Mostbesep/Modularis/types"
	"io"
)

type Header struct {
	Version   uint32
	PrevBloc  types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.PrevBloc)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.Timestamp)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.Height)
	if err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &h.Nonce)

}

func (h *Header) DecodeBinary(r io.Reader) error {
	err := binary.Read(r, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &h.PrevBloc)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &h.Timestamp)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &h.Height)
	if err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

type Block struct {
	Header
	Transactions []Transaction

	// Cached version of the header hash
	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	b.Header.EncodeBinary(buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {
	err := b.Header.EncodeBinary(w)
	if err != nil {
		return err
	}
	for _, tx := range b.Transactions {
		if err = tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	err := binary.Read(r, binary.LittleEndian, &b.Header)
	if err != nil {
		return err
	}
	for _, tx := range b.Transactions {
		if err = tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil

}
