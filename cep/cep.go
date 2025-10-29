package cep

import (
	"encoding/csv"
	"regexp"
	"strings"
	"sync"

	"github.com/thiagozs/go-xutils/convs"
	"github.com/thiagozs/go-xutils/randutil"
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
	return reCEP.MatchString(cep)
}

// Format formats a CEP
func (c *CEP) Format(cep string) string {
	cep = c.Trim(cep)
	return cep[:5] + "-" + cep[5:]
}

// Generate generates a random CEP
func (c *CEP) Generate() string {
	rec, err := loadCepRecords()
	if err != nil || len(rec) == 0 {
		return ""
	}

	cepRand := rec[randutil.Global.Intn(len(rec))]

	randomInRange := func(start, end int) int {
		if start >= end {
			return start
		}
		return start + randutil.Global.Intn(end-start+1)
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
	return reNonDigits.ReplaceAllString(cep, "")
}

var (
	cepRecords [][]string
	cepOnce    sync.Once
	cepLoadErr error
)

func loadCepRecords() ([][]string, error) {
	cepOnce.Do(func() {
		r := csv.NewReader(strings.NewReader(cepsrangecsv))
		cepRecords, cepLoadErr = r.ReadAll()
	})
	return cepRecords, cepLoadErr
}

var (
	reCEP       = regexp.MustCompile(`^\d{8}$`)
	reNonDigits = regexp.MustCompile(`\D`)
)
