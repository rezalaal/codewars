package kata8

import (
	"strings"
	"unicode"
)

/*
#altERnaTIng cAsE <=> ALTerNAtiNG CaSe

altERnaTIng cAsE <=> ALTerNAtiNG CaSe
Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:

ToAlternatingCase("hello world"); // => "HELLO WORLD"
ToAlternatingCase("HELLO WORLD"); // => "hello world"
ToAlternatingCase("hello WORLD"); // => "HELLO world"
ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
As usual, your function/method should be pure, i.e. it should not mutate the original string.
*/

func ToAlternatingCase(str string) string {
	result := ""
	for _, char := range str {
		if IsUpper(string(char)) {
			result += strings.ToLower(string(char))
		} else {
			result += strings.ToUpper(string(char))
		}
	}
	return result
}

func IsUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func ToAlternatingCase2(str string) string {
	alternate := func(r rune) rune {
	  if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	  } else {
		return unicode.ToLower(r)
	  }
	}
	
	return strings.Map(alternate, str)
}