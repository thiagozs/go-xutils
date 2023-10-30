package structs

import "testing"

func TestToQueryParams(t *testing.T) {
	s := New()

	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name: "Test with simple struct",
			input: struct {
				Name  string
				Age   int
				Email string
			}{
				Name:  "John Doe",
				Age:   30,
				Email: "john.doe@example.com",
			},
			expected: "age=30&email=john.doe%40example.com&name=John+Doe",
		},
		{
			name: "Test with struct containing slice",
			input: struct {
				Name    string
				Hobbies []string
			}{
				Name:    "Jane Doe",
				Hobbies: []string{"reading", "cycling", "coding"},
			},
			expected: "hobbies=reading%2Ccycling%2Ccoding&name=Jane+Doe",
		},
		{
			name:     "Test with non-struct type",
			input:    "This is a string",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.ToQueryParams(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
