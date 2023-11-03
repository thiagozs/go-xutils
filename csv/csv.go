package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type CSV struct{}

func New() *CSV {
	return &CSV{}
}

func (c *CSV) ParseToMap(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var records []map[string]string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		row := make(map[string]string)
		for i, cell := range record {
			row[header[i]] = cell
		}

		records = append(records, row)
	}

	return records, nil
}

func (c *CSV) ToXLSX(csvFilePath, xlsxFilePath string) error {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV file: %w", err)
	}

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	xlsx.NewSheet(sheetName)

	for i, record := range records {
		for j, field := range record {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
			xlsx.SetCellValue(sheetName, cell, field)
		}
	}

	if err := xlsx.SaveAs(xlsxFilePath); err != nil {
		return fmt.Errorf("error saving XLSX file: %w", err)
	}

	return nil
}

func (c *CSV) GetHeaders(csvFilePath string) ([]string, error) {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %w", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %w", err)
	}

	return headers, nil
}

func (c *CSV) GetHeadersFromMap(mapper []map[string]string) ([]string, error) {
	if len(mapper) == 0 {
		return nil, fmt.Errorf("mapper is empty")
	}

	firstRow := mapper[0]

	var headers []string
	for i := 0; ; i++ {
		key := fmt.Sprintf("col%d", i)
		header, ok := firstRow[key]
		if !ok {
			break
		}
		headers = append(headers, header)
	}

	return headers, nil
}

func (c *CSV) GetRows(csvFilePath string) ([][]string, error) {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %w", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	// Read and discard the header row
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error reading header from CSV file: %w", err)
	}

	// Read the rest of the rows
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %w", err)
	}

	return rows, nil
}

func (c *CSV) GetRowsFromMap(mapper []map[string]string) ([][]string, error) {
	if len(mapper) == 0 {
		return nil, fmt.Errorf("mapper is empty")
	}

	// Start from the second element to skip the headers.
	var rows [][]string
	for _, record := range mapper[1:] { // Skip the first map which contains headers.
		var row []string
		for i := 0; ; i++ {
			key := fmt.Sprintf("col%d", i)
			value, ok := record[key]
			if !ok {
				break
			}
			row = append(row, value)
		}
		rows = append(rows, row)
	}

	return rows, nil
}
