package kata7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
Sum of integers in string
Your task in this kata is to implement a function that calculates the sum of the integers inside a string. For example,
in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.

Note: only positive integers will be tested.
*/
func SumOfIntegersInString(strng string) int {
	var res string
	for _, char := range strng {		
		if char >= '0' && char <= '9' {
			res += string(char)
		}else{
			res += " "
		}
	}
	
	numbers := strings.Fields(res)
	sum := 0
	n := 0
	for _, number :=range numbers {
		n,_ = strconv.Atoi(number)
		sum += n
	}
	return sum
}

func SumOfIntegersInString2(strng string) int {
	r := 0
	for _, s := range regexp.MustCompile(`\d+`).FindAllString(strng, -1) {
	  x, _ := strconv.Atoi(s)
	  r += x
	}
	return r
}
func SumOfIntegersInString3(strng string) int {
	var end, buf int
	  for _, ch := range strng {
		  if ch >= '0' && ch <= '9' {
			  buf = buf*10 + int(ch-'0')
		  } else {
			  end += buf
			  buf = 0
		  }
	  }
	  return end + buf
}

func SumOfIntegersInString4(strng string) int {
	//Your task in this kata is to implement a function that calculates the sum of the integers inside a string.
	// For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.

	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(strng, -1))

	sum := 0
	for _, i := range re.FindAllString(strng, -1) {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum
}