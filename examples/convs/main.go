package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-xutils/v2/convs"
)

func main() {
	quantity, err := convs.Parse[int]("42")
	if err != nil {
		log.Fatal(err)
	}
	formatted, err := convs.Format(quantity)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("valor: %d; texto: %q\n", quantity, formatted)
}
