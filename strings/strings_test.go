package strings

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StringsSuite struct {
	suite.Suite
	str *Strings
}

func (suite *StringsSuite) SetupTest() {
	suite.str = New()
}

func (suite *StringsSuite) TestToCamelCase() {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"Hello, world! This is an example string.", "helloWorldThisIsAnExampleString"},
		{"  multiple   spaces  ", "multipleSpaces"},
		{"special@#characters!!", "specialCharacters"},
		{"123 numbers 456", "123Numbers456"},
		{"helloword", "helloword"},
		{"helloWords", "helloWords"},
	}

	for _, test := range tests {
		result := suite.str.ToCamelCase(test.input)
		if result != test.expected {
			suite.T().Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, result)
		}
	}
}

func (suite *StringsSuite) TestGenerateUniqueSlug() {

	tests := []struct {
		input string
	}{
		{"Hello World"},
		{"Another Test"},
		{"123 Test"},
	}

	slugRegex, _ := regexp.Compile("^[a-z0-9]+(-[a-z0-9]+)*-[a-z0-9]{6}$")

	for _, test := range tests {
		slug := suite.str.GenerateUniqueSlug(test.input)
		assert.Regexp(suite.T(), slugRegex, slug, "The slug does not match the expected pattern")

		// Ensure uniqueness by generating another slug and comparing
		anotherSlug := suite.str.GenerateUniqueSlug(test.input)
		assert.NotEqual(suite.T(), slug, anotherSlug, "The slugs are not unique")
	}
}

func (suite *StringsSuite) TestToSnakeCase() {

	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"Another Test", "another_test"},
		{"123 Test", "123_test"},
		{"withSpecialCharacters!", "with_special_characters!"},
	}

	for _, test := range tests {
		result := suite.str.ToSnakeCase(test.input)
		assert.Equal(suite.T(), test.expected, result, "The snake case conversion did not produce the expected result")
	}
}

func (suite *StringsSuite) TestRemoveSpecialChars() {

	specialCharTests := []struct {
		char     rune
		expected bool
	}{
		{'á', true},
		{'A', false},
		{'$', false},
		{'Ç', true},
	}

	for _, test := range specialCharTests {
		result := suite.str.isBrazilianSpecialChar(test.char)
		assert.Equal(suite.T(), test.expected, result, "The special character check did not produce the expected result")
	}

	removeSpecialCharTests := []struct {
		input    string
		expected string
	}{
		{"hello world!", "hello world"},
		{"$100 is 100 dollars.", "100 is 100 dollars"},
		{"áéíóú are Brazilian special chars.", "áéíóú are Brazilian special chars"},
		{"Remove #$%^&*() special characters", "Remove  special characters"},
	}

	for _, test := range removeSpecialCharTests {
		result := suite.str.removeSpecialChars(test.input)
		assert.Equal(suite.T(), test.expected, result, "The remove special characters function did not produce the expected result")
	}
}

func (suite *StringsSuite) TestRemoveStopWords() {

	tests := []struct {
		input    string
		expected string
	}{
		{"Eu amo programar em Go", "amo programar em Go"},
		{"O rato roeu a roupa do rei de Roma", "rato roeu roupa rei Roma"},
		{"Aqui é um lugar maravilhoso para se viver", "é lugar maravilhoso viver"},
		{"Ela sempre soube que ele estava mentindo", "soube estava mentindo"},
		{"Onde você vai estar amanhã à tarde?", "você vai estar à tarde?"},
	}

	for _, test := range tests {
		result := suite.str.RemoveStopWords(test.input)
		assert.Equal(suite.T(), test.expected, result, "The remove stop words function did not produce the expected result")
	}
}

func (suite *StringsSuite) TestEscapeString() {

	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "Hello World"},
		{"It's a beautiful day", "It\\'s a beautiful day"},
		{"She said, \"Hello!\"", "She said, \\\"Hello!\\\""},
		{"This is a test\\nNew Line", "This is a test\\\\nNew Line"},
		{"Carriage Return\\rTest", "Carriage Return\\\\rTest"},
		{"Comment Test --", "Comment Test --"},
		{"Wildcard_%Test", "Wildcard\\_\\%Test"},
		{"Multi Comment /* Test */", "Multi Comment /\\* Test \\*/"},
	}

	for _, test := range tests {
		result := suite.str.EscapeString(test.input)
		assert.Equal(suite.T(), test.expected, result, "The escape string function did not produce the expected result")
	}
}

func (suite *StringsSuite) TestRandomStrE() {
	length := 10
	result := suite.str.RandomStrE(length)
	if len(result) != length {
		suite.T().Errorf("Expected string of length %d, got %d", length, len(result))
	}

	// Check if all characters in result are in the allowed charset
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	for _, char := range result {
		if !strings.ContainsRune(charset, char) {
			suite.T().Errorf("Character '%c' not in allowed charset", char)
		}
	}
}

func (suite *StringsSuite) TestRandomStr() {
	length := 10
	result := suite.str.RandomStr(length)
	if len(result) != length {
		suite.T().Errorf("Expected string of length %d, got %d", length, len(result))
	}

	// Check if result contains only alphanumeric characters
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(result) {
		suite.T().Errorf("Result contains special characters: %s", result)
	}
}

func TestStringsSuite(t *testing.T) {
	suite.Run(t, new(StringsSuite))
}
