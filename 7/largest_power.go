package kata7

import (
	"math"
)

/*
Powers of 3

Given a positive integer N, return the largest integer k such that 3^k < N.

For example,

LargestPower(3) // returns 0
LargestPower(4) // returns 1
You may assume that the input to your function is always a positive integer.
*/
func LargestPower(n uint64) int {
	
	i := 0
	for  {
		if math.Pow(3, float64(i)) >= float64(n) {
			return i - 1
		}

		i++
	}
}

func LargestPower1(n uint64) int {
    k := 0.0
    for math.Pow(3,k) < float64(n) {
        k++
    }
    return int(k-1)
}

func LargestPower2(n uint64) int {
	res := int(math.Log(float64(n))/math.Log(3))
	if uint64(math.Pow(3, float64(res))) < n { return res }
	return res-1
}
