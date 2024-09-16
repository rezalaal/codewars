package kata7

/*How many are smaller than me?
Write a function that given, an array arr, returns an array containing at each index i the amount of numbers that 
are smaller than arr[i] to the right.
چند تا عدد سمت راست آرایه ورودی از این عدد کوچک تر است
For example:

* Input [5, 4, 3, 2, 1] => Output [4, 3, 2, 1, 0]
* Input [1, 2, 0] => Output [1, 1, 0]
If you've completed this one and you feel like testing your performance tuning of this same kata, 
head over to the much tougher version How many are smaller than me II?
http://www.codewars.com/kata/56a1c63f3bc6827e13000006
*/
func Smaller(arr []int) []int {
	var result = make([]int, len(arr))
	for i, n := range arr {
		count := 0
		for _, m := range arr[i+1:] {
			if m < n {
				count++
			}
		}
		result[i] = count
	}
	return result
}

func Smaller2(a []int) []int {
	b := make([]int, len(a))
	for i := range a {
	  for j := range a[i+1:] {
		if a[i]>a[i+j+1] {b[i]++ }
	  }
	}
	return b
}
