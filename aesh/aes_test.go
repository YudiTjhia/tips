package aesh

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Aes(t *testing.T) {
	raw := "123456"
	aesh := Aesh{}
	err := aesh.GenerateKey(AesKeySize32)
	if err != nil {
		t.Fatal(err)
	}

	encrypted, err := aesh.Encrypt(raw)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := aesh.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, raw, decrypted, "Decrypted text should match the original plaintext")
}
