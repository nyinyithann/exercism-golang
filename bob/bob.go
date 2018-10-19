// Package bob describes about a lackadaisical teenager called Bob
package bob

import (
	"strings"
	"unicode"
)

// Hey describes the way Bob answers particular questions
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if strings.HasSuffix(remark, "?") {
		rk := remark[:len(remark)-1]
		if strings.ToUpper(remark) == remark && !areWordsNumbers(rk) && !areWordsPuncts(rk) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if strings.ToUpper(remark) == remark && !areWordsNumbers(remark) {
		return "Whoa, chill out!"
	} else if strings.TrimSpace(remark) == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func areWordsNumbers(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			continue
		}
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func areWordsPuncts(s string) bool {
	for _, r := range s {

		if unicode.IsSpace(r) {
			continue
		}
		if !unicode.IsPunct(r) {
			return false
		}
	}
	return true
}
