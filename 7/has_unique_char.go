package kata7

import (
	"strings"
)

/*
#All unique

Write a program to determine if a string contains only unique characters. Return true if it does and false otherwise.

The string may contain any of the 128 ASCII characters. Characters are case-sensitive, e.g.
 'a' and 'A' are considered different characters.
*/
func HasUniqueChar(str string) bool {
	count := 0
	for i, char := range str {

		if strings.Contains(str[i+1:], string(char)) {
			count++
		}

	}
	if count > 0 {
		return false
	}
	return true
}

func HasUniqueChar2(str string) bool {
	seen := make(map[rune]bool)
	for _, c := range str {
		if seen[c] {
		return false
	  }
	  seen[c] = true
	}
	
	return true
}

func HasUniqueChar3(str string) bool {
	var m = map[rune]bool{}
	for _, w := range str {
	  m[w] = true
	}
	return len(m)==len(str)
}


func HasUniqueChar4(str string) bool {
	for _,letter := range str {
		if strings.Count(str,string(letter)) > 1 {
		   return false
		   }
		}
	return true
}

func HasUniqueChar5(str string) bool {
	for i, x := range str {
	  if strings.LastIndex(str, string(x)) != i { return false }
	}
	return true
}