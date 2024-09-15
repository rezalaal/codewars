package kata7

import (
	"strings"
)



// Use the preloaded Tuple struct as return type
type Tuple struct {
	Char  rune
	Count int
}

/*
#Ordered Count of Characters

Count the number of occurrences of each character and return it as a (list of tuples) in order of appearance. For empty output return (an empty list).

Consult the solution set-up for the exact data structure implementation depending on your language.

Example:

OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

// Where
type Tuple struct {
    Char  rune
    Count int
}
*/
func OrderedCount(text string) []Tuple {
	var result = []Tuple{}
	var countMap = make(map[rune]int)
	for _, char := range text {
		countMap[char]++
	}

	var seen = make(map[rune]bool)
	for _, char := range text {
		if !seen[char] {
			result = append(result, Tuple{Char: char, Count: countMap[char]})
			seen[char] = true
		}
	}
	return result
}


func OrderedCount1(text string) []Tuple {
  counts := make([]Tuple, 0, len(text))
  counted := ""
  
  for _, r := range(text) {
    if strings.Count(counted, string(r)) == 0 {
      counts = append(counts, Tuple{Char: r, Count: strings.Count(text, string(r))})
      counted += string(r)
    }
  }
  return counts 
}