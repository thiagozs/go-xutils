package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-xutils/v2/bools"
)

func main() {
	value, err := bools.Parse("true")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("valor:", value)
	fmt.Println("formatado:", bools.Format(value))
}
