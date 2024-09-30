package kata7


/*
Simple string reversal II

In this Kata, you will be given a string and two indexes (a and b). Your task is to reverse the portion of that string between those two indices inclusive.

solve("codewars",1,5) = "cawedors" -- elements at index 1 to 5 inclusive are "odewa". So we reverse them.
solve("cODEWArs", 1,5) = "cAWEDOrs" -- to help visualize.
Input will be lowercase and uppercase letters only.

The first index a will always be lower that than the string length; the second index b can be greater than the string length. More examples in the test cases. Good luck!
*/
func SimpleStringReversalII(s string, a, b int) string {
	if b > len(s) {
		b = len(s) - 1
	}
	tail := s[a:b+1]
	r := ReverseLetters(tail)

	return s[:a] + r + s[b+1:]
}

func SimpleStringReversalII2(s string, a, b int) string {
	if a<0 {a=0}
	if b>len(s) {b=len(s)-1}
	r:=[]rune(s)
	for a<b {
	  r[a],r[b]=r[b],r[a]
	  a++
	  b--
	}
	return string(r)
}

func SimpleStringReversalII3(s string, a, b int) (r string) {
	if b >= len(s) { b = len(s) - 1 } // limit b to prevent index out of range
  
  r += s[:a] // add the staring part
  
  // reverse the middle
  for i := b; i >= a; i-- {
    r += string(s[i])
  }
  
  return r + s[b + 1:] // add the end and return
}