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
	errOpenCSVFile   = "error opening CSV file: %w"
	errReadCSVFile   = "error reading CSV file: %w"
	errSaveXLSXFile  = "error saving XLSX file: %w"
	errReadingHeader = "error reading header from CSV file: %w"
)

var (
	ErrEmptyHeader     = errors.New("csv: header names must not be empty")
	ErrDuplicateHeader = errors.New("csv: header names must be unique")
	ErrMapperEmpty     = errors.New("csv: mapper is empty")
)

type CSV struct {
	Comma rune
}

func New() *CSV {
	return &CSV{Comma: ','}
}

func (c *CSV) reader(source io.Reader) *csv.Reader {
	reader := csv.NewReader(source)
	reader.Comma = c.Comma
	return reader
}

func (c *CSV) ParseToMap(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf(errOpenCSVFile, err)
	}
	defer func() { _ = file.Close() }()
	return parse(c.reader(file))
}

// Parse reads a CSV stream using its first record as the map keys. Stream-based
// APIs are easier to test and compose than APIs tied to filesystem paths.
func Parse(source io.Reader) ([]map[string]string, error) {
	return parse(csv.NewReader(source))
}

func parse(reader *csv.Reader) ([]map[string]string, error) {

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf(errReadingHeader, err)
	}
	seenHeaders := make(map[string]struct{}, len(header))
	for _, name := range header {
		if name == "" {
			return nil, ErrEmptyHeader
		}
		if _, exists := seenHeaders[name]; exists {
			return nil, fmt.Errorf("%w: %q", ErrDuplicateHeader, name)
		}
		seenHeaders[name] = struct{}{}
	}

	var records []map[string]string
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf(errReadCSVFile, err)
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
		return fmt.Errorf(errOpenCSVFile, err)
	}
	defer func() { _ = csvFile.Close() }()

	reader := c.reader(csvFile)

	// create xlsx file and stream writer to avoid loading all data in memory
	f := excelize.NewFile()
	defer func() { _ = f.Close() }()
	sheetName := "Sheet1"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf(errSaveXLSXFile, err)
	}

	// use StreamWriter for memory-efficient writes
	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return fmt.Errorf(errSaveXLSXFile, err)
	}

	rowNum := 1
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf(errReadCSVFile, err)
		}

		// convert []string to []interface{}
		vals := make([]interface{}, len(record))
		for i, v := range record {
			vals[i] = v
		}

		cell, err := excelize.CoordinatesToCellName(1, rowNum)
		if err != nil {
			return fmt.Errorf(errSaveXLSXFile, err)
		}
		if err := sw.SetRow(cell, vals); err != nil {
			return fmt.Errorf(errSaveXLSXFile, err)
		}
		rowNum++
	}

	// flush stream
	if err := sw.Flush(); err != nil {
		return fmt.Errorf(errSaveXLSXFile, err)
	}

	f.SetActiveSheet(index)
	if err := f.SaveAs(xlsxFilePath); err != nil {
		return fmt.Errorf(errSaveXLSXFile, err)
	}
	return nil
}

func (c *CSV) GetHeaders(csvFilePath string) ([]string, error) {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf(errOpenCSVFile, err)
	}
	defer func() { _ = csvFile.Close() }()

	reader := c.reader(csvFile)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf(errReadCSVFile, err)
	}

	return headers, nil
}

func (c *CSV) GetHeadersFromMap(mapper []map[string]string) ([]string, error) {
	if len(mapper) == 0 {
		return nil, ErrMapperEmpty
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
		return nil, fmt.Errorf(errOpenCSVFile, err)
	}
	defer func() { _ = csvFile.Close() }()

	reader := c.reader(csvFile)

	// Read and discard the header row
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf(errReadingHeader, err)
	}

	// Read the rest of the rows iteratively to avoid ReadAll for large files
	var rows [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf(errReadCSVFile, err)
		}
		rows = append(rows, record)
	}

	return rows, nil
}

func (c *CSV) GetRowsFromMap(mapper []map[string]string) ([]map[string]string, error) {
	if len(mapper) == 0 {
		return nil, ErrMapperEmpty
	}

	dataRows := mapper[1:]

	return dataRows, nil
}
