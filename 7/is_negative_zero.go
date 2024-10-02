package kata7

import (
	"fmt"
	"math"
)

/*
Is It Negative Zero (-0)?
There exist two zeroes: +0 (or just 0) and -0.

Write a function that returns true if the input number is -0 and false otherwise (True and False for Python).
*/
func IsNegativeZero(n float64) bool {
	return n == 0 && math.Signbit(n)
}

func IsNegativeZero1(n float64) bool {
	return fmt.Sprintf("%v", n) == "-0"
}

func IsNegativeZero2(n float64) bool {
	return n == 0.0 && math.Copysign(1, n) == -1
}

func IsNegativeZero3(n float64) bool {
	return n == 0 && 1 / n < 0
}
