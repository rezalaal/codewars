package kata6

import (
	"fmt"
	"strings"
)

/*
#Valid Braces

Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!

All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

What is considered Valid?
A string of braces is considered valid if all braces are matched with the correct brace.

Examples
"(){}[]"   =>  True
"([{}])"   =>  True
"(}"       =>  False
"[(])"     =>  False
"[({})](]" =>  False
*/
func ValidBraces(str string) bool {
	stack := []rune{}

	for _, char := range str {
		switch char {
		case '{', '[', '(':
			stack = append(stack, char)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}

var m = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

func ValidBraces1(str string) bool {
	s := make([]string, 0)
	for _, r := range str {
		r := string(r)
		if len(s) > 0 && m[s[len(s)-1]] == r {
			s = s[:len(s)-1]
		} else {
			s = append(s, r)
		}
	}
	return len(s) == 0
}

func ValidBraces2(str string) bool {
	stack := make([]rune, 0)
	for _, c := range str {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func ValidBraces3(str string) bool {
	for x := 0; x <= len(str)+2; x++ {
		str = strings.Replace(str, "()", "", -1)
		str = strings.Replace(str, "[]", "", -1)
		str = strings.Replace(str, "{}", "", -1)
	}
	fmt.Println("STR=", str)
	if len(str) > 0 {
		return false
	}
	return true
}
