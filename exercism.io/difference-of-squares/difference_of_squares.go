// Package diffsquares implements calculatios for diffs of squares
package diffsquares

// SumOfSquares finds the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference find difference between the square of
// the sum of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareOfSum find sum of square of the first N natural numbers.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}
