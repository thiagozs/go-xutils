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

func (p *Phone) IsValid(phone, country string) bool {
	_, err := p.Normalize(phone, country)
	return err == nil
}

func (p *Phone) Generate(limit int) []string {
	return p.phonegen.Random(limit)
}

func (p *Phone) GenMobile(country string, limit int) []string {
	return p.phonegen.RandomMobile(limit)
}

func (p *Phone) GenLandline(country string, limit int) []string {
	return p.phonegen.RandomLandline(limit)
}

func (p *Phone) GenMobileWithMask(limit int) []string {
	return p.phonegen.RandomMobileWithMask(limit)
}

func (p *Phone) GenLandlineWithMask(limit int) []string {
	return p.phonegen.RandomLandlineWithMask(limit)
}
