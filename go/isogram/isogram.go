// Package isogram contains the solution to an exercism
// exercise which teaches sequences and strings.
package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram returns true if the given word is an isogram, false otherwise.
func IsIsogram(word string) bool {
	// Remove all non-letter characters.
	ignoredCharacters := regexp.MustCompile(`[^a-z]`)
	word = ignoredCharacters.ReplaceAllString(strings.ToLower(word), "")

	lettersUsed := map[byte]bool{}

	for i := 0; i < len(word); i++ {
		if lettersUsed[word[i]] {
			return false
		}

		lettersUsed[word[i]] = true
	}

	return true
}
