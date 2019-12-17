// An exercism exercise to learn string parsing and conditionals
package bob

import (
	"strings"
	"regexp"
)

// Reponds to the input sentence with simple rules
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	if isEmpty(trimmedRemark) {
		return "Fine. Be that way!"
	} else if isQuestion(trimmedRemark) && isYelling(trimmedRemark) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(trimmedRemark) {
		return "Sure."
	} else if isYelling(trimmedRemark) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

// Returns true if the string is empty, falses otherwise
func isEmpty(remark string) bool {
	return len(remark) == 0
}

// Returns true if the string ends with a '?' character, false otherwise
func isQuestion(remark string) bool {
	return remark[len(remark) - 1] == '?'
}

// Returns if the string contains at least one letter and all letters
// are capitalized, false otherwise
func isYelling(remark string) bool {
	matched, _ := regexp.Match(`[A-Za-z]`, []byte(remark))
	return matched && strings.ToUpper(remark) == remark
}

