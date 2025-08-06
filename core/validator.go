package core

import (
	"fmt"
)

//var (
//	ErrDuplicateBlock = errors.New("chain already has block at this height")
//)

type Validator interface {
	ValidateBlock(block *Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{bc: bc}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlockAt(b.Height) {
		return fmt.Errorf("validate error > chain already contains block {%d} with hash {%s}",
			b.Height, b.Hash(BlocHasher{}))
	}

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf(
			"validate error > block with hash{%s} height mismatch: expected %d, got %d, chain height: %d",
			b.Hash(BlocHasher{}), v.bc.Height()+1, b.Height, v.bc.Height())
	}

	prevHeader, err := v.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}
	hash := BlocHasher{}.Hash(prevHeader)
	if hash != b.PrevBlockHash {
		return fmt.Errorf("validate error > block with hash{%s} prevhash mismatch: expected {%s}, got {%s}",
			b.Hash(BlocHasher{}), b.PrevBlockHash, hash)
	}

	return b.Verify()
}
