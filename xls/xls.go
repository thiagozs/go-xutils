package xls

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type XLS struct{}

func New() *XLS {
	return &XLS{}
}

func (x *XLS) ParseToMap(filePath string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening xlsx file: %w", err)
	}

	defer f.Close()

	rows, err := f.GetRows(f.GetSheetList()[0])
	if err != nil {
		return nil, fmt.Errorf("error getting rows from xlsx file: %w", err)
	}

	var records []map[string]string
	for _, row := range rows {
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
		return fmt.Errorf("error opening xlsx file: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetList()[0])
	if err != nil {
		return fmt.Errorf("error getting rows from xlsx file: %w", err)
	}

	csvFile, err := os.Create(csvPath)
	if err != nil {
		return fmt.Errorf("error creating csv file: %w", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing row to csv file: %w", err)
		}
	}

	return nil
}

func (x *XLS) GetHeadersFromMap(mapper []map[string]string) ([]string, error) {
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

func (x *XLS) GetHeaders(xlsxPath string) ([]string, error) {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return nil, fmt.Errorf("error opening xlsx file: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetList()[0])
	if err != nil {
		return nil, fmt.Errorf("error getting rows from xlsx file: %w", err)
	}

	return rows[0], nil
}

func (x *XLS) GetRowsFromMap(mapper []map[string]string) ([][]string, error) {
	var rows [][]string
	for _, v := range mapper {
		var row []string
		for _, vv := range v {
			row = append(row, vv)
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func (x *XLS) GetRows(xlsxPath string) ([][]string, error) {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return nil, fmt.Errorf("error opening xlsx file: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetList()[0])
	if err != nil {
		return nil, fmt.Errorf("error getting rows from xlsx file: %w", err)
	}

	return rows, nil
}
