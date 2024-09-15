package kata7

/*
#Sum of Cubes

Write a function that takes a positive integer n, sums all the cubed values from 1 to n (inclusive), and returns that sum.

Assume that the input n will always be a positive integer.

Examples: (Input --> output)

2 --> 9 (sum of the cubes of 1 and 2 is 1 + 8)
3 --> 36 (sum of the cubes of 1, 2, and 3 is 1 + 8 + 27)
*/
func SumCubes(n int) int {
	sum := 0
	for i:=1;i<=n;i++ {
		sum += i * i * i
	}
	return sum
}


func SumCubes2(n int) int {
	n = n * (n + 1)
	return n * n / 4
}