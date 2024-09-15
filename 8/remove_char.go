package kata8

// import "fmt"

/*
#Remove First and Last Character

It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string.
You're given one parameter, the original string. You don't have to worry about strings with less than two characters.
*/
func RemoveChar(word string) string {
	l := len(word)

	result := ""
	for i, char := range word {
		if i == 0 || i == l-1 {
			continue
		}
		result += string(char)
	}
	return result
}

func RemoveChar2(word string) string {
	if len(word) == 0 || len(word) == 1 {
		return ""
	}
	return word[1:len(word)-1]
}

func RemoveChar3(word string) string {
	if len(word) == 0 || len(word) == 1 {
		return ""
	}
	var newWord = []rune(word)
	return string(newWord[1:len(newWord) - 1])
}


func RemoveChar4(word string) string {
	var str string
	for i := 1; i < len(word)-1; i++ {
		str += string(word[i])
	}
	return str
}