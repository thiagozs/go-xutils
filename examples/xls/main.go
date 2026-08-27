package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/thiagozs/go-xutils/v2/xls"
	"github.com/xuri/excelize/v2"
)

func main() {
	dir, err := os.MkdirTemp("", "go-xutils-xls-")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = os.RemoveAll(dir) }()

	xlsxPath := filepath.Join(dir, "dados.xlsx")
	workbook := excelize.NewFile()
	if err := workbook.SetSheetRow("Sheet1", "A1", &[]any{"nome", "email"}); err != nil {
		log.Fatal(err)
	}
	if err := workbook.SetSheetRow("Sheet1", "A2", &[]any{"Ana", "ana@example.com"}); err != nil {
		log.Fatal(err)
	}
	if err := workbook.SaveAs(xlsxPath); err != nil {
		log.Fatal(err)
	}
	if err := workbook.Close(); err != nil {
		log.Fatal(err)
	}

	service := xls.NewForSheet("Sheet1")
	headers, err := service.GetHeaders(xlsxPath)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := service.GetRows(xlsxPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("cabeçalhos:", headers)
	fmt.Println("linhas:", rows)
}
