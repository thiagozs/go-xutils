package rsa

import (
	"errors"
	"testing"
)

func TestOAEPAndPEMRoundTrip(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPair(2048)
	if err != nil {
		t.Fatal(err)
	}
	privatePEM, err := ExportPrivateKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	publicPEM, err := ExportPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}
	parsedPrivate, err := ParsePrivateKey(privatePEM)
	if err != nil {
		t.Fatal(err)
	}
	parsedPublic, err := ParsePublicKey(publicPEM)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext, err := EncryptOAEP(parsedPublic, []byte("secret"))
	if err != nil {
		t.Fatal(err)
	}
	plaintext, err := DecryptOAEP(parsedPrivate, ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	if string(plaintext) != "secret" {
		t.Fatalf("unexpected plaintext: %q", plaintext)
	}
}

func TestInvalidKeys(t *testing.T) {
	if _, err := ParsePublicKey("invalid"); !errors.Is(err, ErrInvalidPEM) {
		t.Fatalf("expected ErrInvalidPEM, got %v", err)
	}
	if _, err := ParsePrivateKey("invalid"); !errors.Is(err, ErrInvalidPEM) {
		t.Fatalf("expected ErrInvalidPEM, got %v", err)
	}
	if _, _, err := GenerateKeyPair(1024); err == nil {
		t.Fatal("expected weak key size rejection")
	}
	if _, err := ExportPrivateKey(nil); !errors.Is(err, ErrNotRSAKey) {
		t.Fatalf("expected ErrNotRSAKey, got %v", err)
	}
	if _, err := ExportPublicKey(nil); !errors.Is(err, ErrNotRSAKey) {
		t.Fatalf("expected ErrNotRSAKey, got %v", err)
	}
}
