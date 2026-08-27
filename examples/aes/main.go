package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-xutils/v2/aes"
)

func main() {
	key := []byte("0123456789abcdef0123456789abcdef")
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	encrypted, err := cipher.Encrypt([]byte("mensagem confidencial"))
	if err != nil {
		log.Fatal(err)
	}
	decrypted, err := cipher.Decrypt(encrypted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("texto cifrado:", encrypted)
	fmt.Println("texto original:", string(decrypted))
}
