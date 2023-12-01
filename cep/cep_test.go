package cep

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CepTestSuite struct {
	suite.Suite
	cep *CEP
}

func (s *CepTestSuite) SetupTest() {
	s.cep = New()
}

func (s *CepTestSuite) TearDownTest() {
	s.cep = nil
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
		result := s.cep.IsValid(c.cep)
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
		result := s.cep.Trim(c.cep)
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
		result := s.cep.Format(c.cep)
		assert.Equal(s.T(), c.expected, result)
	}
}

func (s *CepTestSuite) GenerateCepTest() {
	result := s.cep.IsValid(s.cep.Generate())
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
		result := s.cep.Normalize(c.cep)
		assert.Equal(s.T(), c.expected, result)
	}
}

func TestCepTestSuite(t *testing.T) {
	suite.Run(t, new(CepTestSuite))
}
