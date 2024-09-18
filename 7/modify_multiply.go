package kata7

import (
	"fmt"
	"strings"
)

/*
Multiply Word in String
You are to write a function that takes a string as its first parameter. This string will be a string of words.

You are expected to then use the second parameter, which will be an integer, to find the corresponding word in the given string. The first word would be represented by 0.

Once you have the located string you are finally going to multiply by it the third provided parameter, which will also be an integer. You are additionally required to add a hyphen in between each word.

# Example

modifyMultiply ("This is a string", 3, 5)
*/
func ModifyMultiply(str string, loc, num int) string {
	words := strings.Split(str, " ")

	if loc < 0 || loc >= len(words) {
		return ""
	}

	if num <= 0 {
		return words[loc]
	}

	result := []string{}
	for i := 0; i < num; i++ {
		result = append(result, words[loc])
	}
	return strings.Join(result, "-")
}

func ModifyMultiply1(str string, loc, num int) string {
	res := strings.Fields(str)
	if num == 0 {
		return res[loc]
	}
	result := strings.Repeat(fmt.Sprintf("%s-", res[loc]), num)

	return strings.TrimSuffix(result, "-")
}

func ModifyMultiply2(str string, loc, num int) string {
	word := strings.Fields(str)[loc]
	if num == 0 {
		return word
	}
	return strings.TrimRight(strings.Repeat(word+"-", num), "-")
}
