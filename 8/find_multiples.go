package kata8

/*
#Find Multiples of a Number

In this simple exercise, you will build a program that takes a value, integer , and returns a list of its multiples up to another value, limit . If limit is a multiple of integer, it should be included as well. There will only ever be positive integers passed into the function, not consisting of 0. The limit will always be higher than the base.

For example, if the parameters passed are (2, 6), the function should return [2, 4, 6] as 2, 4, and 6 are the multiples of 2 up to 6.
*/
func FindMultiples(integer, limit int) []int {
	result := []int{}
	n := integer
	for {
		if n <= limit {
			result = append(result, n)
			n += integer
		}else{ break }

	}
	return result
}

func FindMultiples1(integer, limit int) (res[]int) {
	for i := integer; i <= limit; i += integer { res = append(res, i)}
	return
}

func FindMultiples2(integer, limit int) (res []int) {
	for i := 1; i*integer <= limit; i++ {
	  res = append(res, i*integer)
	}
	return res
}