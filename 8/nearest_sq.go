package kata8

import (
	"math"
)

/*
#Find Nearest square number
Your task is to find the nearest square number, nearest_sq(n) or nearestSq(n), of a positive integer n.

For example, if n = 111, then nearest\_sq(n) (nearestSq(n)) equals 121, since 111 is closer to 121, the square of 11, than 100, the square of 10.

If the n is already the perfect square (e.g. n = 144, n = 81, etc.), you need to just return n.

Good luck :)

Check my other katas:

Alphabetically ordered

Case-sensitive!

Not prime numbers

Find your caterer
*/
func NearestSq(n int) int {
	// Calculate the integer square root of n
	sqrt := int(math.Sqrt(float64(n)))

	// Calculate the nearest perfect squares
	lowerSquare := sqrt * sqrt
	upperSquare := (sqrt + 1) * (sqrt + 1)

	// Check if n is already a perfect square
	if lowerSquare == n {
		return n
	}

	// Determine which square is closer
	if (n - lowerSquare) < (upperSquare - n) {
		return lowerSquare
	}
	return upperSquare
}

func NearestSq1(n int) int {
	n = int(math.Round(math.Sqrt(float64(n))))

	return n * n
}