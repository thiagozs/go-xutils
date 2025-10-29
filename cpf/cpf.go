package cpf

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/thiagozs/go-xutils/randutil"
)

type CPF struct{}

func New() *CPF {
	return &CPF{}
}

// GenerateCPF generates a valid CPF number
func (c *CPF) Generate() string {
	// Generate the first 9 random digits of the CPF
	numbers := make([]int, 9)
	for i := range numbers {
		numbers[i] = randutil.Global.Intn(10)
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

// IsValidCPF validates if the provided CPF is valid
func (c *CPF) IsValid(cpf string) bool {
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

// TrimCPF trims CPF
func (c *CPF) TrimCPF(cpf string) string {
	return reNonDigits.ReplaceAllString(cpf, "")
}

var reNonDigits = regexp.MustCompile(`\D`)
