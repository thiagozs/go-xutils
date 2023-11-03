package xls

import (
	"os"
	"reflect"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestToCSV(t *testing.T) {
	xlsxPath := "./test.xlsx"
	csvPath := "./test.csv"

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Hello")
	f.SetCellValue("Sheet1", "B1", "World")
	if err := f.SaveAs(xlsxPath); err != nil {
		t.Fatalf("Unable to create test xlsx file: %v", err)
	}

	defer os.Remove(xlsxPath)
	defer os.Remove(csvPath)

	x := New()

	err := x.ToCSV(xlsxPath, csvPath)
	if err != nil {
		t.Errorf("ToCSV() error = %v", err)
	}

	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		t.Errorf("CSV file was not created")
	}

}

func TestParseToMap(t *testing.T) {
	xlsxPath := "./testp.xlsx"
	sheetName := "Sheet1"

	f := excelize.NewFile()

	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Header1")
	f.SetCellValue(sheetName, "B1", "Header2")
	f.SetCellValue(sheetName, "A2", "Value1")
	f.SetCellValue(sheetName, "B2", "Value2")

	if err := f.SaveAs(xlsxPath); err != nil {
		t.Fatalf("Unable to create test xlsx file: %v", err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			t.Fatalf("Unable to close the xlsx file: %v", err)
		}
		err = os.Remove(xlsxPath)
		if err != nil {
			t.Fatalf("Unable to remove the xlsx file: %v", err)
		}
	}()

	expected := []map[string]string{
		{"col0": "Header1", "col1": "Header2"},
		{"col0": "Value1", "col1": "Value2"},
	}

	x := New()

	records, err := x.ParseToMap(xlsxPath)
	if err != nil {
		t.Errorf("ParseToMap() error = %v", err)
	}

	if !reflect.DeepEqual(records, expected) {
		t.Errorf("ParseToMap() got = %v, want %v", records, expected)
	}
}

func TestParseToMapError(t *testing.T) {
	// Non-existent file path
	xlsxPath := "./nonexistent.xlsx"

	// Execute the function to test.
	x := New()

	_, err := x.ParseToMap(xlsxPath)
	if err == nil {
		t.Errorf("ParseToMap() expected error, got nil")
	}
}
