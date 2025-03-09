package core

import "github.com/Mostbesep/Modularis/crypto"

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
