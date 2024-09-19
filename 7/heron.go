package kata7

import "math"

/*
Heron's formula
Write function heron which calculates the area of a triangle with sides a, b, and c (x, y, z in COBOL).

Heron's formula:

sqr(s∗(s−a)∗(s−b)∗(s−c))
​

where
s=  a+b+c / 2
​

Output should have 2 digits precision.
*/
func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2

	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	
	return math.Round(area*100) / 100
}
func Heron1(a, b, c float64) (area float64) {
	area = (a + b + c) / 2
	area = math.Sqrt(area * (area - a) * (area - b) * (area - c))
	return
}