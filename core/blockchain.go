package core

import (
	"errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidHeaderIndex = errors.New("invalid header index")
)

type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)
	err := bc.addBBlockWithoutValidation(genesis)
	return bc, err
}

func (bc *Blockchain) AddBlock(block *Block) error {
	if err := bc.validator.ValidateBlock(block); err != nil {
		return err
	}
	return bc.addBBlockWithoutValidation(block)

}

func (bc *Blockchain) GetHeader(index uint32) (*Header, error) {
	if index < 0 || index >= uint32(len(bc.headers)) {
		return nil, ErrInvalidHeaderIndex
	}
	return bc.headers[index], nil
}

func (bc *Blockchain) HasBlockAt(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBBlockWithoutValidation(block *Block) error {
	logrus.WithFields(logrus.Fields{
		"Height": block.Height,
		"Hash":   block.Hash(BlocHasher{}),
	}).Info("Adding new Block")

	bc.headers = append(bc.headers, block.Header)

	return bc.store.Put(block)
}
