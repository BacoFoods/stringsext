package stringsext

import "testing"

func TestCapitalizeWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"hello world", "Hello World"},
		{"123abc", "123abc"},
		{"a", "A"},
		{"hello beautiful world", "Hello Beautiful World"},
		{"multiple   spaces   here", "Multiple Spaces Here"},
		{"tab\tseparated\twords", "Tab Separated Words"},
		{"new\nline characters", "New Line Characters"},
		{"hyphenated-word", "Hyphenated-word"},
		{"underscore_separated_words", "Underscore_separated_words"},
	}

	for _, test := range tests {
		result := CapitalizeWords(test.input)
		if result != test.expected {
			t.Errorf("CapitalizeWords(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"hello world", "Hello World"},
		{"123abc", "123Abc"},
		{"a", "A"},
		{"hello beautiful world", "Hello Beautiful World"},
	}

	for _, test := range tests {
		result := Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
