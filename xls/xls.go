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
	ErrorGetingRows                = "error getting rows from xlsx file: %w"
	ErrorOpeningXLSXFile           = "error opening xlsx file: %w"
	ErrorCreatingCSVFile           = "error creating csv file: %w"
	ErrorWritingRowToCSVFile       = "error writing row to csv file: %w"
	ErrorMapperIsEmpty             = "mapper is empty"
	ErrorMapperOnlyContainsHeaders = "mapper is empty or only contains headers"
)

type XLS struct{}

func New() *XLS {
	return &XLS{}
}

func (x *XLS) ParseToMap(filePath string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf(ErrorOpeningXLSXFile, err)
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return nil, nil
	}
	sheet := sheetList[0]

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return nil, fmt.Errorf(ErrorGetingRows, err)
	}
	defer rowsIter.Close()

	var records []map[string]string
	for rowsIter.Next() {
		row, err := rowsIter.Columns()
		if err != nil {
			return nil, fmt.Errorf(ErrorGetingRows, err)
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
		return fmt.Errorf(ErrorOpeningXLSXFile, err)
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return nil
	}
	sheet := sheetList[0]

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return fmt.Errorf(ErrorGetingRows, err)
	}
	defer rowsIter.Close()

	csvFile, err := os.Create(csvPath)
	if err != nil {
		return fmt.Errorf(ErrorCreatingCSVFile, err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for rowsIter.Next() {
		row, err := rowsIter.Columns()
		if err != nil {
			return fmt.Errorf(ErrorGetingRows, err)
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf(ErrorWritingRowToCSVFile, err)
		}
	}

	return nil
}

func (x *XLS) GetHeadersFromMap(mapper []map[string]string) ([]string, error) {
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

func (x *XLS) GetHeaders(xlsxPath string) ([]string, error) {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return nil, fmt.Errorf(ErrorOpeningXLSXFile, err)
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return nil, nil
	}
	sheet := sheetList[0]

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return nil, fmt.Errorf(ErrorGetingRows, err)
	}
	defer rowsIter.Close()

	if !rowsIter.Next() {
		return nil, fmt.Errorf(ErrorGetingRows, io.EOF)
	}
	row, err := rowsIter.Columns()
	if err != nil {
		return nil, fmt.Errorf(ErrorGetingRows, err)
	}

	return row, nil
}

func (x *XLS) GetRowsFromMap(mapper []map[string]string) ([]map[string]string, error) {
	if len(mapper) <= 1 {
		return nil, errors.New(ErrorMapperIsEmpty)
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
		return nil, fmt.Errorf(ErrorOpeningXLSXFile, err)
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return nil, nil
	}
	sheet := sheetList[0]

	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return nil, fmt.Errorf(ErrorGetingRows, err)
	}
	defer rowsIter.Close()

	var rows [][]string
	for rowsIter.Next() {
		row, err := rowsIter.Columns()
		if err != nil {
			return nil, fmt.Errorf(ErrorGetingRows, err)
		}
		rows = append(rows, row)
	}

	return rows, nil
}
