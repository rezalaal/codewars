package kata7

import (
	"strconv"
	"strings"
)

/*
Char Code Calculation
Given a string, turn each character into its ASCII character code and join them together to create a number - let's call this number total1:

'ABC' --> 'A' = 65, 'B' = 66, 'C' = 67 --> 656667
Then replace any incidence of the number 7 with the number 1, and call this number 'total2':

total1 = 656667
              ^
total2 = 656661
              ^
Then return the difference between the sum of the digits in total1 and total2:

  (6 + 5 + 6 + 6 + 6 + 7)
- (6 + 5 + 6 + 6 + 6 + 1)
-------------------------
                       6
*/
func Calc(s string) int {
	total1 := ""
	total2 := ""
	for _, char :=range s {
		total1 += strconv.Itoa(int(char))
		
	}
	total2 += strings.ReplaceAll(total1, "7", "1")	
	n1 := sumOfDigits(total1)
	n2 := sumOfDigits(total2)

	return n1 - n2
}
func sumOfDigits(s string) int {
	sum := 0
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char)) // Convert character to integer
		sum += digit
	}
	return sum
}

func Calc1(s string) int {
    n := 0
    for _,c := range s {
        if (c/10) == 7 {n++}
        if (c%10) == 7 {n++}
    }
    return 6*n
}