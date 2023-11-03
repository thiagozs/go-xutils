package csv

import (
	"os"
	"reflect"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestParseToMap(t *testing.T) {
	c := New()

	tests := []struct {
		filePath string
		expected []map[string]string
		err      bool
	}{
		{
			"data/test.csv",
			[]map[string]string{
				{"Name": "Alice", "Age": "30", "City": "New York"},
				{"Name": "Bob", "Age": "35", "City": "Los Angeles"},
			},
			false,
		},
		{
			"nonexistent.csv",
			nil,
			true,
		},
	}

	for _, test := range tests {
		result, err := c.ParseToMap(test.filePath)
		if err != nil && !test.err {
			t.Errorf("unexpected error for file %v: %v", test.filePath, err)
			continue
		}

		if err == nil && test.err {
			t.Errorf("expected error for file %v but got none", test.filePath)
			continue
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v for file %v but got %v", test.expected, test.filePath, result)
		}
	}
}

func TestToXLSX(t *testing.T) {

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test.*.xlsx")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	c := New()

	err = c.ToXLSX("./data/test.csv", tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to convert CSV to XLSX: %v", err)
	}

	xlsx, err := excelize.OpenFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to open XLSX file: %v", err)
	}

	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		t.Fatalf("Failed to get rows from XLSX file: %v", err)
	}

	if len(rows) != 3 {
		t.Fatalf("Expected 2 rows in XLSX file, got %d", len(rows))
	}

	expected := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "35", "Los Angeles"},
	}

	for i, row := range rows {
		for j, cell := range row {
			if cell != expected[i][j] {
				t.Errorf("Expected cell (%d, %d) to be %s, got %s", i, j, expected[i][j], cell)
			}
		}
	}
}

func TestGetHeaders(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString("Name,Age,City\nAlice,30,NY\n")
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	c := New()

	headers, err := c.GetHeaders(tmpFile.Name())
	if err != nil {
		t.Errorf("GetHeaders() error = %v", err)
		return
	}

	expectedHeaders := []string{"Name", "Age", "City"}
	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("GetHeaders() = %v, want %v", headers, expectedHeaders)
	}
}

func TestGetHeadersFromMap(t *testing.T) {
	c := New()

	mapper := []map[string]string{
		{"col0": "Name", "col1": "Age", "col2": "City"},
		{"col0": "John", "col1": "30", "col2": "New York"},
		{"col0": "Bob", "col1": "35", "col2": "Los Angeles"},
	}

	headers, err := c.GetHeadersFromMap(mapper)
	if err != nil {
		t.Errorf("GetHeadersFromMap() error = %v", err)
		return
	}

	expectedHeaders := []string{"Name", "Age", "City"}

	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("GetHeadersFromMap() = %v, want %v", headers, expectedHeaders)
	}
}

func TestGetRows(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString("Name,Age,City\nJohn,30,New York\n")
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	c := New()

	rows, err := c.GetRows(tmpFile.Name())
	if err != nil {
		t.Errorf("GetRows() error = %v", err)
		return
	}
	expectedRows := [][]string{
		{"John", "30", "New York"},
	}

	if !reflect.DeepEqual(rows, expectedRows) {
		t.Errorf("GetRows() = %v, want %v", rows, expectedRows)
	}
}

func TestGetRowsFromMap(t *testing.T) {
	c := New()

	mapper := []map[string]string{
		{"col0": "Name", "col1": "Age", "col2": "City"},
		{"col0": "John", "col1": "30", "col2": "New York"},
	}

	rows, err := c.GetRowsFromMap(mapper)
	if err != nil {
		t.Errorf("GetRowsFromMap() error = %v", err)
		return
	}

	expectedRows := [][]string{
		{"John", "30", "New York"},
	}
	if !reflect.DeepEqual(rows, expectedRows) {
		t.Errorf("GetRowsFromMap() = %v, want %v", rows, expectedRows)
	}
}
