package cep

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CepTestSuite struct {
	suite.Suite
}

func (s *CepTestSuite) IsValidTest() {
	cases := []struct {
		cep      string
		expected bool
		result   string
	}{
		{"12345-678", true, "12345678"},
		{"12345678", true, "12345678"},
		{"12345-6789", false, ""},
		{"123456789", false, ""},
		{"1234567", false, ""},
		{"1234567890", false, ""},
		{"12345678901", false, ""},
		{"123456789012", false, ""},
		{"1234567890123", false, ""},
	}

	for _, c := range cases {
		result := IsValid(c.cep)
		assert.Equal(s.T(), c.expected, result)
	}
}

func (s *CepTestSuite) TrimCepTest() {
	cases := []struct {
		cep      string
		expected string
	}{
		{"12345-678", "12345678"},
		{"12345678", "12345678"},
		{"12345-6789", "123456789"},
		{"123456789", "123456789"},
		{"1234567", "1234567"},
		{"1234567890", "1234567890"},
		{"12345678901", "12345678901"},
		{"123456789012", "123456789012"},
		{"1234567890123", "1234567890123"},
	}

	for _, c := range cases {
		result := Trim(c.cep)
		assert.Equal(s.T(), c.expected, result)
	}
}

func (s *CepTestSuite) FormatCepTest() {
	cases := []struct {
		cep      string
		expected string
	}{
		{"12345-678", "12345-678"},
		{"12345678", "12345-678"},
		{"12345-6789", "12345-6789"},
		{"123456789", "12345-6789"},
		{"1234567", "12345-67"},
		{"1234567890", "12345-67890"},
		{"12345678901", "12345-678901"},
		{"123456789012", "12345-6789012"},
		{"1234567890123", "12345-67890123"},
	}

	for _, c := range cases {
		result := Format(c.cep)
		assert.Equal(s.T(), c.expected, result)
	}
}

func (s *CepTestSuite) GenerateCepTest() {
	result := IsValid(Generate())
	assert.True(s.T(), result)
}

func (s *CepTestSuite) NormalizeCepTest() {
	cases := []struct {
		cep      string
		expected string
	}{
		{"12345-678", "12345678"},
		{"12345678", "12345678"},
		{"12345-6789", "123456789"},
		{"123456789PPP", "123456789"},
		{"123****4567", "1234567"},
		{"$$%$1234567890000", "1234567890000"},
		{"@@#$%AAA12345678901", "12345678901"},
		{"&&&12345---6789012", "123456789012"},
		{"#$%1122334455", "1122334455"},
	}

	for _, c := range cases {
		result := Normalize(c.cep)
		assert.Equal(s.T(), c.expected, result)
	}
}

func TestCepTestSuite(t *testing.T) {
	suite.Run(t, new(CepTestSuite))
}

func TestFormatShortInputDoesNotPanic(t *testing.T) {
	if got := Format("123"); got != "123" {
		t.Fatalf("Format short input = %q", got)
	}
}

func TestEmbeddedRangesAndPackageAPI(t *testing.T) {
	records, err := loadCepRecords()
	if err != nil {
		t.Fatal(err)
	}
	if len(records) == 0 {
		t.Fatal("expected valid embedded CEP ranges")
	}
	for _, record := range records {
		if len(record) < 4 || !reCEP.MatchString(record[2]) || !reCEP.MatchString(record[3]) {
			t.Fatalf("invalid embedded record: %#v", record)
		}
	}
	for i := 0; i < 1000; i++ {
		if value := Generate(); !IsValid(value) {
			t.Fatalf("Generate() returned invalid CEP %q", value)
		}
	}
	if got := Format("01000000"); got != "01000-000" {
		t.Fatalf("Format() = %q", got)
	}
	if IsValid("01-000-000") {
		t.Fatal("misplaced CEP separators must be invalid")
	}
}
