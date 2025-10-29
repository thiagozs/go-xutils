package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

type RSA struct{}

func New() *RSA {
	return &RSA{}
}

type RSAPublicKey struct {
	PublicKey string
}

type RSAPrivateKey struct {
	PrivateKey string
}

func (r *RSA) PublicKey(publicKey string) *RSAPublicKey {
	return &RSAPublicKey{
		PublicKey: publicKey,
	}
}

func (r *RSA) PrivateKey(privateKey string) *RSAPrivateKey {
	return &RSAPrivateKey{
		PrivateKey: privateKey,
	}
}

func (pub *RSAPublicKey) Encrypt(encryptStr string) (string, error) {
	// pem
	block, _ := pem.Decode([]byte(pub.PublicKey))

	// x509
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)

	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(encryptStr))
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(encryptedStr), nil
}

func (pri *RSAPrivateKey) Decrypt(decryptStr string) (string, error) {
	// pem
	block, _ := pem.Decode([]byte(pri.PrivateKey))

	// X509
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decryptBytes)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
