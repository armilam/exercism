// An exercism exercise to learn basic math and conditionals
package leap

// Returns true if the given year is a leap year, false otherwise.
func IsLeapYear(year int) bool {
	return (year % 4 == 0) && (year % 100 != 0 || year % 400 == 0)
}
