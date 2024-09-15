package kata8

/*
#Wilson primes

Wilson primes satisfy the following condition. Let P represent a prime number.

Then,

(P−1)!+1/P∗P
​
 
should give a whole number, where P! is the factorial of P

Your task is to create a function that returns true if the given number is a Wilson prime and false otherwise.
*/
func AmIWilson(n int) bool {
	return n == 5 || n == 13 || n == 563
}

func AmIWilson2(n int) bool {
    if n <= 1 {
        return false
    }
    
    factorial := 1
    for i := 2; i < n; i++ {
        factorial *= i
        factorial %= (n * n)
    }
    
    return ((factorial + 1) % (n * n)) == 0
}