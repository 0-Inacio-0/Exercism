package diffsquares

// SquareOfSum returns the square of the sum.
// (1+2+3+...+n)^2.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the square.
// (1^2+2^2+3^2+...+n^2)
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the square of the sum and the sum of the square.
// (1+2+3+...+n)^2 - (1^2+2^2+3^2+...+n^2)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
