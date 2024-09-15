package kata8

import (
	"strings"
	"regexp"
)

type MyString string


/*
#Is the string uppercase?

Is the string uppercase?
Task
Create a method to see whether the string is ALL CAPS.

Examples (input -> output)
"c" -> False
"C" -> True
"hello I AM DONALD" -> False
"HELLO I AM DONALD" -> True
"ACSKLDFJSgSKLDFJSKLDFJ" -> False
"ACSKLDFJSGSKLDFJSKLDFJ" -> True
In this Kata, a string is said to be in ALL CAPS whenever it does not contain any lowercase letter so any string containing no letters at all is trivially considered to be in ALL CAPS.
*/
func (s MyString) IsUpperCase() bool {
	if strings.ToUpper(string(s)) == string(s) {
		return true
	}
	return false
}

func (s MyString) IsUpperCase2() bool {
	return s == MyString(strings.ToUpper(string(s)))
}

func (s MyString) IsUpperCase3() bool {
	r, _ := regexp.Compile(`^[A-Z\s]+$`)
	  return r.MatchString(string(s))
}
