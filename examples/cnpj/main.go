package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/cnpj"
)

func main() {
	value := "11.222.333/0001-81"
	fmt.Println("válido:", cnpj.IsValid(value))
	fmt.Println("normalizado:", cnpj.Normalize(value))
	fmt.Println("gerado:", cnpj.Generate())
}
