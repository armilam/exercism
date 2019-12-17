// Package diffsquares contains functions to determine the
// sum of squares, square of sums, and difference between
// those two for the first n natural numbers.
package diffsquares

import (
	"math"
)

// SquareOfSum computes the square of the sum of the
// first n natural numbers.
func SquareOfSum(n int) int {
	return int(math.Pow(float64((n*(n+1))/2), 2))
}

// SumOfSquares computes the sum of the squares of the
// first n natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference computes the difference between the square
// of the sum and the sum of the squares of the first n
// natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
