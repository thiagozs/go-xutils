package cpf

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/thiagozs/go-xutils/v2/randutil"
)

// Generate returns a valid, unformatted CPF.
func Generate() string {
	// Generate the first 9 random digits of the CPF
	numbers := make([]int, 9)
	for i := range numbers {
		numbers[i] = randutil.Default().Intn(10)
	}

	// Calculate the first check digit
	numbers = append(numbers, calculateCheckDigit(numbers, 10))

	// Calculate the second check digit
	numbers = append(numbers, calculateCheckDigit(numbers, 11))

	// Convert the CPF numbers to a string using strings.Builder
	var b strings.Builder
	for _, number := range numbers {
		b.WriteString(strconv.Itoa(number))
	}
	return b.String()
}

// IsValid validates formatted or unformatted CPF values.
func IsValid(cpf string) bool {
	if !reCPFInput.MatchString(cpf) {
		return false
	}
	cpf = Normalize(cpf)
	if len(cpf) != 11 {
		return false
	}

	// Check if all digits are equal
	allEqual := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return false
	}

	// Convert string to slice of integers
	numbers := make([]int, 11)
	for i, digit := range cpf {
		num, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
		numbers[i] = num
	}

	// Validate the first and second check digits
	if calculateCheckDigit(numbers[:9], 10) != numbers[9] {
		return false
	}
	return calculateCheckDigit(numbers[:10], 11) == numbers[10]
}

func calculateCheckDigit(numbers []int, length int) int {
	sum := 0
	for i, number := range numbers {
		sum += number * (length - i)
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

// Normalize removes every non-digit character from a CPF.
func Normalize(value string) string { return reNonDigits.ReplaceAllString(value, "") }

var (
	reNonDigits = regexp.MustCompile(`\D`)
	reCPFInput  = regexp.MustCompile(`^(?:\d{11}|\d{3}\.\d{3}\.\d{3}-\d{2})$`)
)
