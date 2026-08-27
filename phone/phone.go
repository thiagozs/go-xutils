package phone

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thiagozs/go-phonegen"
	"github.com/ttacon/libphonenumber"
)

type Phone struct {
	phonegen *phonegen.PhoneGen
}

func New() *Phone {
	return &Phone{
		phonegen: phonegen.New(),
	}
}

func (p *Phone) Normalize(phone, country string) (string, error) {
	return Normalize(phone, country)
}

// Normalize parses and validates a phone number and returns E.164 notation.
// For compatibility, Brazilian numbers are returned without the +55 prefix.
func Normalize(phone, country string) (string, error) {
	num, err := libphonenumber.Parse(phone, country)
	if err != nil {
		return "", fmt.Errorf("invalid phone number: %w", err)
	}
	if !libphonenumber.IsValidNumber(num) {
		return "", errors.New("invalid phone number")
	}

	normalizedPhone := libphonenumber.Format(num, libphonenumber.E164)

	if strings.ToUpper(country) == "BR" {
		normalizedPhone = strings.TrimPrefix(normalizedPhone, "+55")
	}

	if strings.Trim(normalizedPhone, "0") == "" {
		return "", errors.New("invalid phone number: all zeros")
	}

	return normalizedPhone, nil
}

func (p *Phone) IsValid(phone, country string) bool {
	return IsValid(phone, country)
}

// IsValid reports whether phone is valid for the supplied ISO region.
func IsValid(phone, country string) bool {
	_, err := Normalize(phone, country)
	return err == nil
}

func (p *Phone) Generate(limit int) []string {
	if limit <= 0 {
		return []string{}
	}
	return p.phonegen.Random(limit)
}

func (p *Phone) GenerateMobileWithMask(limit int) []string {
	if limit <= 0 {
		return []string{}
	}
	return p.phonegen.RandomMobileWithMask(limit)
}

func (p *Phone) GenerateLandlineWithMask(limit int) []string {
	if limit <= 0 {
		return []string{}
	}
	return p.phonegen.RandomLandlineWithMask(limit)
}

// GenerateMobile returns Brazilian mobile numbers.
func (p *Phone) GenerateMobile(limit int) []string {
	if limit <= 0 {
		return []string{}
	}
	return p.phonegen.RandomMobile(limit)
}

// GenerateLandline returns Brazilian landline numbers.
func (p *Phone) GenerateLandline(limit int) []string {
	if limit <= 0 {
		return []string{}
	}
	return p.phonegen.RandomLandline(limit)
}
