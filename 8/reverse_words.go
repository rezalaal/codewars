package kata8

import (
	"strings"
)

/*
#Reversed Words

Complete the solution so that it reverses all of the words within the string passed in.

Words are separated by exactly one space and there are no leading or trailing spaces.

Example(Input --> Output):

"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"
*/
func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	result := ""
	j := len(words)
	for i:=0;i<len(words);i++ {
		result = result + " " + words[j-1]
		j--
	}
	return strings.TrimLeft(result, " ")
}

func ReverseWords1(str string) string {
	words := strings.Split(str, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		  words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}