package email

import (
	"regexp"
	"strings"
)

var reEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsValid reports whether value has a conventional email address shape.
// It intentionally does not claim that the mailbox or domain exists.
func IsValid(value string) bool {
	if len(value) > 254 || !reEmail.MatchString(value) {
		return false
	}
	parts := strings.Split(value, "@")
	if len(parts) != 2 || len(parts[0]) > 64 || strings.HasPrefix(parts[0], ".") ||
		strings.HasSuffix(parts[0], ".") || strings.Contains(parts[0], "..") {
		return false
	}
	for _, label := range strings.Split(parts[1], ".") {
		if label == "" || len(label) > 63 || strings.HasPrefix(label, "-") || strings.HasSuffix(label, "-") {
			return false
		}
	}
	return true
}
