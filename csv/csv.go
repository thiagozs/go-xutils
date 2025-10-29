package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/xuri/excelize/v2"
)

const (
	ErrorOpenCSVFile   = "error opening CSV file: %w"
	ErrorReadCSVFile   = "error reading CSV file: %w"
	ErrorSaveXLSXFile  = "error saving XLSX file: %w"
	ErrorMapperIsEmpty = "mapper is empty"
	ErrorReadingHeader = "error reading header from CSV file: %w"
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
		return fmt.Errorf(ErrorOpenCSVFile, err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	// create xlsx file and stream writer to avoid loading all data in memory
	f := excelize.NewFile()
	sheetName := "Sheet1"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf(ErrorSaveXLSXFile, err)
	}

	// use StreamWriter for memory-efficient writes
	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return fmt.Errorf(ErrorSaveXLSXFile, err)
	}

	rowNum := 1
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" || err == io.EOF {
				break
			}
			// for csv reader, io.EOF could be returned; use fmt error
			if err != nil {
				break
			}
		}

		// convert []string to []interface{}
		vals := make([]interface{}, len(record))
		for i, v := range record {
			vals[i] = v
		}

		cell, _ := excelize.CoordinatesToCellName(1, rowNum)
		if err := sw.SetRow(cell, vals); err != nil {
			return fmt.Errorf(ErrorSaveXLSXFile, err)
		}
		rowNum++
	}

	// flush stream
	if err := sw.Flush(); err != nil {
		return fmt.Errorf(ErrorSaveXLSXFile, err)
	}

	f.SetActiveSheet(index)
	if err := f.SaveAs(xlsxFilePath); err != nil {
		return fmt.Errorf(ErrorSaveXLSXFile, err)
	}

	return nil
}

func (c *CSV) GetHeaders(csvFilePath string) ([]string, error) {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf(ErrorOpenCSVFile, err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf(ErrorReadCSVFile, err)
	}

	return headers, nil
}

func (c *CSV) GetHeadersFromMap(mapper []map[string]string) ([]string, error) {
	if len(mapper) == 0 {
		return nil, errors.New(ErrorMapperIsEmpty)
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
		return nil, fmt.Errorf(ErrorOpenCSVFile, err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	// Read and discard the header row
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf(ErrorReadingHeader, err)
	}

	// Read the rest of the rows iteratively to avoid ReadAll for large files
	var rows [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf(ErrorReadCSVFile, err)
		}
		rows = append(rows, record)
	}

	return rows, nil
}

func (c *CSV) GetRowsFromMap(mapper []map[string]string) ([]map[string]string, error) {
	if len(mapper) == 0 {
		return nil, errors.New(ErrorMapperIsEmpty)
	}

	dataRows := mapper[1:]

	return dataRows, nil
}
