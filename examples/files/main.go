package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/thiagozs/go-xutils/v2/files"
)

func main() {
	dir, err := os.MkdirTemp("", "go-xutils-files-")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = os.RemoveAll(dir) }()

	service := files.New()
	path := filepath.Join(dir, "exemplo.txt")
	if err := service.SaveFile(path, []byte("primeira linha\n")); err != nil {
		log.Fatal(err)
	}
	if err := service.AppendFile(path, []byte("segunda linha\n")); err != nil {
		log.Fatal(err)
	}

	lines, err := service.ReadFileLines(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("linhas:", lines)
	fmt.Println("arquivo existe:", service.FileExists(path))
}
