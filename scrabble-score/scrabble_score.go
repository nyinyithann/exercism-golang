// Package scrabble includes Score method to calculate scrabble score
package scrabble

import "strings"

// Score calculate scrabble score
func Score(input string) int {
	s := 0
	for _, r := range input {
		switch strings.ToUpper(string(r)) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			s++
		case "D", "G":
			s += 2
		case "B", "C", "M", "P":
			s += 3
		case "F", "H", "V", "W", "Y":
			s += 4
		case "K":
			s += 5
		case "J", "X":
			s += 8
		case "Q", "Z":
			s += 10
		}
	}
	return s
}
