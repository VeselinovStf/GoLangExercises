// Package isogram implements isogram functions
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter,
// however spaces and hyphens are allowed to appear multiple times.
package isogram

import "unicode"

// IsIsogram Determine if a word or phrase is an isogram.
func IsIsogram(input string) (result bool) {	
	// Store and count each string rune
	contents := make(map[rune]int)

	for _, r := range input {		
		// Format input string
		r = unicode.ToLower(r)
		// spaces and hyphens are allowed to appear multiple times
		if r != ' ' && r != '-' {
			// if current rune count is to return false
			if contents[r]+1 == 2 {
				return
			}

			contents[r]++
		}
	}

	// string is isogram
	return true
}
