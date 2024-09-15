package kata7

import (
	"math"
)

/*
#Beginner Series #3 Sum of Numbers

Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.

Note: a and b are not ordered!

Examples (a, b) --> output (explanation)
(1, 0) --> 1 (1 + 0 = 1)
(1, 2) --> 3 (1 + 2 = 3)
(0, 1) --> 1 (0 + 1 = 1)
(1, 1) --> 1 (1 since both are same)
(-1, 0) --> -1 (-1 + 0 = -1)
(-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)
Your function should only return a number, not the explanation about how you get that number.
*/
func GetSum(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		a, b = b, a
	}
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}

func GetSum1(a, b int) int {
	if a > b {
		a, b = b, a
	}
	return (a + b) * (b - a + 1) / 2
}

func GetSum2(a, b int) int {
	if a == b {
		return a
	}

	max := int(math.Max(float64(a), float64(b)))
	min := int(math.Min(float64(a), float64(b)))

	sum := 0
	for num := min; num <= max; num++ {
		sum += num
	}

	return sum
}
