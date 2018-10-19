// Package isogram is for checking if an input string is isogram or not
package isogram

import "strings"

// IsIsogram checks if an input string is isogram or not
func IsIsogram(input string) bool {
	m := make(map[string]struct{})
	v := struct{}{}
	c := 0
	input = strings.Replace(input, "-", "", -1)
	input = strings.Replace(input, " ", "", -1)
	for _, r := range strings.ToUpper(input) {
		m[string(r)] = v
		c++
	}
	if c == len(m) {
		return true
	}
	return false
}
