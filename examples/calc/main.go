package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-xutils/v2/calc"
)

func main() {
	limit, offset, err := calc.LimitOffset(3, 20)
	if err != nil {
		log.Fatal(err)
	}
	random, err := calc.RandomInt32(10, 50)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("paginação: LIMIT %d OFFSET %d\n", limit, offset)
	fmt.Println("número aleatório:", random)
}
