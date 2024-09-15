package kata7

import "math"

/*
Return the closest number multiple of 10
Given a number return the closest number to it that is divisible by 10.

Example input:

22
25
37
Expected output:

20
30
40
*/
func ClosestMultipleOf10(n uint32) uint32 {
	if n%10 < 5 {
		return (n / 10) * 10
	}
	return ((n / 10) + 1) * 10
}
func ClosestMultipleOf102(n uint32) uint32 {
	return (n + 5) / 10 * 10
}

func ClosestMultipleOf103(n uint32) uint32 {
	lastDigit := n % 10
	if lastDigit >= 5 {
		return n + 10 - lastDigit
	}
	return n - lastDigit
}

func ClosestMultipleOf104(n uint32) uint32 {
	return uint32(math.Round(float64(n)/10.) * 10)
}