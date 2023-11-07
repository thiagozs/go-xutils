package phone

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ttacon/libphonenumber"
)

type Phone struct{}

func New() *Phone {
	return &Phone{}
}

func (p *Phone) Normalize(phone, country string) (string, error) {

	num, err := libphonenumber.Parse(phone, country)
	if err != nil {
		return "", fmt.Errorf("invalid phone number: %w", err)
	}

	normalizedPhone := libphonenumber.Format(num, libphonenumber.E164)

	if strings.ToUpper(country) == "BR" {
		normalizedPhone = normalizedPhone[3:]
	}

	if strings.Trim(normalizedPhone, "0") == "" {
		return "", errors.New("invalid phone number: all zeros")
	}

	return normalizedPhone, nil
}
