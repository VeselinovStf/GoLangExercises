// Package collatzconjecture solves The Collatz Conjecture or 3x+1 problem :)
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1. by a given number n
func CollatzConjecture(n int) (s int, err error) {
	if n <= 0 {
		return 0, errors.New("zero is invalid start point")
	}

	for n != 1{
		if n % 2 == 0 {
			//even
			n = n / 2
		}else{
			//odd
			n = (n * 3) + 1
		}

		s++
	}

	return

}
