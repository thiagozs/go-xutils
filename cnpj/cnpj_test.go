package cnpj

import (
	"strings"
	"testing"
)

func TestCNPJ(t *testing.T) {
	t.Run("Generate CNPJ", func(t *testing.T) {
		generatedCNPJ := Generate()
		if len(generatedCNPJ) != 14 {
			t.Errorf("expected length: 14, actual length: %d", len(generatedCNPJ))
		}
	})

	t.Run("Validate CNPJ", func(t *testing.T) {
		invalidCNPJ := "11111111111111"
		validCNPJ := "15757747000166"

		if !IsValid(validCNPJ) {
			t.Errorf("expected valid, but got invalid")
		}

		if IsValid(invalidCNPJ) {
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
				actual := Normalize(test.cnpj)
				if actual != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			})
		}
	})
}

func TestPackageAPIAndRepeatedDigits(t *testing.T) {
	if !IsValid("11.444.777/0001-61") {
		t.Fatal("expected formatted CNPJ to be valid")
	}
	if IsValid("abc11.444.777/0001-61") {
		t.Fatal("CNPJ with arbitrary prefix must be invalid")
	}
	for digit := byte('0'); digit <= '9'; digit++ {
		value := strings.Repeat(string(digit), 14)
		if IsValid(value) {
			t.Fatalf("repeated CNPJ %q must be invalid", value)
		}
	}
	for i := 0; i < 1000; i++ {
		if value := Generate(); !IsValid(value) {
			t.Fatalf("Generate() returned invalid CNPJ %q", value)
		}
	}
}
