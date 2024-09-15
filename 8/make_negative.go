package kata8

import (
	"math"
)

/*
#Return Negative

In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

Examples

MakeNegative(1)    // return -1
MakeNegative(-5)   // return -5
MakeNegative(0)    // return 0
Notes
The number can be negative already, in which case no change is required.
Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.
*/
func MakeNegative(x int) int {	
	if x > 0 {		
		return -x
	}	
	return x
}

func MakeNegative1(x int) int {
	return -int(math.Abs(float64(x)))
}