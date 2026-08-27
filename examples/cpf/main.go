package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/cpf"
)

func main() {
	value := "529.982.247-25"
	fmt.Println("válido:", cpf.IsValid(value))
	fmt.Println("normalizado:", cpf.Normalize(value))
	fmt.Println("gerado:", cpf.Generate())
}
