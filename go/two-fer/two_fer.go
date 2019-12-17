// An exercism exercise to learn string formatting and conditionals.
package twofer

// Used for string formatting
import "fmt"

// Returns the string "One for you, one for me.".
// Replaces "you" with name if a name is given.
func ShareWith(name string) string {

	// If no name is given, use the word "you".
	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
