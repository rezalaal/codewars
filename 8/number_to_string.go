package kata8

import (
	"fmt"
	"strconv"
)

/*
#Convert a Number to a String!

We need a function that can transform a number (integer) into a string.

What ways of achieving this do you know?

Examples (input --> output):
123  --> "123"
999  --> "999"
-100 --> "-100"
*/
// var NumberToString = strconv.Itoa

func NumberToString(n int) string {
	return fmt.Sprintf("%d", n)
}

func NumberToString1(n int) string {
	return fmt.Sprint(n)
}

func NumberToString2(n int) string {
	return strconv.Itoa(n)
}

func NumberToString3(n int) string {
	return strconv.FormatInt(int64(n), 10)
}


func DigitToString(n int) string {
	if n == 0 {return "0"}
	if n == 1 {return "1"}
	if n == 2 {return "2"}
	if n == 3 {return "3"}
	if n == 4 {return "4"}
	if n == 5 {return "5"}
	if n == 6 {return "6"}
	if n == 7 {return "7"}
	if n == 8 {return "8"}
	if n == 9 {return "9"}
	return ""
}
  
  
  
func NumberToString4(n int) string {
	var str string
	var isNegative bool = false
	if n < 0 {
	  n *= -1
	  isNegative = true
	}
	for ; n > 0; {
	  str = DigitToString(n % 10) + str
	  n /= 10
	}
	if isNegative {
	  return "-" + str
	}
	return str
}

func NumberToString5(n int) string {
	neg := ""
	if n < 0 {
	  n *= -1
	  neg = "-"
	}
	
	var res string
	for n > 0 {
	  res = string(rune(n%10+int('0'))) + res
	  n/=10
	}
	return neg + res
}