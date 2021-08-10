// Package diffsquares implements calculatios for diffs of squares
package diffsquares

var sumOfSquares = func(n int)int{
	return  n*(n+1)*(2*n+1) / 6
}

var squareOfSum = func(n int) int{
	eq := n*(n+1) / 2 
	return eq * eq
}

// SumOfSquares finds the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int{
	return  sumOfSquares(n)
}

// Difference find difference between the square of
// the sum of the first N natural numbers.
func Difference(n int) int{
	return  squareOfSum(n) - sumOfSquares(n)
}

// SquareOfSum find sum of square of the first N natural numbers.
func SquareOfSum(n int) int{
	return squareOfSum(n)
}