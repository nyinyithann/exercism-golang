// Package acronym includes a method to convert a sentence to an acronym
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a sentence to an acronym
func Abbreviate(s string) string {
	arr := strings.FieldsFunc(s, func(c rune) bool {
		if unicode.IsSpace(c) || string(c) == "-" {
			return true
		}
		return false
	})
	r := make([]byte, len(arr))
	for i, v := range arr {
		r[i] = v[0]
	}
	return strings.ToUpper(string(r))
}
