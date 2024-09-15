package kata7

/*
#Vowel Count

Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/
func GetCount(str string) (count int) {
	count = 0
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, char := range str {
		for _, vowel := range vowels {
			if string(char) == vowel {
				count++
			}
		}
	}
	return count
}

func GetCount1(str string) (count int) {
	for _, c := range str {
	  switch c {
	  case 'a', 'e', 'i', 'o', 'u':
		count++
	  }
	}
	return count
}

func GetCount2(str string) (count int) {
	m := map[byte]int{
	  'a': 1,
	  'e': 1,
	  'i': 1,
	  'o': 1,
	  'u': 1,
	}
	for i := range str {
	  count += m[str[i]]
	}
	return count
}