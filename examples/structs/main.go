package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-xutils/v2/structs"
)

type filters struct {
	Query string   `json:"q"`
	Tags  []string `json:"tags,omitempty"`
	Page  int      `json:"page"`
}

func main() {
	query, err := structs.EncodeQuery(filters{
		Query: "bibliotecas Go",
		Tags:  []string{"go", "utils"},
		Page:  2,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(query)
}
