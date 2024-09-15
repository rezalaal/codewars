package kata7

import (
	"sort"
)

/*
#Sort array by string length

Write a function that takes an array of strings as an argument and returns a sorted array containing the same strings, 
ordered from shortest to longest.

For example, if this array were passed as an argument:

["Telescopes", "Glasses", "Eyes", "Monocles"]

Your function would return the following array:

["Eyes", "Glasses", "Monocles", "Telescopes"]

All of the strings in the array passed to your function will be different lengths, 
so you will not have to decide how to order multiple strings of the same length.
*/
func SortByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	return arr
}

func SortByLength1(arr []string) []string {
  
    for i := 0; i < len(arr); i++{
        for j := i; j < len(arr); j++{
            if len(arr[j]) < len(arr[i]){
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    return arr

}

func SortByLength2(arr []string) []string {
	if len(arr) < 2 {
	  return arr
  }

  pivot := arr[0]
  var left, right []string
  for i := 1; i < len(arr); i++ {
	  if len(pivot) > len(arr[i]) {
		  left = append(left, arr[i])
	  } else {
		  right = append(right, arr[i])
	  }

  }
  result := append(SortByLength(left), pivot)
  result = append(result, SortByLength(right)...)
  return result
}