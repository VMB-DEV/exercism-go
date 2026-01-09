package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	registeredChars := make(map[rune]struct{})
	for _, char := range strings.ToLower(word) {
		if unicode.IsLetter(char) {
			if _, exists := registeredChars[char]; !exists {
				registeredChars[char] = struct{}{}
			} else {
				return false
			}
		}
	}
	return true
}
