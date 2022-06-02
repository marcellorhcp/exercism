package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	w := []rune(strings.ToLower(word))
	lenght := len(word)

	for i := 0; i < lenght-1; i++ {
		for j := i + 1; j < lenght; j++ {
			if w[i] == w[j] && unicode.IsLetter(w[i]) {
				return false
			}
		}
	}
	return true
}
