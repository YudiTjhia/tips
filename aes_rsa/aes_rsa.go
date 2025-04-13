package aes_rsa

import (
	"errors"
	"sales-services/account/shared/aesh"
	"sales-services/account/shared/rsah"
	"strings"
)

type AesRsa struct {
	aes       aesh.Aesh
	rsa       rsah.Rsah
	separator string
}

func NewAesRsa(aes aesh.Aesh,
	rsah rsah.Rsah,
	separator string) AesRsa {
	return AesRsa{
		aes:       aes,
		rsa:       rsah,
		separator: separator,
	}
}

func (a AesRsa) Encrypt(raw string, aesKeySize aesh.AesKeySize) (string, error) {
	if raw == "" {
		return "", errors.New("raw data is empty")
	}
	if a.rsa.PrivateKey == nil {
		return "", errors.New("private key is nil")
	}
	err := a.aes.GenerateKey(aesKeySize)
	if err != nil {
		return "", err
	}
	encryptedRaw, err := a.aes.Encrypt(raw)
	if err != nil {
		return "", err
	}
	aesKey, err := a.aes.GetKeys64()
	if err != nil {
		return "", err
	}
	encryptedKey, err := a.rsa.Encrypt(aesKey)
	if err != nil {
		return "", err
	}
	return strings.Join([]string{encryptedKey, encryptedRaw}, a.separator), nil
}

func (a AesRsa) Decrypt(encrypted string) (string, error) {
	if encrypted == "" {
		return "", errors.New("encrypted data is empty")
	}
	if a.rsa.PrivateKey == nil {
		return "", errors.New("private key is nil")
	}

	parts := strings.Split(encrypted, a.separator)
	if len(parts) != 2 {
		return "", errors.New("invalid encrypted format")
	}
	encryptedKey := parts[0]
	encryptedRaw := parts[1]

	aesKey, err := a.rsa.Decrypt(encryptedKey)
	if err != nil {
		return "", err
	}
	err = a.aes.SetKeyFrom64(aesKey)
	if err != nil {
		return "", err
	}
	decryptedRaw, err := a.aes.Decrypt(encryptedRaw)
	if err != nil {
		return "", err
	}
	return decryptedRaw, nil
}
