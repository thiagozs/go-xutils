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

func (p *Phone) Generate(countryCode string, areaCode string, localNumber string, isMobile bool) (string, error) {

	ccode := libphonenumber.GetCountryCodeForRegion("Brazil")

	fmt.Println(ccode)

	// Combine parts to create a E.164 format phone number
	rawNumber := fmt.Sprintf("+%d%s%s", ccode, areaCode, localNumber)

	// Parse the raw phone number
	num, err := libphonenumber.Parse(rawNumber, countryCode)
	if err != nil {
		return "", fmt.Errorf("failed to parse phone number: %w", err)
	}

	// Check if the number is valid
	if !libphonenumber.IsValidNumber(num) {
		return "", fmt.Errorf("invalid phone number")
	}

	// Get the actual type of the parsed number
	actualType := libphonenumber.GetNumberType(num)

	// Determine the expected number type based on the isMobile flag
	var expectedType libphonenumber.PhoneNumberType
	if isMobile {
		expectedType = libphonenumber.MOBILE
	} else {
		expectedType = libphonenumber.FIXED_LINE
	}

	// Compare the actual type to the expected type
	if actualType != expectedType {
		return "", fmt.Errorf("the generated number type %v does not match the expected type %v", actualType, expectedType)
	}

	// Format the number in international format
	formattedNum := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)

	return formattedNum, nil
}
