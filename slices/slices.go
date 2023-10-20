package slices

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	xstr "github.com/thiagozs/go-xutils/strings"
)

type Slices struct {
	*xstr.Strings
}

func New(strs *xstr.Strings) *Slices {
	return &Slices{
		Strings: strs,
	}
}

// AreKeysValid checks if all required keys are present in the incoming keys
func (s *Slices) AreKeysValid(requiredKeys, incomingKeys []string) bool {
	requiredKeysMap := make(map[string]bool)
	incomingKeysMap := make(map[string]bool)

	for _, key := range requiredKeys {
		requiredKeysMap[key] = true
	}

	for _, key := range incomingKeys {
		incomingKeysMap[key] = true
	}

	for key := range requiredKeysMap {
		if _, exists := incomingKeysMap[key]; !exists {
			return false
		}
	}
	return true
}

// TrimSpaces trims spaces
func (s *Slices) TrimSpaces(strs []string) []string {
	for i, v := range strs {
		strs[i] = strings.TrimSpace(v)
	}
	return strs
}

// RemoveDuplicates removes duplicates
func (s *Slices) RemoveDuplicates(strs []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range strs {
		lowerEntry := strings.ToLower(entry)
		if _, value := keys[lowerEntry]; !value {
			keys[lowerEntry] = true
			list = append(list, entry)
		}
	}

	return list
}

// RemoveEmpty removes empty
func (s *Slices) RemoveEmpty(strs []string) []string {
	list := []string{}

	for _, entry := range strs {
		if entry != "" {
			list = append(list, entry)
		}
	}

	return list
}

// RemoveEmptyAndDuplicates removes empty and duplicates
func (s *Slices) RemoveEmptyAndDuplicates(strs []string) []string {
	return s.RemoveDuplicates(s.RemoveEmpty(strs))
}

// RemoveEDTS removes empty, duplicates, trim spaces and converts to lower case
func (s *Slices) RemoveEDTS(strs []string) []string {
	return s.RemoveEmptyAndDuplicates(s.TrimSpaces(strs))
}

// SliceToLower converts all strings in a slice to lower case
func (s *Slices) SliceToLower(strs []string) []string {
	for i, v := range strs {
		strs[i] = strings.ToLower(v)
	}
	return strs
}

// SliceToUpper converts all strings in a slice to upper case
func (s *Slices) ToUpper(strs []string) []string {
	for i, v := range strs {
		strs[i] = strings.ToUpper(v)
	}
	return strs
}

// SliceToTitle converts all strings in a slice to title case
func (s *Slices) ToTitle(strs []string) []string {
	for i, v := range strs {
		strs[i] = cases.Title(language.BrazilianPortuguese).String(v)
	}
	return strs
}

// SliceToCamel converts all strings in a slice to camel case
func (s *Slices) ToCamel(strs []string) []string {
	results := []string{}
	for _, v := range strs {
		results = append(results, s.ToCamelCase(v))
	}
	return results
}

// SliceToSnake converts all strings in a slice to snake case
func (s *Slices) ToSnake(strs []string) []string {
	for i, v := range strs {
		strs[i] = s.ToSnakeCase(v)
	}
	return strs
}

func (s *Slices) RemoveStopWordsFromSlice(strs []string) []string {
	for i, v := range strs {
		strs[i] = s.RemoveStopWords(v)
	}
	return strs
}
