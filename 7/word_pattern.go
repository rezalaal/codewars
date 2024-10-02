package kata7

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Cryptanalysis Word Patterns
In cryptanalysis, words patterns can be a useful tool in cracking simple ciphers.

A word pattern is a description of the patterns of letters occurring in a word,
where each letter is given an integer code in order of appearance. So the first letter is given the code 0,
and second is then assigned 1 if it is different to the first letter or 0 otherwise, and so on.

As an example, the word "hello" would become "0.1.2.2.3".
For this task case-sensitivity is ignored, so "hello", "helLo" and "heLlo" will all return the same word pattern.

Your task is to return the word pattern for a given word.
All words provided will be non-empty strings of alphabetic characters only, i.e. matching the regex "[a-zA-Z]+".
*/
func WordPattern(word string) string {
	words := make(map[string]int)
	var res strings.Builder
	counter := 0

	for _, char := range word {
		w := strings.ToLower(string(char))
		if _, exists := words[w]; !exists {
			words[w] = counter
			counter++
		}
		res.WriteString(strconv.Itoa(words[w]))
		res.WriteString(".")
	}

	result := res.String()
	if len(result) > 0 {
		result = result[:len(result)-1]
	}

	return result
}
func WordPattern1(word string) string {
	m := map[rune]int{}
	pos := 0
	d := ""
	var sb strings.Builder
	for _,r := range strings.ToLower(word) {
	  if _,ok := m[r]; !ok { m[r] = pos; pos++ }
	  fmt.Fprintf(&sb, "%s%d",d,m[r])
	  d = "."
	}
	return sb.String()
}
func WordPattern2(word string) string {
	hsh := make(map[string]int)
	i := 0
	for _, v := range word {
	  char := strings.ToLower(string(v))
	  _, ok := hsh[char]
	  if ok { continue }
	  hsh[char] = i
	  i += 1
	}
	res := []string{}
	for _, v := range word {
	  char := strings.ToLower(string(v))
	  n := strconv.Itoa(hsh[char])
	  res = append(res, n)
	}
	return strings.Join(res, ".")
}