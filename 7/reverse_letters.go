package kata7

import (
	"strings"
	"unicode"
)


/*
#Simple Fun #176: Reverse Letter

Task
Given a string str, reverse it and omit all non-alphabetic characters.

Example
For str = "krishan", the output should be "nahsirk".

For str = "ultr53o?n", the output should be "nortlu".

Input/Output
[input] string str
A string consists of lowercase latin letters, digits and symbols.

[output] a string
*/
func ReverseLetters(s string) string {
	var letters []rune
	
	for _, letter :=range s {
		if unicode.IsLetter(letter) {
			letters = append(letters, letter)
		}
		
	}
	var res strings.Builder
	for i := len(letters) - 1; i >= 0; i-- {
		res.WriteRune(letters[i])
	}

	return res.String() 
}


func ReverseLetters1(s string) string {
	arr := []rune{}
	for _, v := range s {
		if unicode.IsLetter(v) {
			arr = append([]rune{v}, arr...)
		}
	}
	return string(arr)

}

func ReverseLetters3(s string) string {
	result := []byte{}
	
	for i := len(s) - 1; i >= 0; i-- {
	  if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
		result = append(result, s[i])
	  }
	}
	
	return string(result)
}


func ReverseLetters2(s string) string {
	var word string = ""
	  for i := len(s) - 1; i >= 0; i-- {
		  //
		  if (s[i] > 64 && s[i] < 91) || (s[i] > 96 && s[i] < 123) {
			  word += string(s[i])
		  }
	  }
	  return word
}