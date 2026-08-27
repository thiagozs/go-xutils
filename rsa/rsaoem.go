package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const DefaultKeyBits = 3072

// GenerateKeyPair generates a key pair and reports generation failures.
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	if bits < 2048 {
		return nil, nil, errors.New("rsa: key size must be at least 2048 bits")
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// ExportPrivateKey returns a PKCS#1 PEM block.
func ExportPrivateKey(privkey *rsa.PrivateKey) (string, error) {
	if privkey == nil {
		return "", ErrNotRSAKey
	}
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem), nil
}

// ExportPublicKey returns a PKIX PEM block.
func ExportPublicKey(pubkey *rsa.PublicKey) (string, error) {
	if pubkey == nil {
		return "", ErrNotRSAKey
	}
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)
	return string(pubkey_pem), nil
}
