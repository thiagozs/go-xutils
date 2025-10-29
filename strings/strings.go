package strings

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

type Strings struct{}

func New() *Strings {
	return &Strings{}
}

var (
	// precompiled regexes to avoid repeated compilation
	slugReg         = regexp.MustCompile("[^a-z0-9]+")
	nonAlnumRe      = regexp.MustCompile(`[^a-zA-Z0-9\\s]+`)
	nonAlphaNumReg  = regexp.MustCompile("[^a-zA-Z0-9]+")
	lowerToUpperReg = regexp.MustCompile("([a-z0-9])([A-Z])")
	seededRand      = rand.New(rand.NewSource(time.Now().UnixNano()))
	stopWordsMap    = map[string]struct{}{}
)

func init() {
	// initialize stopWords map for O(1) lookup
	for _, w := range stopWords {
		stopWordsMap[w] = struct{}{}
	}
}

// GenerateRandomString generates a random string
func (s *Strings) GenerateUniqueSlug(input string) string {
	input = strings.ToLower(input)
	slug := slugReg.ReplaceAllString(input, "-")

	slug = strings.Trim(slug, "-")

	shortUUID := uuid.New().String()[:6]

	slug = fmt.Sprintf("%s-%s", slug, shortUUID)

	return slug
}

// ToCamelCase converts a string to camel case
func (s *Strings) ToCamelCase(str string) string {
	return s.toCamelCase(str)
}

func (s *Strings) toCamelCase(str string) string {
	processedString := nonAlphaNumReg.ReplaceAllString(str, " ")

	words := strings.Fields(processedString)

	for index, word := range words {
		if index == 0 {
			words[index] = strings.ToLower(string(word[0])) + word[1:]
		} else {
			words[index] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}

	return strings.Join(words, "")
}

// ToSnakeCase converts a string to snake case
func (s *Strings) ToSnakeCase(str string) string {
	return s.toSnakeCase(str)
}

func (s *Strings) toSnakeCase(str string) string {
	str = strings.ReplaceAll(str, " ", "_")
	str = lowerToUpperReg.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(str)
}

// RemoveSpecialChars removes special chars
func (s *Strings) RemoveSpecialChars(str string) string {
	return s.removeSpecialChars(str)
}

func (s *Strings) isBrazilianSpecialChar(r rune) bool {
	specialChars := "áàâãäéèêëíìîïóòôõöúùûüçÁÀÂÃÄÉÈÊËÍÌÎÏÓÒÔÕÖÚÙÛÜÇ"
	return strings.ContainsRune(specialChars, r)
}

func (s *Strings) removeSpecialChars(input string) string {
	var result strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) || s.isBrazilianSpecialChar(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// brazilian stop words
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

// RemoveStopWords removes stop(brazilian-ptBR) words from a string
func (s *Strings) RemoveStopWords(text string) string {
	words := strings.Fields(text)
	var filteredWords []string
	for _, word := range words {
		if _, ok := stopWordsMap[strings.ToLower(word)]; !ok {
			filteredWords = append(filteredWords, word)
		}
	}
	return strings.Join(filteredWords, " ")
}

// EscapeString escapes special characters in the input string.
func (s *Strings) EscapeString(input string) string {
	var builder strings.Builder

	for _, ch := range input {
		switch ch {
		case '\\':
			builder.WriteString(`\\`)
		case '\'':
			builder.WriteString(`\'`)
		case '"':
			builder.WriteString(`\"`)
		case '\n':
			builder.WriteString(`\\n`)
		case '\r':
			builder.WriteString(`\\r`)
		case '_':
			builder.WriteString(`\_`)
		case '%':
			builder.WriteString(`\%`)
		case '*':
			builder.WriteString(`\*`)
		default:
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}

func (s *Strings) RandomStrE(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (s *Strings) RandomStr(length int) string {
	var result string
	for len(result) < length {
		str := s.RandomStrE(length)
		str = nonAlnumRe.ReplaceAllString(str, "")
		result += str
	}

	return result[:length]
}
