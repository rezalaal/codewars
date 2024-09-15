package kata8

/*
#Grasshopper - Grade book

Grade book
Complete the function so that it finds the average of the three scores passed to it and returns the letter value associated with that grade.

Numerical Score	Letter Grade
90 <= score <= 100	'A'
80 <= score < 90	'B'
70 <= score < 80	'C'
60 <= score < 70	'D'
0 <= score < 60	'F'
Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.
*/
func GetGrade(a, b, c int) rune {
	average := (a + b + c) / 3
	if average <= 100 && average >= 90 {
		return 'A'
	}
	if average < 90 && average >= 80 {
		return 'B'
	}
	if average < 80 && average >= 70 {
		return 'C'
	}
	if average < 70 && average >= 60 {
		return 'D'
	}
	if average < 60 && average >= 0 {
		return 'F'
	}
	return 'F'
}

func GetGrade1(a, b, c int) rune {
	switch mean := (a + b + c) / 3; {
	case mean < 60:
		return 'F'
	case mean < 70:
		return 'D'
	case mean < 80:
		return 'C'
	case mean < 90:
		return 'B'
	default:
		return 'A'
	}

}
