package isogram

import "strings"

func IsIsogram(word string) bool {
	registeredChars := make(map[int32]struct{})
	for _, char := range strings.ToLower(word) {
		if _, exists := registeredChars[char]; !exists {
			if char != int32('-') && char != int32(' ') {
				registeredChars[char] = struct{}{}
			}
		} else {
			return false
		}
	}
	return true
}
