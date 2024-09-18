package kata7

import (
	"strconv"
)

/*
Previous multiple of three
Given a positive integer n: 0 < n < 1e6, remove the last digit until you're left with a number that is a multiple of three.

Return n if the input is already a multiple of three, and if no such number exists, return null, a similar empty value, or -1.

Examples
1      => null
25     => null
36     => 36
1244   => 12
952406 => 9
*/
func PrevMultOfThree(n int) interface{} {
	if n < 10 {
		return nil
	}
	if n % 3 == 0 {
		return n
	}
	str := strconv.Itoa(n)
	last := -1
	for i:=len(str)-1;i>0;i-- {
		number,_ := strconv.Atoi(str[0:i])
		if number % 3 == 0 {
			last = number
			return last
		}
	}
	if last % 3 != 0 {
		return nil
	}
	return last
}
func PrevMultOfThree1(n int) interface{} {
  
    for i := n; i > 0; i /= 10 {
        if i % 3 == 0 {
            return i
        }
    }

    return nil
}
func PrevMultOfThree2(n int) interface{} {
	for n%3 != 0{
	  n /= 10
	}
	if n==0 {return nil}
	return n
}
func PrevMultOfThree3(n int) interface{} {
	// write your code here
	// your function should return an int or a nil
	if(n < 10 && n % 3 != 0) {
		return nil
	}
	
	if(n % 3 == 0){
	  return n
	} else {
	  return PrevMultOfThree(n / 10)
	}
}
func PrevMultOfThree4(n int) interface{} {
 
	for n%3 != 0 {
	  n = n/10
	}  
	switch {
	  case n == 0: return nil
	  default: return n
	}
}