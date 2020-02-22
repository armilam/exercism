package grains

import (
	"math"
	"fmt"
)

// Square returns the number of grains on the given square.
func Square(square int) (uint64, error) {
	if square < 1 {
		return 0, fmt.Errorf("Square %d is less than 1", square)
	} else if square > 64 {
		return 0, fmt.Errorf("Square %d is greater than 64", square)
	} else {
		return uint64(math.Pow(2, float64(square - 1))), nil
	}
}

// Total returns the total number of grains on the chess board.
func Total() uint64 {
	var sum uint64 = 0

	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		sum += square
	}

	return sum
}
