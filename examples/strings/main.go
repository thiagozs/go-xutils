package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/randutil"
	xstrings "github.com/thiagozs/go-xutils/v2/strings"
)

func main() {
	fmt.Println("camel case:", xstrings.CamelCase("minha biblioteca Go"))
	fmt.Println("snake case:", xstrings.SnakeCase("MinhaBibliotecaGo"))
	fmt.Println("slug único:", xstrings.UniqueSlug("Introdução ao Go"))
	fmt.Println("SQL LIKE:", xstrings.EscapeSQLLike(`100%_concluído`))

	generator := xstrings.NewGeneratorWithSource(randutil.New(42))
	fmt.Println("texto aleatório:", generator.RandomAlphanumeric(12))
}
