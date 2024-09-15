package kata8

import (
	"math"
)

/*
#To square(root) or not to square(root)

Write a method, that will get an integer array as parameter and will process every number from this array.

Return a new array with processing every number of the input-array like this:

If the number has an integer square root, take this, otherwise square the number.

Example
[4,3,9,7,2,1] -> [2,9,3,49,4,1]
Notes
The input array will always contain only positive numbers, and will never be empty or null.
*/
func SquareOrSquareRoot(arr []int) []int {
	var result = make([]int, len(arr))
	for i, n := range arr {
		sqrt := math.Sqrt(float64(n))
		if sqrt == float64(int(sqrt)) {
			result[i] = int(sqrt)
		} else {
			result[i] = int(math.Pow(float64(n), 2))
		}
	}
	return result
}

func SquareOrSquareRoot1(arr []int) []int {
	results := make([]int, len(arr))

	for i, x := range arr {
		sqrt := math.Sqrt(float64(x))

		if sqrt == math.Floor(sqrt) {
			results[i] = int(sqrt)
		} else {
			results[i] = x * x
		}
	}

	return results
}

func SquareOrSquareRoot2(arr []int) []int {
	var r []int
	for _, num := range arr {
		s := math.Sqrt(float64(num))
		if math.Mod(s, 1.0) == 0 {
			r = append(r, int(s))
		} else {
			r = append(r, num*num)
		}
	}

	return r
}

func SquareOrSquareRoot3(arr []int) []int {
	o := make([]int, len(arr))
	for i, v := range arr {
		s := int(math.Sqrt(float64(v)))
		if v == s*s {
			o[i] = s
		} else {
			o[i] = v * v
		}
	}
	return o
}
