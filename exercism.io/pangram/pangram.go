// Package pangram implements pangram functions
package pangram

import (

	"unicode"
)

// IsPangram determines if a sentence is a pangram.
// A pangram (Greek: παν γράμμα, pan gramma, "every letter")
//is a sentence using every letter of the alphabet at least once.
func IsPangram(s string) bool{
	l := make(map[rune]int)
	for _,r := range s{
		f := unicode.ToLower(r)
		if ok := l[f] == 0 && !unicode.IsSpace(f) && !unicode.IsPunct(f) && !unicode.IsNumber(f); ok {
			l[f]++
		}
	}

	return len(l) == 26
}