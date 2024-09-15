package kata7

import (
	"regexp"
	"unicode"
)

/*
In this Kata, you will be given a string and your task will be to return a list of ints detailing the count of
uppercase letters, lowercase, numbers and special characters, as follows.

Solve("*'&ABCDabcde12345") = [4,5,5,3].
--the order is: uppercase letters, lowercase, numbers and special characters.
More examples in the test cases.

Good luck!
*/
func SimpleStringCharacters(s string) []int {
	upperCases := 0
	lowerCases := 0
	numbers := 0
	specialChars := 0
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			upperCases++
		}else if char >= 'a' && char <= 'z' {
			lowerCases++
		}else if char >= '0' && char <= '9' {
			numbers++
		}else{
			specialChars++
		}
	}
	
	return []int{upperCases, lowerCases, numbers, specialChars}
}

func SimpleStringCharacters2(s string) []int {
	r := make([]int, 4)
	for _, c := range s {
	  switch {
		case c >= 'A' && c <= 'Z': r[0]++;
		case c >= 'a' && c <= 'z': r[1]++;
		case c >= '0' && c <= '9': r[2]++;
		default: r[3]++;
	  }
	}
	return r
}

func SimpleStringCharacters3(s string) []int {
	
	uppercase := 0;
	lowercase := 0;
	number    := 0;
	character := 0;
	
	for _, v := range s {
	  if unicode.IsUpper(v){
		uppercase++
	  }else if unicode.IsLower(v){
		lowercase++
	  }else if unicode.IsNumber(v){
		number++
	  }else if unicode.IsPunct(v) || unicode.IsSymbol(v){
		character++
	  }
	}
	
	
	return []int{uppercase, lowercase, number, character}
	
}

func SimpleStringCharacters4(s string) (end []int) {
	end = append(end, len(regexp.MustCompile("[A-Z]").FindAllString(s, -1)))
	end = append(end, len(regexp.MustCompile("[a-z]").FindAllString(s, -1)))
	end = append(end, len(regexp.MustCompile("[0-9]").FindAllString(s, -1)))
	end = append(end, len(regexp.MustCompile(`[\W]`).FindAllString(s, -1)))
  return
}