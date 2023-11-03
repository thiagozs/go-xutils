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
