package kata8

import (
	"strings"
)

/*
#Is it a palindrome?

Write a function that checks if a given string (case insensitive) is a palindrome.

A palindrome is a word, number, phrase, or other sequence of symbols that reads the same backwards as forwards, such as madam or racecar.
*/
func IsPalindrome(str string) bool {
	for i, v := range str {
		if strings.ToLower(string(str[len(str) - i - 1])) != strings.ToLower(string(v)) {
			return false
		}
	}
	return true
}