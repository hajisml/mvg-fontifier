package transformer

import "strings"

// customAlphabet is a map for substitutions (uppercase and lowercase combined)
var customAlphabet = map[rune]rune{
	'A': 'Δ', 'D': 'Đ', 'E': 'Σ', 'F': 'Ϝ',
	'J': 'J', 'N': 'Ͷ', 'O': 'Ø', 'S': 'Ϟ',
	'T': 'Ͳ', 'a': 'α', 'c': 'ϲ', 'd': 'ɗ',
	'f': 'ϝ', 'j': 'ʝ', 'm': 'ϻ', 'n': 'η',
	'o': 'o', 's': 'ϟ', 'u': 'υ', 'v': 'v',
	'z': 'ʐ',
}

// Transform takes a string and returns it with characters substituted based on the custom alphabet.
func Transform(input string) string {
	var result strings.Builder
	for _, char := range input {
		if replacement, exists := customAlphabet[char]; exists {
			result.WriteRune(replacement)
		} else {
			// Keep original character if no substitution exists
			result.WriteRune(char)
		}
	}
	return result.String()
}
