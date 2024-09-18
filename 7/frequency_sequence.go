package kata7

import (
	"strconv"
	"strings"
)

/*
Frequency sequence
Your task is to return an output string that translates an input string s by replacing each character in s
with a number representing the number of times that character occurs in s and separating each number
with the sep character(s).

Example (s, sep --> Output)

"hello world", "-" --> "1-1-3-3-2-1-1-2-1-3-1"
"19999999"   , ":" --> "1:7:7:7:7:7:7:7"
"^^^**$"     , "x" --> "3x3x3x2x2x1"
*/
func FreqSeq(str string, sep string) string {
	res := ""
	for i, char := range str {
		res += strconv.Itoa(strings.Count(str, string(char)))
		if i<len(str) - 1 {
			res += sep
		}		
	}
	return res
}

func FreqSeq1(str string, sep string) string {
	numbers := make([]string, 0, len(str))
	  for _, v := range str {
		  numbers = append(numbers, strconv.Itoa(strings.Count(str, string(v))))
	  }
	  return strings.Join(numbers, sep)
}

func FreqSeq2(str string, sep string) string {

	buf := map[rune]int{}
	result := []string{}

	for _, ch := range str {
		buf[ch]++
	}
	for _, ch := range str {
		result = append(result, strconv.Itoa(buf[ch]))
	}
	return strings.Join(result, sep)
}