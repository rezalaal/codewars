package kata7

import (
	"math"
)

/*
#Deodorant Evaporator

This program tests the life of an evaporator containing a gas.

We know the content of the evaporator (content in ml), the percentage of foam or gas lost every day (evap_per_day) and the threshold (threshold) in percentage beyond which the evaporator is no longer useful. All numbers are strictly positive.

The program reports the nth day (as an integer) on which the evaporator will be out of use.

Example:
evaporator(10, 10, 5) -> 29
Note:
Content is in fact not necessary in the body of the function "evaporator", you can use it or not use it, as you wish. Some people might prefer to reason with content, some other with percentages only. It's up to you but you must keep it as a parameter because the tests have it as an argument.

Example Calculation
For the input Evaporator(10, 10, 5):
Initial content = 10 ml
Evaporation rate = 10% per day
Threshold = 5% of initial content = 0.5 ml
Day-by-Day Breakdown
Day 1: Content = 10 - (10 * 10 / 100) = 9.0 ml
Day 2: Content = 9.0 - (9.0 * 10 / 100) = 8.1 ml
Day 3: Content = 8.1 - (8.1 * 10 / 100) = 7.29 ml
Day 4: Content = 7.29 - (7.29 * 10 / 100) = 6.561 ml
Day 5: Content = 6.561 - (6.561 * 10 / 100) = 5.9049 ml
Day 6: Content = 5.9049 - (5.9049 * 10 / 100) = 5.31441 ml
Day 7: Content = 5.31441 - (5.31441 * 10 / 100) = 4.783969 ml
Day 8: Content = 4.783969 - (4.783969 * 10 / 100) = 4.3055721 ml
Day 9: Content = 4.3055721 - (4.3055721 * 10 / 100) = 3.87499989 ml
Day 10: Content = 3.87499989 - (3.87499989 * 10 / 100) = 3.487499901 ml
Day 11: Content = 3.487499901 - (3.487499901 * 10 / 100) = 3.1387499109 ml
Day 12: Content = 3.1387499109 - (3.1387499109 * 10 / 100) = 2.82487491981 ml
Day 13: Content = 2.82487491981 - (2.82487491981 * 10 / 100) = 2.542387427829 ml
Day 14: Content = 2.542387427829 - (2.542387427829 * 10 / 100) = 2.2881486850461 ml
Day 15: Content = 2.2881486850461 - (2.2881486850461 * 10 / 100) = 2.05933381654149 ml
Day 16: Content = 2.05933381654149 - (2.05933381654149 * 10 / 100) = 1.853400434887341 ml
Day 17: Content = 1.853400434887341 - (1.853400434887341 * 10 / 100) = 1.6680603913986079 ml
Day 18: Content = 1.6680603913986079 - (1.6680603913986079 * 10 / 100) = 1.5012543522587471 ml
Day 19: Content = 1.5012543522587471 - (1.5012543522587471 * 10 / 100) = 1.3511289170338724 ml
Day 20: Content = 1.3511289170338724 - (1.3511289170338724 * 10 / 100) = 1.2160150253304852 ml
Day 21: Content = 1.2160150253304852 - (1.2160150253304852 * 10 / 100) = 1.0944135227974367 ml
Day 22: Content = 1.0944135227974367 - (1.0944135227974367 * 10 / 100) = 0.9849721705176931 ml
Day 23: Content = 0.9849721705176931 - (0.9849721705176931 * 10 / 100) = 0.8864749534659238 ml
Day 24: Content = 0.8864749534659238 - (0.8864749534659238 * 10 / 100) = 0.7978274571185314 ml
Day 25: Content = 0.7978274571185314 - (0.7978274571185314 * 10 / 100) = 0.7180447114066783 ml
Day 26: Content = 0.7180447114066783 - (0.7180447114066783 * 10 / 100) = 0.6452392402656105 ml
Day 27: Content = 0.6452392402656105 - (0.6452392402656105 * 10 / 100) = 0.5807153162380495 ml
Day 28: Content = 0.5807153162380495 - (0.5807153162380495 * 10 / 100) = 0.5226437846142446 ml
Day 29: Content = 0.5226437846142446 - (0.5226437846142446 * 10 / 100) = 0.4703794061528194 ml
At the end of Day 29, the content drops below the threshold of 0.5 ml, so the function returns 29.
*/
func Evaporator(content float64, evapPerDay int, threshold int) int {
	n := 0
	thresholdContent := content * float64(threshold) / 100
	for content > thresholdContent {
		content -= content * float64(evapPerDay) / 100
		n++
	}
	return n
}

func Evaporator2(content float64, evapPerDay int, threshold int) int {
	base := 1.0 - float64(evapPerDay)/100.0
	top := float64(threshold) / 100.0
	N := math.Log(top) / math.Log(base)

	return int(math.Ceil(N))
}

func Evaporator3(_ float64, evapPerDay int, threshold int) int {
	return int(math.Ceil(math.Log(float64(threshold)/100) / math.Log(1-float64(evapPerDay)/100)))
}
