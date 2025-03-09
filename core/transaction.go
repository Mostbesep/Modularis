package core

import "github.com/Mostbesep/Modularis/crypto"

type Transaction struct {
	Data      []byte
	publicKey crypto.PublicKey
	signature *crypto.Signature
}
