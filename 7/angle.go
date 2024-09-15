package kata7

/*
#Sum of angles

Find the total sum of internal angles (in degrees) in an n-sided simple polygon. N will be greater than 2.
*/
func Angle(n int) int {
	return (n - 2) * 180
}
