package kata7

import "regexp"
/*
Longest vowel chain
The vowel substrings in the word codewarriors are o,e,a,io. The longest of these has a length of 2.
Given a lowercase string that has alphabetic characters only 
(both vowels and consonants) and no spaces, return the length of the longest vowel substring. 
Vowels are any of aeiou.

Good luck!
*/
func LongestVowelChain(s string) int {
	count := 0
	max := 0
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
			if count > max {
				max = count
			}
		}else{
			count = 0
		}

	}
	return max
}
func LongestVowelChain1(s string) int {
	n, max := 0, 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			if n++; n > max {
				max = n
			}
		default:
			n = 0
		}
	}
	return max
}

func LongestVowelChain2(s string) int {
    max := 0
    for _,sub := range regexp.MustCompile(`[aeiou]+`).FindAllString(s,-1) {
        if len(sub) > max {
            max = len(sub)
        }
    }
    return max
}