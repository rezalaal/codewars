package kata7

import (
	"sort"
)

/*
#Sort Numbers

Finish the solution so that it sorts the passed in array of numbers. If the function passes in an empty array or null/nil value then it should return an empty array.

For example:

solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
solution(NULL)              # should return NULL
*/
func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func SortNumbers1(numbers []int) []int {
	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Ints(sortedNumbers)
	return sortedNumbers
}


func SortNumbers2(numbers []int) []int {
  
	for i := 0; i < len(numbers); i++{
	  for j := 0; j < len(numbers); j++ {
		if numbers[j] > numbers[i] {
		  temp := numbers[i]
		  numbers[i] = numbers[j]
		  numbers[j] = temp
		}
	  }
	  
	}
	
	return numbers // your code here
}