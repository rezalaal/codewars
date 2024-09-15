package kata8

import (
	"sort"
)

/*
#Merge two sorted arrays into one

You are given two sorted arrays that both only contain integers. Your task is to find a way to merge them into a single one, sorted in asc order. Complete the function mergeArrays(arr1, arr2), where arr1 and arr2 are the original sorted arrays.

You don't need to worry about validation, since arr1 and arr2 must be arrays with 0 or more Integers. If both arr1 and arr2 are empty, then just return an empty array.

Note: arr1 and arr2 may be sorted in different orders. Also arr1 and arr2 may have same integers. Remove duplicated in the returned result.

Examples (input -> output)
* [1, 2, 3, 4, 5], [6, 7, 8, 9, 10] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9], [10, 8, 6, 4, 2] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9, 11, 12], [1, 2, 3, 4, 5, 10, 12] -> [1, 2, 3, 4, 5, 7, 9, 10, 11, 12]
Happy coding!
*/
func MergeArrays(arr1, arr2 []int) []int {
	result := []int{}
	for _,n :=range arr1 {
		if !exists(result, n) {
			result = append(result, n)
		}
	}
	for _,n :=range arr2 {
		if !exists(result, n) {
			result = append(result, n)
		}
	}
	sort.Ints(result)
	return result
}

func exists(arr []int, n int) bool {
	for _, number :=range arr {
		if number == n {
			return true
		}
	}
	return false
}

func MergeArrays1(a, b []int) (result []int ){
	joined := append(a, b...)
	  sort.Ints(joined)
	  for i := 0; i < len(joined); i++ {
		  if i == 0 || joined[i] != joined[i-1] {
			  result = append(result, joined[i])
		  }
	  }
	  return
}

func MergeArrays2(arr1, arr2 []int) []int {

	amount := append(arr1, arr2...)
  sort.Ints(amount)
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range amount {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}