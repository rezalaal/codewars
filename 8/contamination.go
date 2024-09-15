package kata8

import "strings"

/*
#Contamination #1 -String-

An AI has infected a text with a character!!

This text is now fully mutated to this character.

If the text or the character are empty, return an empty string.
There will never be a case when both are empty as nothing is going on!!

Note: The character is a string of length 1 or an empty string.

Example
text before = "abc"
character   = "z"
text after  = "zzz"
*/
func Contamination(text, char string) (out string) {
	if len(text) == 0 || len(char) == 0 {
		return
	}
	for _ = range text {
		out += char
	}
	return
}

func Contamination1(text, char string) string {
	result := make([]rune, len(text))
	if text == "" || char == "" {
		return ""
	}
	for i := range text {
		result[i] += rune(char[0])
	}
	return string(result)
}

func Contamination2(text, char string) string {
	return strings.Repeat(char, len(text))
}
