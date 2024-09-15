package kata7

import (
	"strings"
)

/*
#Spacify

Modify the spacify function so that it returns the given string with spaces inserted between each character.

spacify=("hello world") -> # returns "h e l l o   w o r l d
*/
func Spacify(s string) string {
	result := ""
	for i, char := range s {		
		result += string(char)
		if i != len(s) - 1 {
			result += " "
		}
	}
	return result
}

func Spacify1(s string) string {
	return strings.Join(strings.Split(s, ""), " ")
}