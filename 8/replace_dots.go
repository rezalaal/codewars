package kata8

import (
	"strings"
	"regexp"
)

/*
#Replace all dots

The code provided is supposed replace all the dots . in the specified String str with dashes -

But it's not working properly.

Task
Fix the bug so we can all go home early.

Notes
String str will never be null.
*/
func ReplaceDots(str string) string {
	result := ""
	for _, char := range str {
		if string(char) == "." {
			result += "-"
		}else{
			result += string(char)
		}		
	}
	return result
}

func ReplaceDots1(str string) string {
	return strings.ReplaceAll(str, ".", "-")
}

func ReplaceDots2(str string) string {
	return regexp.MustCompile(`\.`).ReplaceAllString(str, "-")
}