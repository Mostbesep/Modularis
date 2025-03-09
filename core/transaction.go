package core

import (
	"errors"
	"github.com/Mostbesep/Modularis/crypto"
)

var (
	TransactionNotSignedErr = errors.New("transaction not signed")
	InvalidSignatureErr     = errors.New("transaction signature or data is invalid")
)

type Transaction struct {
	Data      []byte
	publicKey crypto.PublicKey
	signature *crypto.Signature
}

func (tx *Transaction) Sign(privateKey crypto.PrivateKey) error {
	sig, err := privateKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.publicKey = privateKey.PublicKey()
	tx.signature = sig
	return nil
}

func (tx *Transaction) Verify() (bool, error) {
	if tx.signature == nil {
		return false, TransactionNotSignedErr
	}
	result := tx.signature.Verify(tx.publicKey, tx.Data)
	if !result {
		return false, InvalidSignatureErr
	}
	return result, nil
}
