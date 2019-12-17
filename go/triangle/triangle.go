// Package triangle is a solution to an exercism exercise which teaches
// conditionals and boolean logic.
package triangle

import "math"

// Kind represents a triangle type.
type Kind string

const (
	// NaT represents a triangle with side lengths that are not possible to make a real triangle
	NaT = Kind("not a triangle")
	// Equ represents a triangle with three equal sides
	Equ = Kind("equilateral")
	// Iso represents a triangle with two equal sides
	Iso = Kind("isosceles")
	// Sca represents a valid triangle with no equal sides
	Sca = Kind("scalene")
)

// KindFromSides returns a Kind representing the triangle defined by the three side lengths.
func KindFromSides(a, b, c float64) Kind {
	if SidesAreInvalid(a, b, c) {
		return NaT
	} else if a+b < c || a+c < b || b+c < a {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else {
		return Sca
	}
}

// SidesAreInvalid returns true if any side is not a number, infinite, or not positive.
func SidesAreInvalid(a, b, c float64) bool {
	return math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) ||
		a <= 0 || b <= 0 || c <= 0
}
