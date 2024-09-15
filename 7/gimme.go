package kata7

import (
	"slices"
	"sort"
)

/*
#Find the middle element

As a part of this Kata, you need to create a function that when provided with a triplet, returns the index of the numerical element that lies between the other two elements.

The input to the function will be an array of three distinct numbers (Haskell: a tuple).

For example:

gimme([2, 3, 1]) => 0
2 is the number that fits between 1 and 3 and the index of 2 in the input array is 0.

Another example (just to make sure it is clear):

gimme([5, 10, 14]) => 1
10 is the number that fits between 5 and 14 and the index of 10 in the input array is 1.
*/
func Gimme1(array [3]int) int {
	var newArray = make([]int, 3)
	copy(newArray, array[:])
	slices.Sort(newArray)
	for i, n := range array {
		if n == newArray[1] {
			return i
		}
	}
	return 0
}

func Gimme(array [3]int) int {
	min := array[0]
	middle := array[1]
	max := array[2]
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		
	}
	for i, n := range array {
		if n != min && n != max{
			middle = i
		}
	}
	return middle
}

func Gimme2(array [3]int) int {
	left, middle, right := array[0], array[1], array[2]
	if (left > middle && left < right) || (left < middle && left > right) {
	  return 0
	}
	if (middle > left && middle < right) || (middle < left && middle > right) {
	  return 1
	}
	return 2
}

func Gimme3(array [3]int) int {
	arrayCopy := array

	arrSlice := arrayCopy[:]
	sort.Ints(arrSlice)
	val := arrSlice[1]

	for i, v := range array {
		if v == val {
			return i
		}
	}

	return 0
}

func Gimme4(array [3]int) int {
	sorted := make([]int, len(array))
	copy(sorted, array[:])
	sort.Ints(sorted)
	for k, v := range array {
	  if v == sorted[1] {
		return k
	  }
	}
	return -1
}