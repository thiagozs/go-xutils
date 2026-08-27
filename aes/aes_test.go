package aes

import (
	"encoding/base64"
	"testing"
)

var testKey = []byte("0123456789abcdef0123456789abcdef")

func TestCipherRoundTripAndAuthentication(t *testing.T) {
	c, err := NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}

	encrypted, err := c.Encrypt([]byte("mensagem confidencial"))
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "mensagem confidencial" {
		t.Fatalf("unexpected plaintext: %q", got)
	}

	tamperedBytes, err := base64.RawURLEncoding.DecodeString(encrypted)
	if err != nil {
		t.Fatal(err)
	}
	tamperedBytes[len(tamperedBytes)-1] ^= 1
	tampered := base64.RawURLEncoding.EncodeToString(tamperedBytes)
	if _, err := c.Decrypt(tampered); err == nil {
		t.Fatal("expected authentication failure for modified ciphertext")
	}
}

func TestCipherRejectsInvalidInput(t *testing.T) {
	if _, err := NewCipher([]byte("short")); err == nil {
		t.Fatal("expected invalid key error")
	}
	c, err := NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Decrypt(""); err == nil {
		t.Fatal("expected invalid ciphertext error")
	}
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	c, err := NewCipher(testKey)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		encrypted, err := c.Encrypt([]byte("123456"))
		if err != nil {
			b.Fatal(err)
		}
		if _, err := c.Decrypt(encrypted); err != nil {
			b.Fatal(err)
		}
	}
}
