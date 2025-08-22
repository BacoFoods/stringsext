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

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected bool
	}{
		// Basic cases
		{"", "", true},
		{"hello", "", true},
		{"", "hello", false},

		// Exact matches
		{"hello", "hello", true},
		{"Hello", "Hello", true},

		// Case insensitive matches
		{"Hello", "hello", true},
		{"hello", "Hello", true},
		{"HELLO", "hello", true},
		{"hello", "HELLO", true},
		{"HeLLo", "hElLo", true},

		// Substring matches
		{"hello world", "world", true},
		{"Hello World", "world", true},
		{"HELLO WORLD", "world", true},
		{"hello world", "WORLD", true},
		{"Hello Beautiful World", "beautiful", true},
		{"Hello Beautiful World", "BEAUTIFUL", true},

		// No matches
		{"hello", "world", false},
		{"Hello", "world", false},
		{"hello world", "xyz", false},
		{"Hello World", "XYZ", false},

		// Single character cases
		{"a", "a", true},
		{"A", "a", true},
		{"a", "A", true},
		{"a", "b", false},
		{"A", "B", false},

		// Partial matches at different positions
		{"hello world", "hell", true},
		{"hello world", "HELL", true},
		{"hello world", "orld", true},
		{"hello world", "ORLD", true},
		{"hello world", "lo w", true},
		{"hello world", "LO W", true},

		// Special characters and numbers
		{"hello123", "123", true},
		{"Hello123", "123", true},
		{"hello-world", "hello-world", true},
		{"Hello-World", "hello-world", true},
		{"hello_world", "HELLO_WORLD", true},
		{"test@example.com", "@example", true},
		{"Test@Example.COM", "@example", true},

		// Unicode characters
		{"café", "café", true},
		{"Café", "café", true},
		{"CAFÉ", "café", true},
		{"naïve", "ïve", true},
		{"Naïve", "ÏVE", true},
	}

	for _, test := range tests {
		result := ContainsIgnoreCase(test.s, test.substr)
		if result != test.expected {
			t.Errorf("ContainsIgnoreCase(%q, %q) = %v, expected %v", test.s, test.substr, result, test.expected)
		}
	}
}
