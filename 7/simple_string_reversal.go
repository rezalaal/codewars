package kata7

import (
	"strings"
)

/*
Simple string reversal
In this Kata, we are going to reverse a string while maintaining the spaces (if any) in their original place.

For example:

solve("our code") = "edo cruo"
-- Normal reversal without spaces is "edocruo".
-- However, there is a space at index 3, so the string becomes "edo cruo"

solve("your code rocks") = "skco redo cruoy".
solve("codewars") = "srawedoc"
More examples in the test cases. All input will be lower case letters and in some cases spaces.

Good luck!
*/
func SimpleStringReversal(s string) string {
	words := strings.Split(s, " ")	
	reversed := reverse(strings.Join(words, ""))
	result := ""
	cur := 0
	for _, char := range s {
		if char == ' ' {
			result += " "
		}else{
			result += string(reversed[cur])
			cur++
		}		
	}
	return result
}
func SimpleStringReversal1(s string) string {
	r := []rune(s)
	for i,j := 0,len(r)-1; i<j; i++ {
	  for ; i<j && r[i]==' '; i++ {  }
	  for ; i<j && r[j]==' '; j-- {  }
	  r[i],r[j] = r[j],r[i]
	  j--
	}
	return string(r)
}
func SimpleStringReversal2(s string) string {
	spaceIndexes := []int{}
	reverseString := ""
	  for i, v := range s {
	  if string(v) == " " {
		spaceIndexes = append(spaceIndexes, i)
	  } else {
		reverseString = string(v) + reverseString
	  }
	}
	
	for _, index := range spaceIndexes {
	  reverseString = reverseString[:index] + " " + reverseString[index:]
	}
	
	return reverseString
}