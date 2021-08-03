// Package scrabble implements Scrable game functions
package scrabble

import "strings"

// Score computes the Scrabble score from given word.
func Score(s string) (res int) {

	for _, r := range strings.ToUpper(s) {
		switch r {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			res++
		case 'D', 'G':
			res += 2
		case 'B', 'C', 'M', 'P':
			res += 3
		case 'F', 'H', 'V', 'W', 'Y':
			res += 4
		case 'K':
			res += 5
		case 'J', 'X':
			res += 8
		case 'Q', 'Z':
			res += 10
		}
	}

	return
}
