// Package strings provides text transformations and random string generation.
package strings

import (
	"fmt"
	"regexp"
	stdstrings "strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/thiagozs/go-xutils/v2/randutil"
	"golang.org/x/text/unicode/norm"
)

var (
	slugReg         = regexp.MustCompile("[^a-z0-9]+")
	lowerToUpperReg = regexp.MustCompile("([a-z0-9])([A-Z])")
	stopWordsMap    = buildStopWordsMap()
)

// Generator generates random strings from an injectable random source.
type Generator struct {
	rng *randutil.Source
}

// NewGenerator creates a generator backed by the package's default source.
func NewGenerator() *Generator {
	return NewGeneratorWithSource(randutil.Default())
}

// NewGeneratorWithSource creates a generator with an injectable random source.
func NewGeneratorWithSource(source *randutil.Source) *Generator {
	if source == nil {
		source = randutil.Default()
	}
	return &Generator{rng: source}
}

// UniqueSlug converts input to a slug and appends a random six-character suffix.
func UniqueSlug(input string) string {
	input = removeDiacritics(stdstrings.ToLower(input))
	slug := stdstrings.Trim(slugReg.ReplaceAllString(input, "-"), "-")
	return fmt.Sprintf("%s-%s", slug, uuid.New().String()[:6])
}

func removeDiacritics(input string) string {
	var result stdstrings.Builder
	for _, r := range norm.NFD.String(input) {
		if !unicode.Is(unicode.Mn, r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// CamelCase converts text to lower camel case while preserving Unicode.
func CamelCase(str string) string {
	words := stdstrings.FieldsFunc(str, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	for index, word := range words {
		if index == 0 {
			words[index] = lowerFirst(word)
		} else {
			words[index] = upperFirst(stdstrings.ToLower(word))
		}
	}
	return stdstrings.Join(words, "")
}

func lowerFirst(value string) string {
	runes := []rune(value)
	if len(runes) == 0 {
		return ""
	}
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func upperFirst(value string) string {
	runes := []rune(value)
	if len(runes) == 0 {
		return ""
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// SnakeCase converts spaces and lower-to-upper boundaries to underscores.
func SnakeCase(str string) string {
	str = stdstrings.ReplaceAll(str, " ", "_")
	str = lowerToUpperReg.ReplaceAllString(str, "${1}_${2}")
	return stdstrings.ToLower(str)
}

// RemoveSpecialChars removes characters other than letters, numbers and spaces.
func RemoveSpecialChars(input string) string {
	var result stdstrings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// RemoveStopWords removes common Brazilian Portuguese stop words.
func RemoveStopWords(text string) string {
	words := stdstrings.Fields(text)
	filteredWords := make([]string, 0, len(words))
	for _, word := range words {
		if _, ok := stopWordsMap[stdstrings.ToLower(word)]; !ok {
			filteredWords = append(filteredWords, word)
		}
	}
	return stdstrings.Join(filteredWords, " ")
}

// EscapeSQLLike escapes backslashes and wildcard characters for a SQL LIKE
// pattern that declares backslash as its ESCAPE character.
func EscapeSQLLike(input string) string {
	var builder stdstrings.Builder
	for _, ch := range input {
		switch ch {
		case '\\':
			builder.WriteString(`\\`)
		case '_':
			builder.WriteString(`\_`)
		case '%':
			builder.WriteString(`\%`)
		default:
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}

// Random generates a string from letters, numbers and symbols.
func (g *Generator) Random(length int) string {
	if length <= 0 {
		return ""
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[g.rng.Intn(len(charset))]
	}
	return string(result)
}

// RandomAlphanumeric generates a string containing only letters and numbers.
func (g *Generator) RandomAlphanumeric(length int) string {
	if length <= 0 {
		return ""
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[g.rng.Intn(len(charset))]
	}
	return string(result)
}

func buildStopWordsMap() map[string]struct{} {
	result := make(map[string]struct{}, len(stopWords))
	for _, word := range stopWords {
		result[word] = struct{}{}
	}
	return result
}

var stopWords = []string{
	"o", "a", "os", "as", "um", "uma", "uns", "umas", "de", "do", "da", "dos", "das",
	"para", "pra", "por", "per", "com", "sem", "sob", "sobre", "entre", "dentro", "e",
	"mas", "porém", "contudo", "ou", "porque", "pois", "quando", "enquanto", "se", "eu",
	"tu", "ele", "ela", "nós", "vós", "eles", "elas", "me", "te", "se", "lhe", "nos",
	"vos", "lhes", "aqui", "ali", "lá", "agora", "já", "sempre", "nunca", "depois",
	"antes", "tarde", "cedo", "hoje", "ontem", "amanhã", "que", "qual", "quais", "como",
	"onde", "quando", "quanto", "quanta", "quantos", "quantas", "este", "esta", "estes",
	"estas", "isso", "isto", "aquilo",
}
