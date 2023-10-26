package slices

import (
	"reflect"
	"testing"

	"github.com/thiagozs/go-xutils/strings"
)

var (
	s = New(strings.New())
)

func TestAreKeysValid(t *testing.T) {
	// Inicialize a instância de Slices aqui
	// Substitua nil pelo valor apropriado se necessário
	s := New(nil)

	tests := []struct {
		requiredKeys []string
		incomingKeys []string
		expected     bool
	}{
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key4"}, false},
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key2"}, true},
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key2", "key3"}, true},
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key2", "key4"}, false},
		{[]string{"key1", "key2", "key3"}, []string{"key3", "key4"}, false},
		{[]string{"key1", "key2", "key3"}, []string{"key3"}, true},
	}

	for _, test := range tests {
		result := s.AreKeysValid(test.requiredKeys, test.incomingKeys)
		if result != test.expected {
			t.Errorf("For requiredKeys: %v and incomingKeys: %v, expected %v, got %v", test.requiredKeys, test.incomingKeys, test.expected, result)
		}
	}
}

func TestTrimSpaces(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"   space   ", "noSpace", "  trim "}, []string{"space", "noSpace", "trim"}},
		{[]string{"   ", "   ", " "}, []string{"", "", ""}},
	}

	for _, test := range tests {
		result := s.TrimSpaces(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"dup", "dup", "unique"}, []string{"dup", "unique"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		result := s.RemoveDuplicates(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestRemoveEmpty(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"", "notEmpty", "", "alsoNotEmpty"}, []string{"notEmpty", "alsoNotEmpty"}},
		{[]string{"", "", ""}, []string{}},
	}

	for _, test := range tests {
		result := s.RemoveEmpty(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestSliceToLower(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"LOWERCASE", "MixedCase", "UPPERCASE"}, []string{"lowercase", "mixedcase", "uppercase"}},
		{[]string{"alreadylower", "123", "with spaces"}, []string{"alreadylower", "123", "with spaces"}},
	}

	for _, test := range tests {
		result := s.ToLower(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestSliceToUpper(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"lowercase", "MixedCase", "UPPERCASE"}, []string{"LOWERCASE", "MIXEDCASE", "UPPERCASE"}},
		{[]string{"alreadyUPPER", "123", "with spaces"}, []string{"ALREADYUPPER", "123", "WITH SPACES"}},
	}

	for _, test := range tests {
		result := s.ToUpper(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestSliceToTitle(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"title case", "MIXED case", "lowercase"}, []string{"Title Case", "Mixed Case", "Lowercase"}},
	}

	for _, test := range tests {
		result := s.ToTitle(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestSliceToCamel(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"camel case", "mixedCase", "lowercase"}, []string{"camelCase", "mixedCase", "lowercase"}},
		{[]string{"alreadyCamel", "123", "with spaces"}, []string{"alreadyCamel", "123", "withSpaces"}},
	}

	for _, test := range tests {
		result := s.ToCamel(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestSliceToSnake(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"snake_case", "MixedCase", "lowercase"}, []string{"snake_case", "mixed_case", "lowercase"}},
		{[]string{"already_snake", "123", "with spaces"}, []string{"already_snake", "123", "with_spaces"}},
	}

	for _, test := range tests {
		result := s.ToSnake(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestRemoveStopWordsFromSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"o bom amor", "a bandeira e preta"}, []string{"bom amor", "bandeira preta"}},
	}

	for _, test := range tests {
		result := s.RemoveStopWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestRemoveEDTS(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"  duplicate ", "DUPLICATE", " unique "}, []string{"duplicate", "unique"}},
		// Note: Adjust the expected results based on the actual implementation and behavior of RemoveEDTS method
	}

	for _, test := range tests {
		result := s.RemoveEDTS(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
