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
