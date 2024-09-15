package kata7

/*
#Factorial

Your task is to write function factorial.
*/
func Factorial(n int) int {
	if n < 2 {
		return 1
	}
	return Factorial(n-1) * n
}

func Factorial1(n int) int {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}