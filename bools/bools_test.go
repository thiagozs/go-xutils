package bools

import (
	"testing"
)

func TestToBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		err      bool
	}{
		{"true", true, false},
		{"True", true, false},
		{"TRUE", true, false},
		{"false", false, false},
		{"False", false, false},
		{"FALSE", false, false},
		{"invalid", false, true},
	}

	for _, test := range tests {
		result, err := Parse(test.input)
		if err != nil && !test.err {
			t.Errorf("unexpected error for input %v: %v", test.input, err)
			continue
		}

		if err == nil && test.err {
			t.Errorf("expected error for input %v but got none", test.input)
			continue
		}

		if result != test.expected {
			t.Errorf("expected %v for input %v but got %v", test.expected, test.input, result)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "true"},
		{false, "false"},
	}

	for _, test := range tests {
		result := Format(test.input)
		if result != test.expected {
			t.Errorf("expected %v for input %v but got %v", test.expected, test.input, result)
		}
	}
}
