// Package raindrops contains the solution to an exercism exercise
// that teaches conditions, filtering, and strings.
package raindrops

import (
	"strconv"
)

// Convert turns an integer into a string of rain drop
// sounds as a string according to its factorization.
// If the number does not convert to rain drop sounds,
// the the number is returned as a string.
func Convert(number int) string {
	sounds := ""

	if number%3 == 0 {
		sounds += "Pling"
	}

	if number%5 == 0 {
		sounds += "Plang"
	}

	if number%7 == 0 {
		sounds += "Plong"
	}

	if len(sounds) == 0 {
		sounds = strconv.Itoa(number)
	}

	return sounds
}
