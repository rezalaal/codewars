package kata7

import (
	"math"
)

/*
#Sum of odd numbers

Given the triangle of consecutive odd numbers:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29
...
Calculate the sum of the numbers in the nth row of this triangle (starting at index 1) e.g.: (Input --> Output)

1 -->  1
2 --> 3 + 5 = 8
*/
func RowSumOddNumbers(n int) int {
    return n * n * n
}

func RowSumOddNumbers2(n int) int {
    return int(math.Pow(float64(n), 3))
}

func RowSumOddNumbers1(n int) int {
	// The first number in the nth row is n^2 - n + 1
    firstNumber := n*n - n + 1
    
    // Initialize the sum to 0
    sum := 0
    
    // Add the odd numbers in the nth row
    for i := 0; i < n; i++ {
        sum += firstNumber
        firstNumber += 2 // Move to the next odd number
    }
    
    return sum

}