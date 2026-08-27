package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/randutil"
)

func main() {
	first := randutil.New(42)
	second := randutil.New(42)

	fmt.Println("sequências determinísticas:", first.Intn(100), second.Intn(100))
	fmt.Println("fonte compartilhada:", randutil.Default().Int63n(1_000))
}
