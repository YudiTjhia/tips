package aesh

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AesKeySize int

const (
	AesKeySize16 AesKeySize = 16
	AesKeySize24 AesKeySize = 24
	AesKeySize32 AesKeySize = 32
)

type Aesh struct {
	keyBytes []byte
}

func (a *Aesh) SetKeyFrom64(key64 string) error {
	if key64 == "" {
		return errors.New("aes key64 is empty")
	}
	key, err := base64.RawStdEncoding.DecodeString(key64)
	if err != nil {
		return err
	}
	a.keyBytes = key
	return nil
}

func (a *Aesh) GetKeys64() (string, error) {
	if a.keyBytes == nil {
		return "", errors.New("aes keyBytes is nil")
	}
	key := make([]byte, base64.RawStdEncoding.EncodedLen(len(a.keyBytes)))
	base64.RawStdEncoding.Encode(key, a.keyBytes)
	return string(key), nil
}

func (a *Aesh) GenerateKey(keySize AesKeySize) error {
	if keySize != AesKeySize16 && keySize != AesKeySize24 && keySize != AesKeySize32 {
		return errors.New("invalid AES key size: must be 16, 24, or 32 bytes")
	}

	key := make([]byte, keySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return err
	}
	a.keyBytes = key
	return nil
}

func (a Aesh) validateKey() error {
	if a.keyBytes == nil {
		return errors.New("aes key is nil")
	}
	if len(a.keyBytes) != 16 && len(a.keyBytes) != 24 && len(a.keyBytes) != 32 {
		return errors.New("invalid AES key size: must be 16, 24, or 32 bytes")
	}
	return nil
}

func (a Aesh) Encrypt(plainText string) (string, error) {
	if a.keyBytes == nil {
		return "", errors.New("aes key is nil")
	}
	err := a.validateKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(a.keyBytes))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	cipherText := aesGCM.Seal(nil, nonce, []byte(plainText), nil)
	final := append(nonce, cipherText...)

	base64Final := base64.StdEncoding.EncodeToString(final)
	return base64Final, nil

}

func (a Aesh) Decrypt(cipherTextBase64 string) (string, error) {
	if a.keyBytes == nil {
		return "", errors.New("aes key is nil")
	}
	err := a.validateKey()
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(a.keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aesGCM.NonceSize() {
		return "", errors.New("cipherText too short")
	}

	nonce := cipherText[:aesGCM.NonceSize()]
	cipherData := cipherText[aesGCM.NonceSize():]

	plainText, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil

}
