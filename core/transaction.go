package core

import (
	"errors"
	"github.com/Mostbesep/Modularis/crypto"
)

var (
	TransactionNotSignedErr        = errors.New("transaction not signed")
	InvalidTransactionSignatureErr = errors.New("transaction signature or data is invalid")
)

type Transaction struct {
	Data      []byte
	From      crypto.PublicKey
	signature *crypto.Signature
}

func (tx *Transaction) Sign(privateKey crypto.PrivateKey) error {
	sig, err := privateKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.From = privateKey.PublicKey()
	tx.signature = sig
	return nil
}

func (tx *Transaction) Verify() error {
	if tx.signature == nil {
		return TransactionNotSignedErr
	}
	result := tx.signature.Verify(tx.From, tx.Data)
	if !result {
		return InvalidTransactionSignatureErr
	}
	return nil
}
