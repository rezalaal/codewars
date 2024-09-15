package kata7

import (
	"math"
)

/*
#Maximum Length Difference

You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z. Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).

Example:
a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13
Bash note:
input : 2 strings with substrings separated by ,
output: number as a string
*/
func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	maxA := 0
	minA := int(^uint(0) >> 1)
	for _, word := range a1 {
		if len(word) > maxA {
			maxA = len(word)
		}
		if len(word) < minA {
			minA = len(word)
		}
	}
	maxB := 0
	minB := int(^uint(0) >> 1)
	for _, word := range a2 {
		if len(word) > maxB {
			maxB = len(word)
		}
		if len(word) < minB {
			minB = len(word)
		}
	}
	diff1 := maxA - minB
	diff2 := maxB - minA
	if diff1 < 0 {
		diff1 = -diff1
	}
	if diff2 < 0 {
		diff2 = -diff2
	}

	if diff1 > diff2 {
		return diff1
	}
	return diff2
}

func findShortestAndLongest(a []string) (shortest, longest int) {
	lA := 0
	sA := math.MaxInt64
	for _, s := range a {
		if len(s) > lA {
			lA = len(s)
		}
		if len(s) < sA {
			sA = len(s)
		}
	}

	return sA, lA
}

func MxDifLg2(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	sA1, lA1 := findShortestAndLongest(a1)
	sA2, lA2 := findShortestAndLongest(a2)

	res1 := math.Abs(float64(sA1 - lA2))
	res2 := math.Abs(float64(sA2 - lA1))

	if res1 > res2 {
		return int(res1)
	}

	return int(res2)
}

func AbsDiff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

func MxDifLg3(a, b []string) int {
	max, cur := -1, 0
	for _, x := range a {
		for _, y := range b {
			cur = AbsDiff(len(x), len(y))
			if cur > max {
				max = cur
			}
		}
	}
	return max
}
