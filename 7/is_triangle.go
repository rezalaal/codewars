package kata7


import (
	"sort"
)

/*
#Is this a triangle?

Implement a function that accepts 3 integer values a, b, c. The function should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).

Examples:

Input -> Output
1,2,2 -> true
4,2,3 -> true
2,2,2 -> true
1,2,3 -> false
-5,1,3 -> false
0,2,3 -> false
1,2,9 -> false 

مجموع طول هر دو ضلع مثلث باید از طول ضلع سوم بزرگتر باشد. 
*/
func IsTriangle(a, b, c int) bool {
	return a > 0 && b > 0 && c > 0 && 
           a+b > c && a+c > b && b+c > a
}


func IsTriangle1(a, b, c int) bool {
	l := []int{a,b,c}
	sort.Ints(l)
	
	return l[0] + l[1] > l[2]
}