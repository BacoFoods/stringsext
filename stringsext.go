// Package stringsext provides extended string manipulation functions
// that are not available in the standard strings package.
package stringsext

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Capitalize returns a copy of the string s with the first letter capitalized
// using the golang.org/x/text/cases package.
// If the string is empty, it returns an empty string.
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	// Create a title caser for English language
	caser := cases.Title(language.English)

	return caser.String(s)
}

// CapitalizeWords returns a copy of the string s with the first letter of each word capitalized.
// If the string is empty, it returns an empty string.
func CapitalizeWords(s string) string {
	if len(s) == 0 {
		return s
	}

	words := strings.Fields(s)
	for i, word := range words {
		if len(word) == 0 {
			continue
		}

		// Capitalize the first letter and make the rest lowercase
		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " ")
}
