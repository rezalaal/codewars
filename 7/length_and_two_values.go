package kata7


/*
Length and two values
Given an integer n and two other values, build an array of size n filled with these two values alternating.

Examples
5, true, false     -->  [true, false, true, false, true]
10, "blue", "red"  -->  ["blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"]
0, "one", "two"    -->  []
Good luck!
*/
func Alternate(n int, firstValue string, secondValue string) []string {
	result := []string{}
	for i:=1;i<=n;i++ {
		if i % 2 == 0 {
			result = append(result, secondValue)
		}else{
			result = append(result, firstValue)
		}
	}
	return result
} 
func Alternate1[T any](n int, x T, y T) []T {
	res := make([]T, n)
	xy := [2]T{x, y}
	for i := 0; i < n; i++ {
	  res[i] = xy[i % 2]
	}
	return res
}