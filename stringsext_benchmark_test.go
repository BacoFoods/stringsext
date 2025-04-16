package stringsext

import (
	"testing"
)

func BenchmarkCapitalizeWords(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"SingleWord", "hello"},
		{"TwoWords", "hello world"},
		{"LongPhrase", "this is a longer phrase with multiple words to capitalize"},
		{"WithHyphens", "hyphenated-words-in-this-string-for-testing"},
		{"WithUnderscores", "underscore_separated_words_for_special_case"},
		{"MixedDelimiters", "words with spaces, tabs\tand-hyphens mixed in"},
		{"VeryLongString", "this is a very long string that is used to test the performance of the capitalize function when dealing with inputs that contain many characters and therefore require more processing time which might reveal performance bottlenecks in the implementation that might not be obvious with shorter strings"},
		{"SpecialCharacters", "!@#$%^&*()_+[]{}|;':\",.<>?"},
		{"NumbersAndSymbols", "12345-67890"},
		{"LeadingSpaces", "   leading spaces"},
		{"TrailingSpaces", "trailing spaces   "},
		{"MixedCase", "mIXed CaSe"},
		{"EmptySpaces", "     "},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CapitalizeWords(tc.input)
			}
		})
	}
}

func BenchmarkCapitalize(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"SingleWord", "hello"},
		{"TwoWords", "hello world"},
		{"LongPhrase", "this is a longer phrase with multiple words to capitalize"},
		{"WithHyphens", "hyphenated-words-in-this-string-for-testing"},
		{"WithUnderscores", "underscore_separated_words_for_special_case"},
		{"MixedDelimiters", "words with spaces, tabs\tand-hyphens mixed in"},
		{"VeryLongString", "this is a very long string that is used to test the performance of the capitalize function when dealing with inputs that contain many characters and therefore require more processing time which might reveal performance bottlenecks in the implementation that might not be obvious with shorter strings"},
		{"SpecialCharacters", "!@#$%^&*()_+[]{}|;':\",.<>?"},
		{"NumbersAndSymbols", "12345-67890"},
		{"LeadingSpaces", "   leading spaces"},
		{"TrailingSpaces", "trailing spaces   "},
		{"MixedCase", "mIXed CaSe"},
		{"EmptySpaces", "     "},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Capitalize(tc.input)
			}
		})
	}
}

// BenchmarkCapitalizeWordsParallel tests the CapitalizeWords function in a parallel context
func BenchmarkCapitalizeWordsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CapitalizeWords("this is a standard test string for parallel benchmark")
		}
	})
}

// BenchmarkCapitalizeParallel tests the Capitalize function in a parallel context
func BenchmarkCapitalizeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Capitalize("this is a standard test string for parallel benchmark")
		}
	})
}
