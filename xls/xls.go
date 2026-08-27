package xls

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/xuri/excelize/v2"
)

const (
	errGettingRows         = "error getting rows from xlsx file: %w"
	errOpeningXLSXFile     = "error opening xlsx file: %w"
	errCreatingCSVFile     = "error creating csv file: %w"
	errWritingRowToCSVFile = "error writing row to csv file: %w"
)

var (
	ErrMapperEmpty       = errors.New("xls: mapper is empty")
	ErrOnlyHeaderPresent = errors.New("xls: mapper only contains a header")
)

type XLS struct {
	Sheet string
}

func New() *XLS {
	return &XLS{}
}

// NewForSheet creates an XLS service pinned to a sheet name.
func NewForSheet(sheet string) *XLS { return &XLS{Sheet: sheet} }

func (x *XLS) sheetName(f *excelize.File) (string, bool) {
	if x.Sheet != "" {
		for _, sheet := range f.GetSheetList() {
			if sheet == x.Sheet {
				return sheet, true
			}
		}
		return "", false
	}
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return "", false
	}
	return sheets[0], true
}

func (x *XLS) ParseToMap(filePath string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf(errOpeningXLSXFile, err)
	}
	defer func() { _ = f.Close() }()

	sheet, ok := x.sheetName(f)
	if !ok {
		return nil, nil
	}

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return nil, fmt.Errorf(errGettingRows, err)
	}
	defer func() { _ = rowsIter.Close() }()

	var records []map[string]string
	for rowsIter.Next() {
		row, err := rowsIter.Columns()
		if err != nil {
			return nil, fmt.Errorf(errGettingRows, err)
		}

		record := make(map[string]string)
		for i, cell := range row {
			record[fmt.Sprintf("col%d", i)] = cell
		}
		records = append(records, record)
	}

	return records, nil
}

func (x *XLS) ToCSV(xlsxPath, csvPath string) error {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return fmt.Errorf(errOpeningXLSXFile, err)
	}
	defer func() { _ = f.Close() }()

	sheet, ok := x.sheetName(f)
	if !ok {
		return nil
	}

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return fmt.Errorf(errGettingRows, err)
	}
	defer func() { _ = rowsIter.Close() }()

	csvFile, err := os.Create(csvPath)
	if err != nil {
		return fmt.Errorf(errCreatingCSVFile, err)
	}
	defer func() { _ = csvFile.Close() }()

	writer := csv.NewWriter(csvFile)

	for rowsIter.Next() {
		row, err := rowsIter.Columns()
		if err != nil {
			return fmt.Errorf(errGettingRows, err)
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf(errWritingRowToCSVFile, err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf(errWritingRowToCSVFile, err)
	}
	if err := csvFile.Close(); err != nil {
		return fmt.Errorf(errWritingRowToCSVFile, err)
	}
	return nil
}

func (x *XLS) GetHeadersFromMap(mapper []map[string]string) ([]string, error) {
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

func (x *XLS) GetHeaders(xlsxPath string) ([]string, error) {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return nil, fmt.Errorf(errOpeningXLSXFile, err)
	}
	defer func() { _ = f.Close() }()

	sheet, ok := x.sheetName(f)
	if !ok {
		return nil, nil
	}

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return nil, fmt.Errorf(errGettingRows, err)
	}
	defer func() { _ = rowsIter.Close() }()

	if !rowsIter.Next() {
		return nil, fmt.Errorf(errGettingRows, io.EOF)
	}
	row, err := rowsIter.Columns()
	if err != nil {
		return nil, fmt.Errorf(errGettingRows, err)
	}

	return row, nil
}

func (x *XLS) GetRowsFromMap(mapper []map[string]string) ([]map[string]string, error) {
	if len(mapper) <= 1 {
		return nil, ErrOnlyHeaderPresent
	}

	var rowsWithoutHeaders []map[string]string

	// Copy over all maps except the first one
	for i, rowMap := range mapper {
		if i == 0 {
			continue
		}
		rowsWithoutHeaders = append(rowsWithoutHeaders, rowMap)
	}

	return rowsWithoutHeaders, nil
}

func (x *XLS) GetRows(xlsxPath string) ([][]string, error) {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return nil, fmt.Errorf(errOpeningXLSXFile, err)
	}
	defer func() { _ = f.Close() }()

	sheet, ok := x.sheetName(f)
	if !ok {
		return nil, nil
	}

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return nil, fmt.Errorf(errGettingRows, err)
	}
	defer func() { _ = rowsIter.Close() }()

	var rows [][]string
	for rowsIter.Next() {
		row, err := rowsIter.Columns()
		if err != nil {
			return nil, fmt.Errorf(errGettingRows, err)
		}
		rows = append(rows, row)
	}

	return rows, nil
}
