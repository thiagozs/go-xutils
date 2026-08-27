package main

import (
	"fmt"
	"log"

	xrsa "github.com/thiagozs/go-xutils/v2/rsa"
)

func main() {
	privateKey, publicKey, err := xrsa.GenerateKeyPair(2048)
	if err != nil {
		log.Fatal(err)
	}

	encrypted, err := xrsa.EncryptOAEP(publicKey, []byte("mensagem confidencial"))
	if err != nil {
		log.Fatal(err)
	}
	decrypted, err := xrsa.DecryptOAEP(privateKey, encrypted)
	if err != nil {
		log.Fatal(err)
	}
	publicPEM, err := xrsa.ExportPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("texto original:", string(decrypted))
	fmt.Println("chave pública exportada:", len(publicPEM) > 0)
}
