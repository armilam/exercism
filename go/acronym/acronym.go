// Package acronym is an exercism exercise to learn string manipulation.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the acronym for the given phrase.
func Abbreviate(s string) string {
	// Replace all non-word/whitespace characters with a space.
	ignoredCharacters := regexp.MustCompile(`[^A-Za-z'\s]`)
	s = ignoredCharacters.ReplaceAllString(s, " ")

	// Split the string on whitespace.
	whitespace := regexp.MustCompile(`\s+`)
	words := whitespace.Split(s, -1)

	// Collect the first letter of each word.
	acronym := ""
	for _, word := range words {
		acronym = acronym + string(word[0])
	}

	// Return the acronym in upper case.
	return strings.ToUpper(acronym)
}
