package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/thiagozs/go-xutils/v2/csv"
)

func main() {
	source := strings.NewReader("nome,email\nAna,ana@example.com\nBruno,bruno@example.com\n")
	records, err := csv.Parse(source)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Printf("%s <%s>\n", record["nome"], record["email"])
	}
}
