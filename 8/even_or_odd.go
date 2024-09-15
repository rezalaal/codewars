package kata8

/*
#Even or Odd

Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

*/
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

/*
If number is even, its binary representation will end in 0, and number & 1 will yield 0.
If number is odd, its binary representation will end in 1, and number & 1 will yield 1.
*/
func EvenOrOdd2(number int) string {
	return []string{"Even", "Odd"}[number&1]
}

func EvenOrOdd3(number int) string {
	return Ternary(number&1 == 0, "Even", "Odd")
}

func Ternary[T any](statement bool, a, b T) T {
	if statement {
		return a
	}
	return b
}

func EvenOrOdd4(number int) string {
	return map[bool]string{true: "Even", false: "Odd"}[number%2 == 0]
}