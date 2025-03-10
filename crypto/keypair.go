package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/Mostbesep/Modularis/types"
	"math/big"
)

type PrivateKey struct {
	Key *ecdsa.PrivateKey
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.Key, data)
	if err != nil {
		return nil, err
	}
	return &Signature{
		r: r,
		s: s,
	}, nil
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		Key: &k.Key.PublicKey,
	}
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(fmt.Sprintf("crypto.GeneratePrivateKey  err-> ecdsa.GenerateKey: %v", err))
	}

	return PrivateKey{Key: key}
}

type PublicKey struct {
	Key *ecdsa.PublicKey
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.Key.Curve, k.Key.X, k.Key.Y)
}

func (k PublicKey) Address() types.Address {
	hash := sha256.Sum256(k.ToSlice())

	return types.AddressFromBytes(hash[len(hash)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (s Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(
		pubKey.Key,
		data,
		s.r,
		s.s)
}

func (s Signature) String() string {
	return fmt.Sprintf("{Signature: s: %d, r: %d}", s.s, s.r)
}
