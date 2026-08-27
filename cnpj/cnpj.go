package cnpj

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/thiagozs/go-xutils/v2/randutil"
)

// Generate returns a valid, unformatted CNPJ.
func Generate() string {
	// Generate the first 12 random digits of the CNPJ
	numbers := make([]int, 12)
	for i := range numbers {
		numbers[i] = randutil.Default().Intn(10)
	}

	// Calculate the first check digit
	numbers = append(numbers, calculateCheckDigit(numbers))

	// Calculate the second check digit
	numbers = append(numbers, calculateCheckDigit(numbers))

	// Convert the CNPJ numbers to a string using strings.Builder
	var b strings.Builder
	for _, number := range numbers {
		b.WriteString(strconv.Itoa(number))
	}
	return b.String()
}

func calculateCheckDigit(numbers []int) int {
	weights := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i, number := range numbers {
		sum += number * weights[len(weights)-len(numbers)+i]
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

// IsValid validates formatted or unformatted CNPJ values.
func IsValid(cnpj string) bool {
	if !reCNPJInput.MatchString(cnpj) {
		return false
	}
	cnpj = Normalize(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	if allDigitsEqual(cnpj) {
		return false
	}

	numbers := make([]int, 14)
	for i, digit := range cnpj {
		num, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
		numbers[i] = num
	}

	// Validate the first and second check digits
	if calculateCheckDigit(numbers[:12]) != numbers[12] {
		return false
	}
	return calculateCheckDigit(numbers[:13]) == numbers[13]
}

func allDigitsEqual(value string) bool {
	for i := 1; i < len(value); i++ {
		if value[i] != value[0] {
			return false
		}
	}
	return true
}

// Normalize removes every non-digit character from a CNPJ.
func Normalize(value string) string { return reNonDigits.ReplaceAllString(value, "") }

var (
	reNonDigits = regexp.MustCompile(`\D`)
	reCNPJInput = regexp.MustCompile(`^(?:\d{14}|\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2})$`)
)
