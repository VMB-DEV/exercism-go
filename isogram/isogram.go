package isogram

import "strings"

func IsIsogram(word string) bool {
	registeredChars := make(map[int32]struct{})
	for _, char := range strings.ToLower(word) {
		countThisChar := char != '-' && char != ' '
		if countThisChar {
			if _, exists := registeredChars[char]; !exists {
				registeredChars[char] = struct{}{}
			} else {
				return false
			}
		}
	}
	return true
}
