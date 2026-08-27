package email

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Valid email",
			email:    "test@example.com",
			expected: true,
		},
		{
			name:     "Invalid email without @",
			email:    "testexample.com",
			expected: false,
		},
		{
			name:     "Invalid email without domain",
			email:    "test@",
			expected: false,
		},
		{
			name:     "Invalid email with special characters",
			email:    "test@!$%^&*().com",
			expected: false,
		},
		{
			name:     "Valid email with numbers and characters",
			email:    "test123@example.co.in",
			expected: true,
		},
		{name: "Invalid consecutive local dots", email: "a..b@example.com", expected: false},
		{name: "Invalid consecutive domain dots", email: "a@example..com", expected: false},
		{name: "Invalid leading domain hyphen", email: "a@-example.com", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsValid(tt.email)
			if actual != tt.expected {
				t.Errorf("expected: %v, actual: %v", tt.expected, actual)
			}
		})
	}
}
