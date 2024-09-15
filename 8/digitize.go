package kata8

import "strconv"

/*
#Convert number to reversed array of digits

Convert number to reversed array of digits
Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

Example(Input => Output):
35231 => [1,3,2,5,3]
0 => [0]
*/
func Digitize(n int) []int {
	s := strconv.Itoa(n)
	arr := make([]int, len(s))
	j := len(s) - 1
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		arr[j] = digit
		j--
	}
	return arr
}

func Digitize2(n int) []int {
	var ret []int
	for {
		ret = append(ret, n%10)
		n /= 10
		if n == 0 {
			return ret
		}
	}
}

func Digitize3(n int) (a[]int) {
    if n == 0 { a = append(a, 0) }
    for ; n>0; n /= 10 { a = append(a, n % 10) }
    return
}