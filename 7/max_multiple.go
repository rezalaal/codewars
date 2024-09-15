package kata7

/*
#Maximum Multiple

Task
Given a Divisor and a Bound , Find the largest integer N , Such That ,

Conditions :
N is divisible by divisor

# N is less than or equal to bound

N is greater than 0.

Notes
The parameters (divisor, bound) passed to the function are only positive values .
It's guaranteed that a divisor is Found .
Input >> Output Examples
divisor = 2, bound = 7 ==> return (6)
Explanation:
(6) is divisible by (2) , (6) is less than or equal to bound (7) , and (6) is > 0 .

divisor = 10, bound = 50 ==> return (50)
Explanation:
(50) is divisible by (10) , (50) is less than or equal to bound (50) , and (50) is > 0 .*

divisor = 37, bound = 200 ==> return (185)
Explanation:
(185) is divisible by (37) , (185) is less than or equal to bound (200) , and (185) is > 0 .
*/
func MaxMultiple(d, b int) int {
	max := 1
	for i := 1; i <= b; i++ {
		if i%d == 0 {
			max = i
		}
	}
	return max
}

func MaxMultiple2(d, b int) int {
	return b - b%d
}

func MaxMultiple3(d, b int) int {
	return (b / d) * d
}

func MaxMultiple4(d, b int) int {
	// your code here
	s := b / d
	return s * d
}
