package rsah

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type Rsah struct {
	PrivateKey       *rsa.PrivateKey
	PrivateKeyString *string
	PublicKeyString  *string
}

func (r *Rsah) GenerateKeys() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	privASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	privPEM := pem.EncodeToMemory(&pem.Block{``
		Type:  "RSA PRIVATE KEY",
		Bytes: privASN1,
	})
	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	strPrivPEM := string(privPEM)
	strPubPEM := string(pubPEM)
	r.PrivateKey = privateKey
	r.PrivateKeyString = &strPrivPEM
	r.PrivateKeyString = &strPubPEM
	return nil
}

func (r Rsah) Encrypt(raw string) (string, error) {
	if raw == "" {
		return "", errors.New("raw is empty")
	}

	if r.PrivateKey == nil {
		return "", errors.New("private key is nil")
	}

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, &r.PrivateKey.PublicKey, []byte(raw))
	if err != nil {
		return "", err
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encryptedBase64, nil
}

func (r Rsah) Decrypt(encrypted64 string) (string, error) {
	if encrypted64 == "" {
		return "", errors.New("encrypted64 is empty")
	}
	if r.PrivateKey == nil {
		return "", errors.New("private key is nil")
	}
	decoded64, err := base64.StdEncoding.DecodeString(encrypted64)
	if err != nil {
		return "", err
	}
	decodedRsa, err := rsa.DecryptPKCS1v15(rand.Reader, r.PrivateKey, decoded64)
	if err != nil {
		return "", err
	}

	return string(decodedRsa), nil
}
