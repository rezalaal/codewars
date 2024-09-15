package kata7

import (
	"sort"
	"strings"
)

/*
#Two to One

Take 2 strings s1 and s2 including only letters from a to z. Return a new sorted string, the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.

Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/
func TwoToOne(s1 string, s2 string) string {
	letters := []string{}
	for _, char := range s1 {
		letters = append(letters, string(char))
	}
	for _, char := range s2 {
		letters = append(letters, string(char))
	}
	sort.Strings(letters)
	
	seen := make(map[string]bool) 
	var result []string           

	for _, char := range letters {
		if !seen[char] { 
			seen[char] = true 
			result = append(result, string(char)) 
		}
	}
	
	return strings.Join(result, "")
}

func TwoToOne1(s1 string, s2 string) string {
	chars := strings.Split(s1 + s2, "")
	sort.Strings(chars)
	result := ""
	for _, s := range chars {
	  chr := string(s)
	  if !strings.Contains(result, chr) {
		result = result + chr
	  }
	}
	return result
}

func TwoToOne2(s1 string, s2 string) string {
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
	  if strings.Contains(s1+s2, char) {
		result += char
	  }
	}
	return result
}