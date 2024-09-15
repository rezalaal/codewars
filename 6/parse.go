package kata6

/*
#Make the Deadfish Swim

Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.

Parse("iiisdoso") == []int{8, 64}
*/
func Parse(data string) []int {
	result := []int{}
	num := 0
	for _, char := range data {
		switch string(char) {
		case "i":
			num++
		case "d":
			num--
		case "s":
			num *= num
		case "o":
			result = append(result, num)
		}
	}
	return result
}