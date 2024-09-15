package kata8

import (
	"math"
	"strconv"
)

/*
#Bin to Decimal

Complete the function which converts a binary number (given as a string) to a decimal number.

*/
func BinToDec(bin string) int {
	number := 0
	if bin == "0" {
		return 0
	}
	for i, char := range bin {
		if string(char) == "1" {
			number += int(math.Pow(2, float64(len(bin)-i)-1))
		}		
	}
	return number
}

func BinToDec2(bin string) int {
	r, _ := strconv.ParseInt(bin, 2, 64)
	return int(r)
}

func BinToDec3(bin string) int {
	n := 0
	for _, r := range bin {
	  n *= 2
	  n += int(r-'0')
	}
	return n
}