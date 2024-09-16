package kata7


/*
Simple Fun #2: Circle of Numbers
Task
Consider integer numbers from 0 to n - 1 written down along the circle in such a way that the distance between any two
neighbouring numbers is equal (note that 0 and n - 1 are neighbouring, too).

Given n and firstNumber/first_number/first-number, find the number which is written in the radially opposite position
to firstNumber.

Example
For n = 10 and firstNumber = 2, the output should be 7

https://codefightsuserpics.s3.amazonaws.com/tasks/circleOfNumbers/img/example.png?_tm=1476003938167

Input/Output
[input] integer n

A positive even integer.

Constraints: 4 ≤ n ≤ 1000.

[input] integer firstNumber/first_number/first-number

Constraints: 0 ≤ firstNumber ≤ n - 1

[output] an integer

*/
func CircleOfNumbers(n int, firstNumber int) int {
	items := map[int]int{}
    
	start := int(n / 2)
	for i:=0;i<n;i++{	
		items[i] = start + i
	}
	for i, number := range items {
		if number >= n {
			items[i] = number % n
		}
	}
	
	return items[firstNumber]
}
func CircleOfNumbers1(n int, firstNumber int) int {
    return (firstNumber + n / 2) % n
}

func CircleOfNumbers2(n int, firstNumber int) int {
	return (n >> 1 + firstNumber) % n
}

func CircleOfNumbers3(n int, firstNumber int) int {
    num:= firstNumber +(n/2)
  if num>=n{
    return num -n
  }
  return num
}