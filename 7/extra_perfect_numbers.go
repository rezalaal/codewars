package kata7

import "fmt"

/*
Extra Perfect Numbers (Special Numbers Series #7)
Definition
Extra perfect number is the number that first and last bits are set bits.

Task
Given a positive integer N , Return the extra perfect numbers in range from 1 to N .

Warm-up (Highly recommended)
https://www.codewars.com/collections/playing-with-numbers

Playing With Numbers Series
https://www.codewars.com/collections/playing-with-numbers

Notes
Number passed is always Positive .

Returned array/list should contain the extra perfect numbers in ascending order from lowest to highest

Input >> Output Examples
extraPerfect(3)  ==>  return {1,3}
Explanation:
(1)10 =(1)2
First and last bits as set bits.

(3)10 = (11)2
First and last bits as set bits.

extraPerfect(7)  ==>  return {1,3,5,7}
Explanation:
(5)10 = (101)2
First and last bits as set bits.

(7)10 = (111)2
First and last bits as set bits.

Playing with Numbers Series
https://www.codewars.com/collections/playing-with-numbers

Playing With Lists/Arrays Series
https://www.codewars.com/collections/playing-with-lists-slash-arrays

For More Enjoyable Katas
http://www.codewars.com/users/MrZizoScream/authored

ALL translations are welcomed
Enjoy Learning !!
Zizou
*/
func ExtraPerfect(n int) []int {
	res := []int{}
	for i := 1; i <= n; i++ {
		if i % 2 != 0 {
			res = append(res, i)
		}
	}
	return res
}

func ExtraPerfect1(n int) []int {
    cur := 1
    result := []int{}
    for cur<=n {
      bits := fmt.Sprintf("%b", cur) 
      if bits[0]=='1' && bits[len(bits)-1]=='1' {
        result = append(result, cur)
      }
      cur+=2
    }
    return result
}
