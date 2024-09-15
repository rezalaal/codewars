package kata8

/*
#Convert boolean values to strings 'Yes' or 'No'.

Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.
*/
func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}
	return "No"
}

func BoolToWord2(word bool) string {
	return map[bool]string{false:"No", true:"Yes"}[word]
}