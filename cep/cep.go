package cep

import (
	"encoding/csv"
	"math/rand"
	"regexp"
	"strings"

	"github.com/thiagozs/go-xutils/convs"
)

type CEP struct {
	conv *convs.Convs
}

func New() *CEP {
	return &CEP{
		conv: convs.New(),
	}
}

// Trim trims a CEP
func (c *CEP) Trim(cep string) string {
	return strings.Replace(cep, "-", "", -1)
}

// IsValid checks if a CEP is valid
func (c *CEP) IsValid(cep string) bool {
	cep = c.Trim(cep)
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

// Format formats a CEP
func (c *CEP) Format(cep string) string {
	cep = c.Trim(cep)
	return cep[:5] + "-" + cep[5:]
}

// Generate generates a random CEP
func (c *CEP) Generate() string {
	rec, err := csv.NewReader(strings.NewReader(cepsrangecsv)).ReadAll()
	if err != nil {
		return ""
	}

	cepRand := rec[rand.Intn(len(rec))]

	randomInRange := func(start, end int) int {
		if start >= end {
			return start
		}
		return start + rand.Intn(end-start+1)
	}

	cep1, _ := c.conv.ToInt(cepRand[2])

	cep2, _ := c.conv.ToInt(cepRand[3])

	cepRandNum := randomInRange(cep1, cep2)

	result, _ := c.conv.ToString(cepRandNum)

	return result
}

// Normalize normalizes a CEP
func (c *CEP) Normalize(cep string) string {
	cep = c.Trim(cep)
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(cep, "")
}
