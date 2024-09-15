package kata7

/*
#Alphabetical Addition

Your task is to add up letters to one letter.

The function will be given a slice of letters (runes), each one being a letter to add. Return a rune.

Notes:
Letters will always be lowercase.
Letters can overflow (see second to last example of the description)
If no letters are given, the function should return 'z'
Examples:
AddLetters([]rune{'a', 'b', 'c'}) = 'f'
AddLetters([]rune{'a', 'b'}) = 'c'
AddLetters([]rune{'z'}) = 'z'
AddLetters([]rune{'z', 'a'}) = 'a'
AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
AddLetters([]rune{}) = 'z'
*/
func AddLetters(letters []rune) rune {

	if len(letters) == 0 {
		return 'z'
	}
	total := 0
	for _, ch := range letters {
		total += int(ch - 'a' + 1)
	}
	total = (total - 1) % 26
	if total < 0 {
		total += 26
	}

	return rune('a' + total)
}

func AddLetters1(letters []rune) rune {
	if len(letters) == 0 {
		return 'z'
	}

	var sum = 0
	for _, letter := range letters {
		sum += int(letter) - 96
	}

	charCode := (sum-1)%26 + 97
	return rune(charCode)
}

var alphabet = map[int]rune{
	1: 'a', 2: 'b', 3: 'c', 4: 'd', 5: 'e', 6: 'f', 7: 'g', 8: 'h', 9: 'i', 10: 'j', 11: 'k', 12: 'l', 13: 'm', 14: 'n', 15: 'o', 16: 'p', 17: 'q', 18: 'r', 19: 's', 20: 't', 21: 'u', 22: 'v', 23: 'w', 24: 'x', 25: 'y', 26: 'z',
}

func AddLetters2(letters []rune) rune {
	count := 0
	alphabet1 := "abcdefghijklmnopqrstuvwxyz"
	if len(letters) == 0 {
		return alphabet[26]
	}
	for i, char := range alphabet1 {

		for _, ch := range letters {
			if string(char) == string(ch) {
				count = count + i + 1

			}
		}

	}

	if count <= 26 {
		return alphabet[count]
	} else if count > 26 && count < 51 {
		return alphabet[count-26]
	} else if count%26 == 0 {
		return 122
	} else {
		return alphabet[count-count/26*26]
	}
	return 'o'
}
