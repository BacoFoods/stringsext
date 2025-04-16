# stringsext

A Go library that provides extended string helpers for easier string manipulation.

## Installation

```bash
go get github.com/BacoFoods/stringsext
```

## Usage

```go
import "github.com/BacoFoods/stringsext"
```

## Functions

### `Capitalize(s string) string`

Returns a copy of the string with the first letter of each word capitalized, using the `golang.org/x/text/cases` package. This provides standardized, locale-aware text handling which is valuable for internationalization.

```go
stringsext.Capitalize("hello world") // Returns "Hello World"
```

### `CapitalizeWords(s string) string`

Returns a copy of the string with the first letter of each word capitalized and the rest of the letters in lowercase. This function is optimized for speed and works by splitting the string by spaces, capitalizing each word, and joining them back together.

```go
stringsext.CapitalizeWords("hello WORLD") // Returns "Hello World"
```

## Behavior Differences

The two capitalization functions have slightly different behaviors in certain cases:

```go
// Test cases showing the differences between Capitalize and CapitalizeWords:

// String starting with numbers
Capitalize("123abc") // Returns "123Abc"
CapitalizeWords("123abc") // Returns "123abc"

// Multiple spaces
CapitalizeWords("multiple   spaces   here") // Returns "Multiple Spaces Here"

// Tab separated
CapitalizeWords("tab\tseparated\twords") // Returns "Tab Separated Words"

// Newline characters
CapitalizeWords("new\nline characters") // Returns "New Line Characters"

// Hyphenated words
CapitalizeWords("hyphenated-word") // Returns "Hyphenated-word"

// Underscore separated
CapitalizeWords("underscore_separated_words") // Returns "Underscore_separated_words"
```

## Performance

The `CapitalizeWords` function is generally faster but has more allocations for complex strings, while `Capitalize` using the standard library has fewer but larger allocations. Handling of space characters and word definitions also changes. So choose accordingly.

```code
goos: darwin
goarch: arm64
pkg: github.com/BacoFoods/stringsext
cpu: Apple M1 Max
BenchmarkCapitalizeWords/Empty-10       537174226                2.056 ns/op           0 B/op          0 allocs/op
BenchmarkCapitalizeWords/SingleWord-10  15902100                76.37 ns/op           32 B/op          3 allocs/op
BenchmarkCapitalizeWords/TwoWords-10     7288099               164.2 ns/op            80 B/op          6 allocs/op
BenchmarkCapitalizeWords/LongPhrase-10   1834941               642.0 ns/op           384 B/op         21 allocs/op
BenchmarkCapitalizeWords/WithHyphens-10                  7710529               154.7 ns/op            72 B/op          3 allocs/op
BenchmarkCapitalizeWords/WithUnderscores-10              8021689               149.2 ns/op            72 B/op          3 allocs/op
BenchmarkCapitalizeWords/MixedDelimiters-10              2372905               492.6 ns/op           280 B/op         16 allocs/op
BenchmarkCapitalizeWords/VeryLongString-10                372403              3147 ns/op            1896 B/op         95 allocs/op
BenchmarkCapitalizeWords/SpecialCharacters-10           10914187               109.6 ns/op            48 B/op          2 allocs/op
BenchmarkCapitalizeWords/NumbersAndSymbols-10           15520041                75.16 ns/op           32 B/op          2 allocs/op
BenchmarkCapitalizeWords/LeadingSpaces-10                7193142               168.3 ns/op            80 B/op          6 allocs/op
BenchmarkCapitalizeWords/TrailingSpaces-10               7026770               170.2 ns/op            80 B/op          6 allocs/op
BenchmarkCapitalizeWords/MixedCase-10                    6271208               193.2 ns/op            88 B/op          7 allocs/op
BenchmarkCapitalizeWords/EmptySpaces-10                 100000000               11.69 ns/op            0 B/op          0 allocs/op
BenchmarkCapitalize/Empty-10                            575325256                2.073 ns/op           0 B/op          0 allocs/op
BenchmarkCapitalize/SingleWord-10                        3663526               328.6 ns/op           437 B/op          4 allocs/op
BenchmarkCapitalize/TwoWords-10                          3095670               410.2 ns/op           448 B/op          4 allocs/op
BenchmarkCapitalize/LongPhrase-10                        1334900               895.1 ns/op           496 B/op          4 allocs/op
BenchmarkCapitalize/WithHyphens-10                       1629960               741.9 ns/op           480 B/op          4 allocs/op
BenchmarkCapitalize/WithUnderscores-10                   1664733               744.3 ns/op           480 B/op          4 allocs/op
BenchmarkCapitalize/MixedDelimiters-10                   1564686               758.7 ns/op           480 B/op          4 allocs/op
BenchmarkCapitalize/VeryLongString-10                     323991              3662 ns/op            1520 B/op          6 allocs/op
BenchmarkCapitalize/SpecialCharacters-10                 2644081               455.7 ns/op           432 B/op          3 allocs/op
BenchmarkCapitalize/NumbersAndSymbols-10                 3541042               344.1 ns/op           432 B/op          3 allocs/op
BenchmarkCapitalize/LeadingSpaces-10                     2618859               461.1 ns/op           456 B/op          4 allocs/op
BenchmarkCapitalize/TrailingSpaces-10                    2482256               473.6 ns/op           456 B/op          4 allocs/op
BenchmarkCapitalize/MixedCase-10                         3074365               379.6 ns/op           448 B/op          4 allocs/op
BenchmarkCapitalize/EmptySpaces-10                       4043048               304.1 ns/op           432 B/op          3 allocs/op
BenchmarkCapitalizeWordsParallel-10                      7287201               163.3 ns/op           352 B/op         19 allocs/op
BenchmarkCapitalizeParallel-10                           5422063               215.6 ns/op           496 B/op          4 allocs/op
```