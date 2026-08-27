package aes

import (
	caes "crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var ErrInvalidCiphertext = errors.New("aes: invalid ciphertext")

// Cipher provides authenticated AES-GCM encryption. It is safe for concurrent
// use and should be preferred over the legacy CBC API.
type Cipher struct {
	aead cipher.AEAD
}

// NewCipher builds an authenticated cipher from a 16, 24, or 32-byte key.
func NewCipher(key []byte) (*Cipher, error) {
	block, err := caes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes: create cipher: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("aes: create GCM: %w", err)
	}
	return &Cipher{aead: aead}, nil
}

// Encrypt returns base64url(nonce || ciphertext || authentication-tag).
func (c *Cipher) Encrypt(plaintext []byte) (string, error) {
	nonce := make([]byte, c.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("aes: generate nonce: %w", err)
	}
	sealed := c.aead.Seal(nonce, nonce, plaintext, nil)
	return base64.RawURLEncoding.EncodeToString(sealed), nil
}

// Decrypt authenticates and decrypts data produced by Encrypt.
func (c *Cipher) Decrypt(encoded string) ([]byte, error) {
	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("aes: decode ciphertext: %w", err)
	}
	nonceSize := c.aead.NonceSize()
	if len(data) < nonceSize+c.aead.Overhead() {
		return nil, ErrInvalidCiphertext
	}
	plaintext, err := c.aead.Open(nil, data[:nonceSize], data[nonceSize:], nil)
	if err != nil {
		return nil, fmt.Errorf("aes: authenticate ciphertext: %w", err)
	}
	return plaintext, nil
}
