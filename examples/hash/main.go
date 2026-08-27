package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/hash"
)

func main() {
	digest := hash.MD5("go-xutils")

	fmt.Println("MD5:", digest)
	fmt.Println("é MD5:", hash.IsMD5(digest))
	fmt.Println("cor hexadecimal válida:", hash.IsHexColor("#1a73e8"))
	fmt.Println("Base64 válido:", hash.IsBase64("Z28teHV0aWxz"))
}
