package rsah

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Rsa(t *testing.T) {
	rsah := Rsah{}
	err := rsah.GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}
	raw := "123456"
	encrypted, err := rsah.Encrypt(raw)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := rsah.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, raw, decrypted, "Decrypted text should match the original plaintext")
}
