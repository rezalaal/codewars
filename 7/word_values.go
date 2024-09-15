package kata7

import "strings"

/*
Word values
Given a string "abc" and assuming that each letter in the string has a value equal to its position in the alphabet,
our string will have a value of 1 + 2 + 3 = 6. This means that: a = 1, b = 2, c = 3 ....z = 26.

You will be given a list of strings and your task will be to return the values of the strings as explained above
multiplied by the position of that string in the list. For our purpose, position begins with 1.

nameValue ["abc","abc abc"] should return [6,24] because of [ 6 * 1, 12 * 2 ]. Note how spaces are ignored.

"abc" has a value of 6, while "abc abc" has a value of 12.
Now, the value at position 1 is multiplied by 1 while the value at position 2 is multiplied by 2.

Input will only contain lowercase characters and spaces.

Good luck!
*/
func NameValue(my_list []string) []int {
	res := []int{}
	for i, word := range my_list {
		sum := 0
		for _, char := range word {
			if char != ' ' {
				sum += int(char - 'a' + 1)
			}
		}
		res = append(res, sum*(i+1))
	}
	return res
}
func NameValue1(my_list []string) []int {
	var result = make([]int, len(my_list))

	for idx, str := range my_list {
		for _, chr := range str {
			if chr >= 'a' && chr <= 'z' {
				result[idx] += int(chr-'a') + 1
			}
		}

		result[idx] = result[idx] * (idx + 1)
	}

	return result
}

func NameValue2(my_list []string) []int {
	var res []int
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, word := range my_list {
		temp := 0
		for _, c := range word {
			temp += strings.Index(alphabet, string(c)) + 1
		}
		res = append(res, temp*(i+1))
	}
	return res
}