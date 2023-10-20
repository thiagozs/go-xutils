package cpf

import (
	"math/rand"
	"strconv"
	"time"
)

type CPF struct{}

func New() *CPF {
	return &CPF{}
}

// GenerateCPF generates a valid CPF number
func (c *CPF) Generate() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Generate the first 9 random digits of the CPF
	numbers := make([]int, 9)
	for i := range numbers {
		numbers[i] = r.Intn(10)
	}

	// Calculate the first check digit
	numbers = append(numbers, calculateCheckDigit(numbers, 10))

	// Calculate the second check digit
	numbers = append(numbers, calculateCheckDigit(numbers, 11))

	// Convert the CPF numbers to a string
	var cpf string
	for _, number := range numbers {
		cpf += strconv.Itoa(number)
	}

	return cpf
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

	// Validate the first check digit
	expectedFirstCheckDigit := calculateCheckDigit(numbers[:9], 10)
	if expectedFirstCheckDigit != numbers[9] {
		return false
	}

	// Validate the second check digit
	expectedSecondCheckDigit := calculateCheckDigit(numbers[:10], 11)
	if expectedSecondCheckDigit != numbers[10] {
		return false
	}

	return true
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
