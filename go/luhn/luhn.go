package luhn

import (
	"strings"
)

// Valid returns true if the number is a valid Luhn number, false otherwise.
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")

	return ValidLength(number) && ValidChecksum(number)
}

// ValidLength returns true if the number is long enough to be a Luhn number, false otherwise.
func ValidLength(number string) bool {
	return len(number) > 1
}

// ValidChecksum returns true if the number has a valid checksum using the Luhn algorithm,
// false otherwise.
func ValidChecksum(number string) bool {
	sum := 0
	length := len(number)

	for i := 0; i < length; i++ {
		// Turn the digit character into a numerical value
		digit := int(number[i] - '0')

		// If the digit value is outside the range 0-9, it's an invalid character
		if digit < 0 || digit > 9 {
			return false
		}

		// If this digit should be doubled, double it
		if i % 2 == length % 2 {
			digit *= 2
		}

		// If the digit is multi-digit, subtract 9
		if digit > 9 {
			digit -= 9
		}

		sum += digit
	}

	return sum % 10 == 0
}
