/*
Package twofer is the solution for the two-fer exercism exercise.
*/
package twofer

import "fmt"

// ShareWith returns a string in the pattern "One for X, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
