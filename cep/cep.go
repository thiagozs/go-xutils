package cep

import (
	"encoding/csv"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/thiagozs/go-xutils/v2/randutil"
)

// Trim removes the conventional CEP separator.
func Trim(value string) string { return strings.ReplaceAll(value, "-", "") }

// IsValid reports whether value contains exactly eight CEP digits, optionally
// separated by a hyphen.
func IsValid(value string) bool { return reCEPInput.MatchString(value) }

// Format inserts the conventional separator without panicking on short input.
func Format(value string) string {
	value = Trim(value)
	if len(value) <= 5 {
		return value
	}
	return value[:5] + "-" + value[5:]
}

// Generate returns a valid CEP from the embedded Brazilian ranges.
func Generate() string {
	rec, err := loadCepRecords()
	if err != nil || len(rec) == 0 {
		return ""
	}

	cepRand := rec[randutil.Default().Intn(len(rec))]

	randomInRange := func(start, end int) int {
		if start >= end {
			return start
		}
		return start + randutil.Default().Intn(end-start+1)
	}

	cep1, err := strconv.Atoi(cepRand[2])
	if err != nil {
		return ""
	}
	cep2, err := strconv.Atoi(cepRand[3])
	if err != nil {
		return ""
	}

	cepRandNum := randomInRange(cep1, cep2)

	return fmt.Sprintf("%08d", cepRandNum)
}

// Normalize removes every non-digit character from a CEP.
func Normalize(value string) string { return reNonDigits.ReplaceAllString(value, "") }

var (
	cepRecords [][]string
	cepOnce    sync.Once
	cepLoadErr error
)

func loadCepRecords() ([][]string, error) {
	cepOnce.Do(func() {
		r := csv.NewReader(strings.NewReader(cepsrangecsv))
		records, err := r.ReadAll()
		if err != nil {
			cepLoadErr = err
			return
		}
		cepRecords = make([][]string, 0, len(records))
		for _, record := range records {
			if len(record) < 4 || !reCEP.MatchString(record[2]) || !reCEP.MatchString(record[3]) {
				continue
			}
			cepRecords = append(cepRecords, record)
		}
	})
	return cepRecords, cepLoadErr
}

var (
	reCEP       = regexp.MustCompile(`^\d{8}$`)
	reCEPInput  = regexp.MustCompile(`^(?:\d{8}|\d{5}-\d{3})$`)
	reNonDigits = regexp.MustCompile(`\D`)
)
