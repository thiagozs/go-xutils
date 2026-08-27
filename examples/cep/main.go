package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/cep"
)

func main() {
	value := "01310-100"
	fmt.Println("válido:", cep.IsValid(value))
	fmt.Println("normalizado:", cep.Normalize(value))
	fmt.Println("formatado:", cep.Format(cep.Normalize(value)))
	fmt.Println("gerado:", cep.Generate())
}
