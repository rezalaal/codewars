package kata7

import (
	"regexp"
	"strings"
)

/*
#Disemvowel Trolls
Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.
*/
func Disemvowel(comment string) string {
	result := ""
	
	for _, char := range comment {
		if !isOnVowels(string(char)) {
			result += string(char)
		}
	}
	return result
}

func Disemvowel1(comment string) string {
	for _,c := range "aiueoAIUEO"{
		   comment = strings.ReplaceAll(comment, string(c), "")
	   }
	 return comment
}

func Disemvowel2(comment string) string {
  return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
}

func Disemvowel3(comment string) string {
	result := ""
	for _, v := range comment {
	  switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		  continue
	  }
	  result += string(v)
	}
	return result
}

func Disemvowel4(comment string) string {
	var replacer = strings.NewReplacer(
		"a", "",
		"e", "",
		"i", "",
		"o", "",
		"u", "",
		"A", "",
		"E", "",
		"I", "",
		"O", "",
		"U", "",
	)
	return replacer.Replace(comment)
}

func isOnVowels(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for _, vowel :=range vowels {
		if s == vowel {
			return true
		}
	}
	return false
}