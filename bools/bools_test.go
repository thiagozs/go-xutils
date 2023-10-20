package bools

import (
	"testing"
)

func TestToBool(t *testing.T) {
	b := New()

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
		result, err := b.ToBool(test.input)
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
	b := New()

	tests := []struct {
		input    bool
		expected string
	}{
		{true, "true"},
		{false, "false"},
	}

	for _, test := range tests {
		result := b.ToString(test.input)
		if result != test.expected {
			t.Errorf("expected %v for input %v but got %v", test.expected, test.input, result)
		}
	}
}
