package kata7

import (
	"math"
)

/*
#Find the next perfect square!

You might know some pretty large perfect squares. But what about the NEXT one?

Complete the findNextSquare method that finds the next integral perfect square after the one passed as a parameter. Recall that an integral perfect square is an integer n such that sqrt(n) is also an integer.

If the argument is itself not a perfect square then return either -1 or an empty value like None or null, depending on your language. You may assume the argument is non-negative.

Examples ( Input --> Output )
121 --> 144
625 --> 676
114 --> -1  #  because 114 is not a perfect square
*/
func FindNextSquare(sq int64) int64 {
	m := math.Round(math.Sqrt(float64(sq)))
	
	if math.Pow(m, 2) == float64(sq) {
		return int64(math.Pow(m+1, 2))
	}
	return -1
}

func FindNextSquare1(sq int64) int64 {
	res := math.Pow(math.Sqrt(float64(sq)) + 1, 2)
	
	if res == math.Trunc(res) {
	  return int64(res)
	}
	return -1
}

func FindNextSquare2(sq int64) int64 {
	sqrt := int64(math.Sqrt(float64(sq)))
	if sq % sqrt != 0 {
	  return -1
	}
	sqrt++
	return sqrt * sqrt
}

func FindNextSquare3(sq int64) int64 {
	res := math.Sqrt(float64(sq))
  
	if res != math.Ceil(res) {
	  return -1
	}
	
	return (int64(res) + 1) * (int64(res) + 1)
}

func FindNextSquare4(sq int64) int64 {
	root := int64(math.Sqrt(float64(sq)))
	
	if root * root == sq {
	  next := root + 1
	  return next * next
	}
	
	return -1
}

func FindNextSquare5(sq int64) int64 {
	sqrt := math.Sqrt(float64(sq))
	if sqrt != math.Round(sqrt) {
		return -1
	}
	return int64(math.Pow(sqrt+1, 2))
}