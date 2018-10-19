// Package hamming implements a function to calculate the Hamming difference between two DNA strands
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance compares two string parameters and returns the number of differences between them.
// If the parameters have different length, zero and error will return.
func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("the length of two parameters must be the same")
	}
	count := 0
	for _, ar := range a {
		br, w := utf8.DecodeRuneInString(b[0:])
		b = b[w:]
		if ar != br {
			count++
		}
	}
	return count, nil
}
