package kata7

import (
	"sort"
)

func MinMax(arr []int) [2]int {
	var res = [2]int{arr[0], arr[0]}
	
	for _, n :=range arr {
		if n > res[1] { res[1] = n}
		if n < res[0] { res[0] = n}
	}
	return res
}

func MinMax1(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}

func MinMax2(a []int) [2]int {
	min, max := a[0], a[0]
	for _,v := range a[1:] {
	  if v<min {min=v}else if v>max {max=v}
	}
	return [2]int{min, max}
}