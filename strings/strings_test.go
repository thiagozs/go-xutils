package strings

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/thiagozs/go-xutils/v2/randutil"
)

type TextSuite struct {
	suite.Suite
	generator *Generator
}

func (suite *TextSuite) SetupTest() {
	suite.generator = NewGenerator()
}

func (suite *TextSuite) TestCamelCase() {
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
		{"ÁRVORE bonita", "áRVOREBonita"},
		{"ação rápida", "açãoRápida"},
	}

	for _, test := range tests {
		result := CamelCase(test.input)
		if result != test.expected {
			suite.T().Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, result)
		}
	}
}

func (suite *TextSuite) TestUniqueSlug() {

	tests := []struct {
		input  string
		prefix string
	}{
		{"Hello World", "hello-world-"},
		{"Another Test", "another-test-"},
		{"123 Test", "123-test-"},
		{"Introdução ao Go", "introducao-ao-go-"},
	}

	slugRegex, _ := regexp.Compile("^[a-z0-9]+(-[a-z0-9]+)*-[a-z0-9]{6}$")

	for _, test := range tests {
		slug := UniqueSlug(test.input)
		assert.Regexp(suite.T(), slugRegex, slug, "The slug does not match the expected pattern")
		assert.True(suite.T(), strings.HasPrefix(slug, test.prefix), "The slug does not preserve normalized words")

		// Ensure uniqueness by generating another slug and comparing
		anotherSlug := UniqueSlug(test.input)
		assert.NotEqual(suite.T(), slug, anotherSlug, "The slugs are not unique")
	}
}

func (suite *TextSuite) TestSnakeCase() {

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
		result := SnakeCase(test.input)
		assert.Equal(suite.T(), test.expected, result, "The snake case conversion did not produce the expected result")
	}
}

func (suite *TextSuite) TestRemoveSpecialChars() {

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
		result := RemoveSpecialChars(test.input)
		assert.Equal(suite.T(), test.expected, result, "The remove special characters function did not produce the expected result")
	}
}

func (suite *TextSuite) TestRemoveStopWords() {

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
		result := RemoveStopWords(test.input)
		assert.Equal(suite.T(), test.expected, result, "The remove stop words function did not produce the expected result")
	}
}

func (suite *TextSuite) TestEscapeSQLLike() {

	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "Hello World"},
		{"It's a beautiful day", "It's a beautiful day"},
		{"She said, \"Hello!\"", "She said, \"Hello!\""},
		{"This is a test\\nNew Line", "This is a test\\\\nNew Line"},
		{"Carriage Return\\rTest", "Carriage Return\\\\rTest"},
		{"Comment Test --", "Comment Test --"},
		{"Wildcard_%Test", "Wildcard\\_\\%Test"},
		{"Multi Comment /* Test */", "Multi Comment /* Test */"},
	}

	for _, test := range tests {
		result := EscapeSQLLike(test.input)
		assert.Equal(suite.T(), test.expected, result, "The escape string function did not produce the expected result")
	}
}

func (suite *TextSuite) TestRandom() {
	length := 10
	result := suite.generator.Random(length)
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

func (suite *TextSuite) TestRandomAlphanumeric() {
	length := 10
	result := suite.generator.RandomAlphanumeric(length)
	if len(result) != length {
		suite.T().Errorf("Expected string of length %d, got %d", length, len(result))
	}

	// Check if result contains only alphanumeric characters
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(result) {
		suite.T().Errorf("Result contains special characters: %s", result)
	}
}

func TestTextSuite(t *testing.T) {
	suite.Run(t, new(TextSuite))
}

func TestNewWithSourceIsDeterministic(t *testing.T) {
	first := NewGeneratorWithSource(randutil.New(42)).Random(32)
	second := NewGeneratorWithSource(randutil.New(42)).Random(32)
	if first != second {
		t.Fatalf("injected sources are not deterministic: %q != %q", first, second)
	}
}
