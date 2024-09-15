package kata8

/*
#Third Angle of a Triangle

You are given two interior angles (in degrees) of a triangle.

Write a function to return the 3rd.

Note: only positive integers will be tested.

https://en.wikipedia.org/wiki/Triangle
*/
func OtherAngle(a int, b int) int {
	return 180 - (a + b)
}
