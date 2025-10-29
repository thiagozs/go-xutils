package cnpj

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type CNPJ struct{}

func New() *CNPJ {
	return &CNPJ{}
}

// GenerateCNPJ generates a random CNPJ
func (c *CNPJ) Generate() string {
	// Generate the first 12 random digits of the CNPJ
	numbers := make([]int, 12)
	for i := range numbers {
		numbers[i] = rand.Intn(10)
	}

	// Calculate the first check digit
	numbers = append(numbers, c.calculateCheckDigit(numbers))

	// Calculate the second check digit
	numbers = append(numbers, c.calculateCheckDigit(numbers))

	// Convert the CNPJ numbers to a string using strings.Builder
	var b strings.Builder
	for _, number := range numbers {
		b.WriteString(strconv.Itoa(number))
	}
	return b.String()
}

func (c *CNPJ) calculateCheckDigit(numbers []int) int {
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

// IsValidCNPJ validates a CNPJ
func (c *CNPJ) IsValid(cnpj string) bool {
	cnpj = c.TrimCNPJ(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	sum := 0
	for _, digit := range cnpj[:12] {
		num, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
		sum += num
	}

	if sum == 0 {
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

	// Validate the first check digit
	expectedFirstCheckDigit := c.calculateCheckDigit(numbers[:12])
	if expectedFirstCheckDigit != numbers[12] {
		return false
	}

	// Validate the second check digit
	expectedSecondCheckDigit := c.calculateCheckDigit(numbers[:13])
	if expectedSecondCheckDigit != numbers[13] {
		return false
	}

	return true
}

// TrimCNPJ trims CNPJ
func (c *CNPJ) TrimCNPJ(cnpj string) string {
	return reNonDigits.ReplaceAllString(cnpj, "")
}

var reNonDigits = regexp.MustCompile(`\D`)
