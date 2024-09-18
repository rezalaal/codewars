package kata7

/*
Fibonacci
Create function fib that returns n'th element of Fibonacci sequence (classic programming task).
*/
func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func Fib1(n int) int {
	a, b := 0, 1
	for ; n > 0; n-- {
		a, b = b, a+b
	}
	return a
}
