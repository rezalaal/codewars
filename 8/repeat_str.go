package kata8

import (
	"strings"
)

/*
#String repeat

Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

Examples (input -> output)
6, "I"     -> "IIIIII"
5, "Hello" -> "HelloHelloHelloHelloHello"
*/
func RepeatStr(repetitions int, value string) string {
	res := ""
	for i := 1; i <= repetitions; i++ {
		res += value
	}

	return res
}

func RepeatStr2(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}