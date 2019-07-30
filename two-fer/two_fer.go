// Package twofer completes the two fer exercise
package twofer

// ShareWith will return the correct "One for X, one for me." phrase
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
