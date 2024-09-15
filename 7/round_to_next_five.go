package kata7

/*
#Round up to the next multiple of 5

Given an integer as input, can you round it to the next (meaning, "greater than or equal") multiple of 5?

Examples:

input:    output:
0    ->   0
2    ->   5
3    ->   5
12   ->   15
21   ->   25
30   ->   30
-2   ->   0
-5   ->   -5
etc.
Input may be any positive or negative integer (including 0).

You can assume that all inputs are valid integers.
*/
func RoundToNextFive(n int) int {
	if n%5 == 0 {
        return n // If n is already a multiple of 5, return it
    }
    
    if n > 0 {
        return ((n + 4) / 5) * 5 // Round up for positive numbers
    } else {
        return (n / 5) * 5 // Round down for negative numbers
    }
}

func RoundToNextFive1(n int) int {
	for n % 5 != 0 {
	  n++
	}
	return n
}

func RoundToNextFive2(n int) int {
	return n + (5 - n % 5) % 5;
}

func RoundToNextFive3(n int) int {
	remainder := n % 5
	if remainder < 0 {
		n += -remainder
	}
	if remainder > 0 {
		n += 5 - remainder
	}
	return n
}