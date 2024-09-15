package kata7

import (
	"fmt"
	"strconv"
)

/*
#Number of Decimal Digits

Determine the total number of digits in the integer (n>=0) given as input to the function. For example, 9 is a single digit, 66 has 2 digits and 128685 has 6 digits. Be careful to avoid overflows/underflows.

All inputs will be valid.
*/
func Digits(n uint64) int {
	
	s := strconv.FormatUint(n, 10)
	return len(s)
}

func Digits1(n uint64) int {
	return len(fmt.Sprintf("%d", n))
}

func Digits2(n uint64) (numDigits int) {
	for numDigits = 1; n > 9; numDigits++ {
	  n /= 10
	}
	return
}

func Digits3(n uint64) int {
	if n < 10 { return 1 } else { return 1 + Digits(n / 10) }
}