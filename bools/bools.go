package bools

import (
	"strconv"
	"strings"
)

// Parse converts a strict true/false string to bool.
func Parse(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, &strconv.NumError{Func: "ParseBool", Num: s, Err: strconv.ErrSyntax}
	}
}

// Format returns "true" or "false".
func Format(value bool) string { return strconv.FormatBool(value) }
