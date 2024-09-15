package kata7

import (
	"strings"
	"regexp"
)

/*
#Remove duplicate words

Your task is to remove all duplicate words from a string, leaving only single (first) words entries.

Example:

Input:

'alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta'

Output:

'alpha beta gamma delta'
*/
func RemoveDuplicateWords(str string) string {
	words := strings.Split(str, " ")
	seen := make(map[string]bool)
	var result []string

	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			result = append(result, word)
		}
	}
	
	return strings.TrimRight(strings.Join(result, " "), " ")
}

func RemoveDuplicateWords1(str string) (s string) {
	m := make(map[string]int)
	for _, word := range strings.Split(str, " ") {
		if _, ok := m[word]; ok {
			continue
		}
		m[word]++
		s += word + " "
	}
	s = strings.TrimRight(s, " ")
  return // reverse those words
}

func RemoveDuplicateWords2(str string) string {
	strArr := strings.Split(str, " ")
	returnStr := " "
	for _, v := range strArr {
	  matched, _ := regexp.MatchString(`\s`+v+`\s`, returnStr)
	  if(!matched) {
		returnStr += v + " "
	  }
	}
	return strings.TrimSpace(returnStr)
}