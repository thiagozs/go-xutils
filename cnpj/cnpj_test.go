package cnpj

import (
	"testing"
)

func TestCNPJ(t *testing.T) {
	c := New()

	t.Run("Generate CNPJ", func(t *testing.T) {
		generatedCNPJ := c.Generate()
		if len(generatedCNPJ) != 14 {
			t.Errorf("expected length: 14, actual length: %d", len(generatedCNPJ))
		}
	})

	t.Run("Validate CNPJ", func(t *testing.T) {
		invalidCNPJ := "11111111111111"
		validCNPJ := "15757747000166"

		if !c.IsValid(validCNPJ) {
			t.Errorf("expected valid, but got invalid")
		}

		if c.IsValid(invalidCNPJ) {
			t.Errorf("expected invalid, but got valid")
		}
	})

	t.Run("Trim CNPJ", func(t *testing.T) {
		type testCase struct {
			cnpj     string
			expected string
		}

		tests := []testCase{
			{"11.444.777/0001-61", "11444777000161"},
			{"11.444.777/0001-61 ", "11444777000161"},
			{" 11...444.777///0001-61", "11444777000161"},
			{" 11.444.777/0001-61 ", "11444777000161"},
			{"11.444.777/0001---61", "11444777000161"},
			{"11.444.777/0001-61adas", "11444777000161"},
			{"11+444+777/0001-61", "11444777000161"},
			{"11.444.777/\\\\0001-61", "11444777000161"},
			{"11444777000161", "11444777000161"},
			{"11.444.777/0001-61$%^", "11444777000161"},
		}

		for _, test := range tests {
			t.Run(test.cnpj, func(t *testing.T) {
				actual := c.TrimCNPJ(test.cnpj)
				if actual != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			})
		}
	})
}
