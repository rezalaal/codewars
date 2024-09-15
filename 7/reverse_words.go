package kata7

import (
	"strings"
)


/*
#Reverse words

Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

Examples
"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"
*/
func ReverseWords(str string) string {
	temp := ""
	res := ""
	for _, char := range str {

		if string(char) != " " {
			temp += string(char)
		} else {
			res += reverseIt(temp) + " "
			temp = ""
		}
		
	}
	res += reverseIt(temp)
	return res
}

func reverseIt(str string) string {
	letters := strings.Split(str, "")
	res := []string{}
	for i := len(letters) - 1; i >= 0; i-- {
		res = append(res, letters[i])
	}
	return strings.Join(res, "")
}


func ReverseWords2(str string) string {
	strList := strings.Split(str, " ")
	var response []string
	for _, word := range strList {
	  letters := []rune(word)
	  for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
			letters[i], letters[j] = letters[j], letters[i]
		}
	  response = append(response, string(letters))
	}
	return strings.Join(response, " ")
}