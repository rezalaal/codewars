package kata8

/*
#Reversed Strings

Complete the solution so that it reverses the string passed into it.

'world'  =>  'dlrow'
'word'   =>  'drow'
*/
func Solution(word string) string {
	ws := make([]string, len(word))
	for i, char := range word {
		ws[i] = string(char)
	}
	for i, j := 0, len(ws)-1; i < j; i, j = i+1, j-1 {
		ws[i], ws[j] = ws[j], ws[i]
	}

	result := ""
	for _, char :=range ws {
		result += char
	}

	return result
}

func Solution2(word string) string {
	var result = ""
	for _,c := range word {
	  result = string(c) + result
	}
	return result
}

func Solution3(word string) string {
	rr := []rune(word)
	l := len(rr)
	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		  rr[i], rr[j] = rr[j], rr[i]
	  }
	  return string(rr)
}