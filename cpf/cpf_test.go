package cpf

import (
	"testing"
)

func TestIsValidCPF(t *testing.T) {
	c := New()
	tests := []struct {
		cpf      string
		expected bool
	}{
		{"12345678909", true},  // valid CPF
		{"39053344705", true},  // valid CPF
		{"39053344704", false}, // invalid check digits
		{"11111111111", false}, // all digits are equal
		{"123", false},         // CPF too short
		{"", false},            // empty string
		{"abcdefghijk", false}, // non-numeric characters
		{"00000000000", false}, // invalid CPF (zeros only)
	}

	for _, test := range tests {
		t.Run(test.cpf, func(t *testing.T) {
			actual := c.IsValid(test.cpf)
			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestGenerateCPF(t *testing.T) {
	c := New()

	for i := 0; i < 10; i++ {
		t.Run("Test case", func(t *testing.T) {
			cpf := c.Generate()
			if !c.IsValid(cpf) {
				t.Errorf("generated CPF is not valid: %v", cpf)
			}
		})
	}
}

func TestTrimCPF(t *testing.T) {
	c := New()
	tests := []struct {
		cpf      string
		expected string
	}{
		{"123.456.789-09", "12345678909"},
		{"390.533.447-05", "39053344705"},
		{"39053344705", "39053344705"},
		{"39053344704", "39053344704"},
		{"111.111.111-11", "11111111111"},
		{"123", "123"},
		{"", ""},
		{"abcdefghijk", ""},
		{"000.000.000-00", "00000000000"},
	}

	for _, test := range tests {
		t.Run(test.cpf, func(t *testing.T) {
			actual := c.TrimCPF(test.cpf)
			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}
