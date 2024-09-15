package kata5

import (
	"math"
)

/*
#Maximum subarray sum

The maximum sum subarray problem consists in finding the maximum sum of a contiguous subsequence in an array or list of integers:

maxSequence [-2, 1, -3, 4, -1, 2, 1, -5, 4]
-- should be 6: [4, -1, 2, 1]
Easy case is when the list is made up of only positive numbers and the maximum sum is the sum of the whole array. If the list is made up of only negative numbers, return 0 instead.

Empty list is considered to have zero greatest sum. Note that the empty list or array is also a valid sublist/subarray.
*/
func MaximumSubarraySum(numbers []int) int {
	max := 0
	currentMax := 0

	for _, num := range numbers {
		currentMax += num
		if currentMax < 0 {
			currentMax = 0			 
		}
		if currentMax > max {
			max = currentMax 
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaximumSubarraySum2(numbers []int) int {
	res, sum := 0, 0
	for i := range numbers {
		sum += numbers[i]
		res = max(res, sum)
		sum = max(sum, 0)
	}
	return res
}

func MaximumSubarraySum3(numbers []int) int {
  
	max, curr := 0, 0
	  for _, e := range numbers {
		  if curr += e; curr < 0 {
			  curr = 0
		  } else if curr > max {
			  max = curr
		  }
	  }
	  return max
}

func MaximumSubarraySum4(numbers []int) int {
	var min, ans, sum = 0, 0, 0
	for _, val := range numbers {
		sum += val
		min = int(math.Min(float64(sum), float64(min)))
		ans = int(math.Max(float64(ans), float64(sum-min)))
	}
	return ans
}