package kata7

import (
	"slices"
	"strings"
)

/*
#String ends with?

Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false
*/
func Solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func Solution3(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
}

func Solution1(str, ending string) bool {
	
	letters := strings.Split(str, "")
	
	letters = DoReverse(letters)
	
	if len(ending) == 0 {
		return true
	}
	if len(str) == 0 {
		return false
	}
	for i, ch := range ending {
		if letters[i] != string(ch) {
			return false
		}
	}
	return true

}

func DoReverse(letters []string) []string {
	res := []string{}
	for i := len(letters) - 1; i >= 0; i-- {
		res = append(res, letters[i])
	}
	return res
}


func Solution2(str, ending string) bool {
	
	letters := strings.Split(str,"")
	
	slices.Reverse(letters)
	if len(ending) == 0 {return true}
	if len(str) == 0 {return false}
	for i,ch :=range ending {
		if letters[i] != string(ch) {
			return false
		}		
	}
	return true

}