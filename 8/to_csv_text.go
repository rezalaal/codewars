package kata8

import (
	"fmt"
	"strconv"
	"strings"
)

/*
#CSV representation of array

Create a function that returns the CSV representation of a two-dimensional numeric array.

Example:

input:
   [[ 0, 1, 2, 3, 4 ],
    [ 10,11,12,13,14 ],
    [ 20,21,22,23,24 ],
    [ 30,31,32,33,34 ]] 
    
output:
     '0,1,2,3,4\n'
    +'10,11,12,13,14\n'
    +'20,21,22,23,24\n'
    +'30,31,32,33,34'
Array's length > 2.

More details here: https://en.wikipedia.org/wiki/Comma-separated_values

Note: you shouldn't escape the \n, it should work as a new line.
*/
func ToCsvText(array [][]int) string {
	array2 := []string{}
	for _ , a := range array {
	  k := []string{}
	  for _, n := range a {
		k = append(k, fmt.Sprint(n))
	  }
	  array2 = append(array2, strings.Join(k,","))
	}
	return strings.Join(array2, "\n")
  }
  
func ToCsvText1(array [][]int) string {
	res := ""
	for i, arr := range array {
		for j, item :=range arr {
			res += strconv.Itoa(item)
			if j != len(arr) - 1 { res += ","}
		}
		if i != len(arr) - 1 { res += "\n"}
	}
	res = strings.TrimRight(res, "\n")
	return res
}