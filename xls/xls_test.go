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
	xlsxPath := "./nonexistent.xlsx"

	x := New()

	_, err := x.ParseToMap(xlsxPath)
	if err == nil {
		t.Errorf("ParseToMap() expected error, got nil")
	}
}

func TestGetHeadersFromMap(t *testing.T) {
	x := New()

	mapper := []map[string]string{
		{"col0": "Name", "col1": "Age", "col2": "City"},
		{"col0": "Test1", "col1": "20", "col2": "NY"},
		{"col0": "Test2", "col1": "15", "col2": "NJ"},
		{"col0": "Test3", "col1": "33", "col2": "CA"},
	}

	expectedHeaders := []string{"Name", "Age", "City"}
	headers, err := x.GetHeadersFromMap(mapper)
	if err != nil {
		t.Errorf("GetHeadersFromMap returned an error: %v", err)
	}
	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("GetHeadersFromMap returned %v, expected %v", headers, expectedHeaders)
	}
}

func TestGetHeaders(t *testing.T) {
	x := New()
	expectedHeaders := []string{"Name", "Age", "City"} // Replace with actual headers from your test.xlsx
	headers, err := x.GetHeaders("./data/test.xlsx")
	if err != nil {
		t.Errorf("GetHeaders returned an error: %v", err)
	}
	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("GetHeaders returned %v, expected %v", headers, expectedHeaders)
	}
}

func TestGetRowsFromMap(t *testing.T) {
	x := New()
	mapper := []map[string]string{
		{"col1": "value1", "col2": "value2"},
		{"col1": "value3", "col2": "value4"},
	}

	expectedRows := [][]string{
		{"value1", "value2"},
		{"value3", "value4"},
	}
	rows, err := x.GetRowsFromMap(mapper)
	if err != nil {
		t.Errorf("GetRowsFromMap returned an error: %v", err)
	}
	if !reflect.DeepEqual(rows, expectedRows) {
		t.Errorf("GetRowsFromMap returned %v, expected %v", rows, expectedRows)
	}
}

func TestGetRows(t *testing.T) {
	x := New()

	expectedRows := [][]string{
		{"Name", "Age", "City"},
		{"Test1", "20", "NY"},
		{"Test2", "15", "NJ"},
		{"Test3", "33", "CA"},
	}

	rows, err := x.GetRows("./data/test.xlsx")
	if err != nil {
		t.Errorf("GetRows returned an error: %v", err)
	}
	if !reflect.DeepEqual(rows, expectedRows) {
		t.Errorf("GetRows returned %v, expected %v", rows, expectedRows)
	}
}
