package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-xutils/v2/phone"
)

func main() {
	normalized, err := phone.Normalize("(11) 98765-4321", "BR")
	if err != nil {
		log.Fatal(err)
	}

	generator := phone.New()
	fmt.Println("normalizado:", normalized)
	fmt.Println("válido:", phone.IsValid(normalized, "BR"))
	fmt.Println("celulares gerados:", generator.GenerateMobileWithMask(2))
}
