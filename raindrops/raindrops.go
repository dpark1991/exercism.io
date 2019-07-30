// Package raindrops completes the raindrops exercise
package raindrops

import (
	"strconv"
)

// Convert will return a string based on the input.
// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
func Convert(in int) string {
	var out string
	if in%3 == 0 {
		out += "Pling"
	}
	if in%5 == 0 {
		out += "Plang"
	}
	if in%7 == 0 {
		out += "Plong"
	}
	if len(out) == 0 {
		out = strconv.Itoa(in)
	}
	return out
}
