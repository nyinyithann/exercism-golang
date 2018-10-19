// Package diffsquares is for finding the difference between the square of sum and the sum of square of first N natural numbers
package diffsquares

// SquareOfSums calculates the square of sum of n numbers
func SquareOfSums(n int) int {
	s := (n * (n + 1)) / 2
	return s * s
}

// SumOfSquares calculates the sum of square of n numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference finds the difference between the square of sum and the sum of square of n numbers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
