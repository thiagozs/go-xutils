package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/slices"
)

func main() {
	values := []string{"  Go ", "GOLANG", "go", "", " Bibliotecas "}
	normalized := slices.Normalize(values)

	fmt.Println("entrada preservada:", values)
	fmt.Println("normalizada:", normalized)
	fmt.Println("snake case:", slices.SnakeCase(normalized))
	fmt.Println("contém chaves:", slices.ContainsAll(normalized, []string{"go", "bibliotecas"}))
}
