// Package raindrops implements string convertion from int to string
// The rules of raindrops are that if a given number:
// has 3 as a factor, add 'Pling' to the result.
// has 5 as a factor, add 'Plang' to the result.
// has 7 as a factor, add 'Plong' to the result.
// does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
package raindrops

import "strconv"

// Rules map of factor to result string
var rules = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Keys order of factors
var keys = []int{3, 5, 7}

// Convert takes a number and converts it into a string that contains
// raindrop sounds corresponding to certain potential factors
func Convert(n int) (raindrop string) {

	for _, v := range keys {
		if factor(n, v) {
			raindrop += rules[v]
		}
	}

	if raindrop != "" {
		return
	}

	return strconv.Itoa(n)
}

// Factor calculates number that evenly divides into another number, leaving no remainder
func factor(number int, factor int) (res bool) {
	if number%factor == 0 {
		return true
	}

	return
}
