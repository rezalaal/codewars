package kata8

import (
	"strings"
)

/*
#Is this my tail?

Some new animals have arrived at the zoo. The zoo keeper is concerned that perhaps the animals do not have the right tails. To help her, you must correct the broken function to make sure that the second argument (tail), is the same as the last letter of the first argument (body) - otherwise the tail wouldn't fit!

If the tail is right return true, else return false.

The arguments will always be non empty strings, and normal letters.

*/
func CorrectTail(body string, tail rune) bool {
	if rune(body[len(body)-1]) == tail {
		return true
	}
	return false
}

func CorrectTail2(body string, tail rune) bool {
	return rune(body[len(body)-1]) == tail
}

func CorrectTail3(body string, tail rune) bool {
	return strings.HasSuffix(body, string(tail))
}