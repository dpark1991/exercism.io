// Package hamming is used to count the differences between DNA strands
package hamming

import "errors"

// Distance returns the differences between the DNA strands
func Distance(a, b string) (int, error) {
	count := 0

	// DNA strands must be the same length
	if len(a) != len(b) {
		return 0, errors.New("DNA strands are different lengths")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
