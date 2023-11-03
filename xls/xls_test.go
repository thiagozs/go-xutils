package xls

import (
	"os"
	"reflect"
	"testing"
)

func TestToCSV(t *testing.T) {

	fileTemp, err := os.CreateTemp("", "test.*.csv")
	if err != nil {
		t.Errorf("ToCSV() error = %v", err)
	}

	defer os.Remove(fileTemp.Name())

	x := New()

	err = x.ToCSV("./data/test.xlsx", fileTemp.Name())
	if err != nil {
		t.Errorf("ToCSV() error = %v", err)
	}

	if _, err := os.Stat(fileTemp.Name()); os.IsNotExist(err) {
		t.Errorf("CSV file was not created")
	}

}

func TestParseToMap(t *testing.T) {

	expected := []map[string]string{
		{"col0": "Name", "col1": "Age", "col2": "City"},
		{"col0": "Test1", "col1": "20", "col2": "NY"},
		{"col0": "Test2", "col1": "15", "col2": "NJ"},
		{"col0": "Test3", "col1": "33", "col2": "CA"},
	}

	x := New()

	records, err := x.ParseToMap("./data/test.xlsx")
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
