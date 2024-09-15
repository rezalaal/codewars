package kata7

import (
	"time"
)

/*
Unlucky Days
Friday 13th or Black Friday is considered as unlucky day. Calculate how many unlucky days are in the given year.

Find the number of Friday 13th in the given year.

Input: Year in Gregorian calendar as integer.

Output: Number of Black Fridays in the year as an integer.

Examples:

unluckyDays(2015) == 3
unluckyDays(1986) == 1
*/
func UnluckyDays(year int) int {
	count := 0
	for month:=1;month<=12;month++ {
		dte := time.Date(year, time.Month(month), 13, 0, 0, 0, 0, time.UTC)
		if dte.Weekday() == time.Friday {
			count++
		}
	}
	
	return count
}