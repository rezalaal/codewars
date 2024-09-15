package kata8

import "strings"

/*
#Abbreviate a Two Word Name

Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

patrick feeney => P.F
*/
func AbbrevName(name string) string {
	parts := strings.Split(name, " ")
	return strings.ToUpper(string(parts[0][0]))+"."+strings.ToUpper(string(parts[1][0]))
}