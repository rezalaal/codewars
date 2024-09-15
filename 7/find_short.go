package kata7

import (
	"strings"
)

/*
#Shortest Word

Simple, given a string of words, return the length of the shortest word(s).

String will never be empty and you do not need to account for different data types.
*/
func FindShort(s string) int {
	words := strings.Split(s, " ")
	min := 0
	for i, word := range words {
		if i == 0 {
			min = len(word)
		} else {
			if min > len(word) {
				min = len(word)
			}
		}
	}
	return min
}

func FindShort1(s string) int {
	shortest := len(s)
	for _, word := range strings.Split(s, " ") {
	  if len(word) < shortest {
		shortest = len(word)
	  }
	}
	return shortest
}

func FindShort2(s string) int {
	var t int
	for i, s := range strings.Fields(s) {
		if len(s) < t || i == 0 {
			t = len(s)
		}
	}
	return t
}