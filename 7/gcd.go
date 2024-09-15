package kata7



import "math"

/*
Greatest common divisor
Find the greatest common divisor of two positive integers. The integers can be large, so you need to find a clever solution.

The inputs x and y are always greater or equal to 1, so the greatest common divisor will always be an integer that
is also greater or equal to 1.
*/
func Gcd(x, y uint32) uint32 {
	max := 0
	minNumber := x
	if y < x {
		minNumber = y
	}
	if x == 1 && y == 1 {
		return 1
	}
    for i := 1; i< int(minNumber); i++ {		
		if int(x) % i == 0 && int(y) % i == 0 && i > max {			
			max = i
		}
	}
	return uint32(max)
}

func Gcd2(x, y uint32) uint32 {
    // Base case: if one of the numbers is zero, return the other number
    if y == 0 {
        return x
    }
    // Recursively apply the Euclidean algorithm
    return Gcd(y, x%y)
}

func Gcd3(x, y uint32) uint32 {
	for y != 0 {
	  x, y = y, x%y 
	}
	return x  
}

func Gcd4(x, y uint32) uint32 {
    max := uint32(math.Max(float64(x), float64(y)))
    min := uint32(math.Min(float64(x), float64(y)))
    for (max % min != 0) {
      temp := max % min
      max = min
      min = temp
    }
    return min
}