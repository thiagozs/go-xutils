// Package slices provides non-mutating transformations for string slices.
package slices

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	xstr "github.com/thiagozs/go-xutils/v2/strings"
)

// ContainsAll reports whether values contains every required value.
func ContainsAll(values, required []string) bool {
	available := make(map[string]struct{}, len(values))
	for _, value := range values {
		available[value] = struct{}{}
	}
	for _, value := range required {
		if _, exists := available[value]; !exists {
			return false
		}
	}
	return true
}

// TrimSpace trims leading and trailing whitespace from every value.
func TrimSpace(values []string) []string {
	return transform(values, strings.TrimSpace)
}

// Unique removes case-insensitive duplicates while preserving order and spelling.
func Unique(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		key := strings.ToLower(value)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, value)
	}
	return result
}

// Compact removes empty values.
func Compact(values []string) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// CompactUnique removes empty values and case-insensitive duplicates.
func CompactUnique(values []string) []string {
	return Unique(Compact(values))
}

// Normalize trims, lowercases, and removes empty and duplicate values.
func Normalize(values []string) []string {
	return CompactUnique(Lower(TrimSpace(values)))
}

// Lower lowercases every value.
func Lower(values []string) []string {
	return transform(values, strings.ToLower)
}

// Upper uppercases every value.
func Upper(values []string) []string {
	return transform(values, strings.ToUpper)
}

// Title converts every value to Brazilian Portuguese title case.
func Title(values []string) []string {
	title := cases.Title(language.BrazilianPortuguese)
	return transform(values, title.String)
}

// CamelCase converts every value to lower camel case.
func CamelCase(values []string) []string {
	return transform(values, xstr.CamelCase)
}

// SnakeCase converts every value to snake case.
func SnakeCase(values []string) []string {
	return transform(values, xstr.SnakeCase)
}

// RemoveStopWords removes common Brazilian Portuguese stop words from every value.
func RemoveStopWords(values []string) []string {
	return transform(values, xstr.RemoveStopWords)
}

func transform(values []string, fn func(string) string) []string {
	result := make([]string, len(values))
	for index, value := range values {
		result[index] = fn(value)
	}
	return result
}
