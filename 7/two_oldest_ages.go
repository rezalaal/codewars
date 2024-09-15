package kata7

import (
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	if len(ages) < 2 {
		return [2]int{}
	}
	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}

func TwoOldestAges1(ages []int) [2]int {
	a, b := 0, 0
	for _, v := range ages {
	  if v > b {
		a, b = b, v
	  } else if v > a {
		a = v
	  }
	}
	return [2]int{a, b}
}

func TwoOldestAges2(ages []int) [2]int {
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))
	return [2]int{ages[1],ages[0]}
}