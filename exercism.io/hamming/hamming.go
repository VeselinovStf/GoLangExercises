/*
	Package Hamming implements simple methods for finding hamming distance
*/
package hamming

import "errors"

const notSameError = "hamming distance is only defined for sequences of equal length"

// Distance  compares two strands of DNA and count the differences between them.
// Returns error if sequences are not with same length
func Distance(a, b string) (count int, err error) {
	if len(a) != len(b) {
		return count, errors.New(notSameError)
	}

	for i, char := range a {
		if char != rune(b[i]) {
			count++
		}
	}

	return
}
