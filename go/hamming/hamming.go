// Package hamming is the solution to an exercism exercise to learn
// about strings and iteration.
package hamming

import (
	"errors"
)

// Distance calculates the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strands are not the same length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
