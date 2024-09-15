package kata8

import (
	"strings"
	"regexp"
)

/*
#Remove String Spaces

Write a function that removes the spaces from the string, then return the resultant string.

Examples:

Input -> Output
"8 j 8   mBliB8g  imjB8B8  jl  B" -> "8j8mBliB8gimjB8B8jlB"
"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd" -> "88Bifk8hB8BB8BBBB888chl8BhBfd"
"8aaaaa dddd r     " -> "8aaaaaddddr"
*/
func NoSpace(word string) string {
	result := ""
	for _, char := range word {
		if char != 32 {
			result += string(char)
		}
	}
	return result
}

func NoSpace2(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func NoSpace3(word string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllString(word, "")
}