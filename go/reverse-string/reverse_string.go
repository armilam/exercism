// Package reverse provides a function which returns the reverse of strings.
package reverse

import (
	"strings"
)

// Reverse reverses the given string.
func Reverse(s string) string {
	reversed := ""
	reader := strings.NewReader(s)

	// Loop over the runes and prepend them to the reversed string.
	for reader.Len() > 0 {
		ch, _, _ := reader.ReadRune()
		reversed = string(ch) + reversed
	}

	return reversed
}
