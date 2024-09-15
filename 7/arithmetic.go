package kata7

/*
#Make a function that does arithmetic!

Given two numbers and an arithmetic operator (the name of it, as a string), return the result of the two numbers having that operator used on them.

a and b will both be positive integers, and a will always be the first number in the operation, and b always the second.

The four operators are "add", "subtract", "divide", "multiply".

A few examples:(Input1, Input2, Input3 --> Output)

5, 2, "add"      --> 7
5, 2, "subtract" --> 3
5, 2, "multiply" --> 10
5, 2, "divide"   --> 2.5
Try to do it without using if statements!
*/
func Arithmetic(a int, b int, operator string) (res int) {
	switch operator {
	case "add":
		res = a + b
	case "subtract":
		res = a - b
	case "multiply":
		res = a * b
	case "divide":
		res = a / b
	}
	return
}

func add(a int, b int) int      { return a + b }
func subtract(a int, b int) int { return a - b }
func multiply(a int, b int) int { return a * b }
func divide(a int, b int) int   { return a / b }

func Arithmetic2(a int, b int, operator string) int {
	f := map[string]func(int, int) int{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":   divide,
	}

	return f[operator](a, b)
}

func Arithmetic3(a int, b int, operator string) int {
	op := map[string]int{"add": a + b, "subtract": a - b, "multiply": a * b, "divide": a / b}
	return op[operator]
}
