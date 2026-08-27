package ip

import (
	"net"
	"strconv"
	"strings"
)

// IsIPv4 validates an IPv4 address in dotted decimal notation.
func IsIPv4(value string) bool {
	parts := strings.Split(value, ".")
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		if len(part) == 0 || len(part) > 3 || (len(part) > 1 && part[0] == '0') {
			return false
		}
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

// IsIPv6 validates an IPv6 address.
func IsIPv6(value string) bool {
	return net.ParseIP(value) != nil && strings.Contains(value, ":")
}

// IsValid reports whether value is either a valid IPv4 or IPv6 address.
func IsValid(value string) bool { return IsIPv4(value) || IsIPv6(value) }
