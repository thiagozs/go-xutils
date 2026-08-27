package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/email"
)

func main() {
	fmt.Println("válido:", email.IsValid("dev@example.com"))
	fmt.Println("inválido:", email.IsValid("endereço-inválido"))
}
