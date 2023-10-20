package email

import "regexp"

type Email struct{}

func New() *Email {
	return &Email{}
}

// IsValidEmail validates an email
func (e *Email) IsValid(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
