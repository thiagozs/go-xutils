package email

import "regexp"

type Email struct{}

func New() *Email {
	return &Email{}
}

var reEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsValidEmail validates an email
func (e *Email) IsValid(email string) bool {
	return reEmail.MatchString(email)
}
