package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	ErrInvalidPEM = errors.New("rsa: invalid PEM data")
	ErrNotRSAKey  = errors.New("rsa: key is not an RSA key")
)

// ParsePublicKey accepts PKIX PUBLIC KEY and PKCS#1 RSA PUBLIC KEY PEM blocks.
func ParsePublicKey(value string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(value))
	if block == nil {
		return nil, ErrInvalidPEM
	}
	if key, err := x509.ParsePKIXPublicKey(block.Bytes); err == nil {
		publicKey, ok := key.(*rsa.PublicKey)
		if !ok {
			return nil, ErrNotRSAKey
		}
		return publicKey, nil
	}
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("rsa: parse public key: %w", err)
	}
	return publicKey, nil
}

// ParsePrivateKey accepts PKCS#1 RSA PRIVATE KEY and unencrypted PKCS#8 PEM.
func ParsePrivateKey(value string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(value))
	if block == nil {
		return nil, ErrInvalidPEM
	}
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("rsa: parse private key: %w", err)
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, ErrNotRSAKey
	}
	return privateKey, nil
}

// EncryptOAEP encrypts plaintext using RSA-OAEP with SHA-256.
func EncryptOAEP(publicKey *rsa.PublicKey, plaintext []byte) (string, error) {
	if publicKey == nil {
		return "", ErrNotRSAKey
	}
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
	if err != nil {
		return "", fmt.Errorf("rsa: encrypt OAEP: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

// DecryptOAEP decrypts ciphertext produced by EncryptOAEP.
func DecryptOAEP(privateKey *rsa.PrivateKey, encoded string) ([]byte, error) {
	if privateKey == nil {
		return nil, ErrNotRSAKey
	}
	ciphertext, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("rsa: decode ciphertext: %w", err)
	}
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("rsa: decrypt OAEP: %w", err)
	}
	return plaintext, nil
}
