package csv

import (
	"reflect"
	"testing"
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
