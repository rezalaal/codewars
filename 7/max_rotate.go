package kata7

import (
	"strconv"
)

/*
Rotate for a Max
Let us begin with an example:

Take a number: 56789. Rotate left, you get 67895.

Keep the first digit in place and rotate left the other digits: 68957.

Keep the first two digits in place and rotate the other ones: 68579.

Keep the first three digits and rotate left the rest: 68597. Now it is over since keeping the first four it remains only one digit which rotated is itself.

You have the following sequence of numbers:

56789 -> 67895 -> 68957 -> 68579 -> 68597

and you must return the greatest: 68957.

Task
Write function max_rot(n) which given a positive integer n returns the maximum number you got doing rotations similar to the above example.

So max_rot (or maxRot or ... depending on the language) is such as:

max_rot(56789) should return 68957

max_rot(38458215) should return 85821534
*/
func MaxRot(n int64) int64 {
    max := n
	s := rotate(strconv.Itoa(int(n)))
	number,_ := strconv.Atoi(s)
	if int64(number) > max {
		max = int64(number)			
	}
	for i := 0; i < len(s); i++ {
		s = s[:i+1]+ rotate(s[i+1:])
		number,_ := strconv.Atoi(s)
		if int64(number) > max {
			max = int64(number)			
		}
	}
	return max
}

func rotate(s string) string {
	if len(s) == 0 {
		return s
	}
	return s[1:] + string(s[0])
}

func MaxRot2(n int64) int64 {
	str := strconv.FormatInt(n, 10)
	max := n
	
	for i := 0; i<len(str)-1 ; i++ {
	  str = str[:i]+str[i+1:]+string(str[i])
	  num, _:= strconv.ParseInt(str, 10, 64)
	  if max < num { max = num }
	}
	
	return max
}

func MaxRot3(n int64) (max int64){
	max = n
	strn := strconv.FormatInt(n, 10)

	var left string
	var middle string
	var right string
	var currentNumber int64

	for index := range strn[:len(strn)-1] {
		left = strn[:index]
		middle = strn[index:index+1]
		right = strn[index+1:]

		strn = left+right+middle

		currentNumber,_ = strconv.ParseInt(strn,10,64)
		if max < currentNumber {max = currentNumber}
		}
	return
}

