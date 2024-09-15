package kata7

import (
	"strings"
)

/*
#Alphabet symmetry

Consider the word "abode". We can see that the letter a is in position 1 and b is in position 2. In the alphabet, a and b are also in positions 1 and 2. Notice also that d and e in abode occupy the positions they would occupy in the alphabet, which are positions 4 and 5.

Given an array of words, return an array of the number of letters that occupy their positions in the alphabet for each word. For example,

solve(["abode","ABc","xyzD"]) = [4, 3, 1]
See test cases for more examples.

Input will consist of alphabet characters, both uppercase and lowercase. No spaces.

Good luck!

Logic:
Iterate over each word and its index.
For each character in the word, calculate its position in the alphabet by converting it to lowercase and finding its ASCII value.
Compare the calculated position with the index of the character (adjusted for 1-based indexing).
If they match, increment the count for that word.
Store the result: After processing each word, store the count in the results slice.
*/
func Solve(slice []string) []int {
	results := make([]int, len(slice))

	for i, word := range slice {
		count := 0
		for j, char := range word {			
			position := int(strings.ToLower(string(char))[0]) - 'a' + 1
			if position == j+1 {
				count++
			}
		}
		results[i] = count
	}
	return results
}