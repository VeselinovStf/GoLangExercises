package hamming

import "errors"

const NotSameError = "hamming distance is only defined for sequences of equal length"

// Distance  compares two strands of DNA and count the differences between them.
// Returns error if sequences are not with same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New(NotSameError)
	}

	var count int

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
