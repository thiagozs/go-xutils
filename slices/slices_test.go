package slices

import (
	"reflect"
	"testing"
)

func TestContainsAll(t *testing.T) {
	tests := []struct {
		requiredKeys []string
		incomingKeys []string
		expected     bool
	}{
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key4"}, false},
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key2"}, false},
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key2", "key3"}, true},
		{[]string{"key1", "key2", "key3"}, []string{"key1", "key2", "key4"}, false},
		{[]string{"key1", "key2", "key3"}, []string{"key3", "key4"}, false},
		{[]string{"key1", "key2", "key3"}, []string{"key3"}, false},
		{[]string{"key1"}, []string{"key1", "extra"}, true},
		{nil, []string{"anything"}, true},
	}

	for _, test := range tests {
		result := ContainsAll(test.incomingKeys, test.requiredKeys)
		if result != test.expected {
			t.Errorf("For requiredKeys: %v and incomingKeys: %v, expected %v, got %v", test.requiredKeys, test.incomingKeys, test.expected, result)
		}
	}
}

func TestTransformsDoNotMutateInput(t *testing.T) {
	input := []string{" Value ", "OTHER"}
	want := append([]string(nil), input...)

	_ = Normalize(input)
	if !reflect.DeepEqual(input, want) {
		t.Fatalf("Normalize mutated its input: got %v, want %v", input, want)
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
		result := TrimSpace(test.input)
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
		result := Unique(test.input)
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
		result := Compact(test.input)
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
		result := Lower(test.input)
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
		result := Upper(test.input)
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
		result := Title(test.input)
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
		result := CamelCase(test.input)
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
		result := SnakeCase(test.input)
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
		result := RemoveStopWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"  duplicate ", "DUPLICATE", " unique "}, []string{"duplicate", "unique"}},
		{[]string{"", "  ", " "}, []string{}},
		{[]string{"  ", "  ", " "}, []string{}},
		{[]string{"  UPPER ", "upper", " Mixed "}, []string{"upper", "mixed"}},
	}

	for _, test := range tests {
		result := Normalize(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
