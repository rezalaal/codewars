package kata7

import (
	"strconv"
)

func Strong(n int) string {
	number := strconv.Itoa(n)
	sum := 0 
	for _, char :=range number {
		digit, _ := strconv.Atoi(string(char))
		sum += Factorial(digit)
	}
	if sum != n  {
		return "Not Strong !!"
	}
	return "STRONG!!!!"
}