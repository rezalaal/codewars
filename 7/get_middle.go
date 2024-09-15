package kata7

/*
#Get the Middle Character

You are going to be given a word. Your job is to return the middle character of the word. If the word's length is odd, return the middle character. If the word's length is even, return the middle 2 characters.

#Examples:

Kata.getMiddle("test") should return "es"

Kata.getMiddle("testing") should return "t"

Kata.getMiddle("middle") should return "dd"

Kata.getMiddle("A") should return "A"
#Input

A word (string) of length 0 < str < 1000 (In javascript you may get slightly more than 1000 in some test cases due to an error in the test cases). You do not need to test for this. This is only here to tell you that you do not need to worry about your solution timing out.

#Output

The middle character(s) of the word represented as a string.
*/
func GetMiddle(s string) string {
	if len(s)%2 == 0 {
		position := (len(s)/2 - 1)
		return string(s[position]) + string(s[position+1])
	} else {
		position := ((len(s)+1)/2 - 1)
		return string(s[position])
	}
}

func GetMiddle1(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	return s[(n-1)/2 : n/2+1]
}

func GetMiddle2(s string) string {
	mid := len(s) / 2
	if len(s)%2 == 1 {
		return string(s[mid])
	}
	return string(s[mid-1]) + string(s[mid])
}

func GetMiddle3(s string) string {
	l := len(s)
	i := l / 2
	if l%2 == 0 {
		return s[i-1 : i+1]
	}
	return string(s[i])
}
