package kata8

import (
	"fmt"
	"strconv"
)

/*
#Sum Mixed Array

Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

Return your answer as a number.
*/
func SumMix(arr []any) int {
	sum := 0
	for _, v := range arr {
		switch v := v.(type) {
		case int:
			sum += v
		case string:
			num, _ := strconv.Atoi(v)
			sum += num
		}
	}
	return sum
}

func SumMix1(arr []any) int {
	sum := 0
	for _, v := range arr {
		iv, _ := strconv.Atoi(fmt.Sprintf("%v", v))
		sum += iv
	}
	return sum
}
