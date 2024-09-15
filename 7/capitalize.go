package kata7

import (
	"strings"
)

/*
#Alternate capitalization

Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.

For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.

The input will be a lowercase string with no spaces.

Good luck!
*/
func Capitalize(st string) []string {
	part1 := ""
	part2 := ""
	for i,char :=range st {
		if i % 2 == 0 {
			part1 += strings.ToLower(string(char))
			part2 += strings.ToUpper(string(char))
		}else{
			part2 += strings.ToLower(string(char))
			part1 += strings.ToUpper(string(char))
		}
	}
	return []string{part2, part1}
}