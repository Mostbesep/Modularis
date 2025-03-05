package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GeneratePrivateKey(t *testing.T) {
	newKey := GeneratePrivateKey()
	assert.NotEmpty(t, newKey)

	pubKey := newKey.PublicKey()
	assert.NotEmpty(t, pubKey)

	address := pubKey.Address()
	assert.NotEmpty(t, address)

	fmt.Println(
		fmt.Sprintf(
			"private :%+v \npublic :%+v,\n address: %+v",
			newKey.Key, pubKey.Key, address))
}
func Test_keypair_Sign_Verify_success(t *testing.T) {
	privateKey := GeneratePrivateKey()
	pubKey := privateKey.PublicKey()

	msg := []byte("Hello World!")

	signedMsg, err := privateKey.Sign(msg)
	assert.NoError(t, err)

	isVerify := signedMsg.Verify(pubKey, msg)
	assert.True(t, isVerify)

	fmt.Println(signedMsg)
}
func Test_keypair_Sign_Verify_failed(t *testing.T) {
	privateKey := GeneratePrivateKey()
	pubKey := privateKey.PublicKey()
	msg := []byte("Hello World!")

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()
	otherMsg := []byte("other Hello World!")

	signedMsg, err := privateKey.Sign(msg)
	assert.NoError(t, err)

	isVerify := signedMsg.Verify(otherPublicKey, msg)
	assert.False(t, isVerify)

	isVerify = signedMsg.Verify(pubKey, otherMsg)
	assert.False(t, isVerify)

	fmt.Println(signedMsg)
}
