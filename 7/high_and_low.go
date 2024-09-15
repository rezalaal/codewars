package kata7

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

/*
#Highest and Lowest

In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

Examples
HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
Notes
All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.

*/
func HighAndLow(in string) string {
	numbers := strings.Split(in, " ")
	high := 0
	low := 0
	for i, n :=range numbers {
		if i == 0 {
			high, _ = strconv.Atoi(string(n))
			low, _ = strconv.Atoi(string(n))
		}else{
			current, _ := strconv.Atoi(string(n))
			if high < current { high = current}
			if low > current { low = current}
		}
	}
	return strconv.Itoa(high) + " " + strconv.Itoa(low)
}

func HighAndLow2(in string) string {
	numStrings := strings.Fields(in)
	var nums = []int{}
	
	for _, i := range numStrings {
	  j, _ := strconv.Atoi(i)
	  nums = append(nums, j)
	}
	sort.Ints(nums)
	return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}