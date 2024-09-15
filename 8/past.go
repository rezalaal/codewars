package kata8

import (
	"fmt"
	"time"
)

/*
#Beginner Series #2 Clock

Clock shows h hours, m minutes and s seconds after midnight.

Your task is to write a function which returns the time since midnight in milliseconds.

Example:
h = 0
m = 1
s = 1

result = 61000
Input constraints:

0 <= h <= 23
0 <= m <= 59
0 <= s <= 59
*/
func Past(h, m, s int) int {
	return (h*60*60 + m*60 + s) * 1000
}

func Past2(h, m, s int) int {
	timeString := fmt.Sprintf("%dh%dm%ds", h, m, s)
	milli, _ := time.ParseDuration(timeString)
	return int(milli.Milliseconds())
}
