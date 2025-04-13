package aes_rsa

import (
	"github.com/stretchr/testify/assert"
	"tips/aesh"
	"tips/rsah"
	"testing"
)

func Test_AesRsa(t *testing.T) {
	aes := aesh.Aesh{}
	rsa := rsah.Rsah{}
	err := rsa.GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}
	aesRsa := NewAesRsa(aes, rsa, "|")
	raw := "123456"
	encrypted, err := aesRsa.Encrypt(raw, aesh.AesKeySize32)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := aesRsa.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, raw, decrypted, "Decrypted text should match the original plaintext")

}
