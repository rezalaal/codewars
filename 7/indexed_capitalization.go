package kata7

import (
	"strings"
)

/*
Indexed capitalization

Given a string and an array of integers representing indices, capitalize all letters at the given indices.

For example:

capitalize("abcdef",[1,2,5]) = "aBCdeF"
capitalize("abcdef",[1,2,5,100]) = "aBCdeF". There is no index 100.
The input will be a lowercase string with no spaces and an array of digits.

Good luck!
*/
func IndexedCapitalization(st string, arr []int) string {
	items := map[int]string{}
	res := ""
	for i, char := range st {
		items[i] = string(char)
	}
	for _, index := range arr {
		if index < len(st) && index >=0 {
			items[index] = strings.ToUpper(items[index])
		}
	}

	for i := 0; i < len(items); i++ {
		res += items[i]
	}
	return res
}
func IndexedCapitalization2(st string, arr []int) string {
	for _, v := range arr {
	  if v < len(st) {
		st = st[:v] + strings.ToUpper(string(st[v])) + st[v+1:]
	  }
	}
	
	return st
  }