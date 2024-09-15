package kata8


import (
	"sort"
)

/*
#Find the smallest integer in the array

Given an array of integers your solution should find the smallest integer.

For example:

Given [34, 15, 88, 2] your solution will return 2
Given [34, -345, -1, 100] your solution will return -345
You can assume, for the purpose of this kata, that the supplied array will not be empty.
*/
func SmallestIntegerFinder(numbers []int) int {
	min := numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}

func SmallestIntegerFinder1(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}