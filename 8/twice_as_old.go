package kata8

import (
	"math"
)

/*
#Twice as old

Your function takes two arguments:

current father's age (years)
current age of his son (years)
Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old). The answer is always greater or equal to 0, no matter if it was in the past or it is in the future.
*/
func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	tao := dadYearsOld - sonYearsOld*2
	if tao < 0 {
		tao *= -1
	}
	return tao
}

func TwiceAsOld1(dadYearsOld int, sonYearsOld int) int { 
	return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
}