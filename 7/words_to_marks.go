package kata7

/*
#Love vs friendship

If　a = 1, b = 2, c = 3 ... z = 26

Then l + o + v + e = 54

and f + r + i + e + n + d + s + h + i + p = 108

So friendship is twice as strong as love :-)

Your task is to write a function which calculates the value of a word based off the sum of the alphabet positions of its characters.

The input will always be made of only lowercase letters and will never be empty.
*/
func WordsToMarks(s string) int {
	alphabets := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	sum := 0
	for _, char := range s {
		sum += alphabets[string(char)]
	}
	return sum 
}


func WordsToMarks2(s string) int {
	count := 0
	for _, i := range s {
	   count += int(i) - 'a' + 1;
	}
	return count
}


func WordsToMarks3(s string) int {
    var result int
    for _, i := range(s) {
      result += int((i-96))
    }
    return result
}


func WordsToMarks4(s string) int {
    t := map[rune]int{}
    for i, c := range "abcdefghijklmnopqrstuvwxyz" { t[c] = i + 1 }
    
    res := 0
    for _, c := range s { res += t[c] }
    
    return res
}

func WordsToMarks5(s string) int {
	abc := "abcdefghijklmnopqrstuvwxyz"
	m := make(map[rune]int)
	for i, c := range abc {
	  m[c] = i+1
	}
	result := 0
	for _, c := range s {
	  result += m[c]
	}
	return result
}