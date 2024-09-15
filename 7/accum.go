package kata7


import (
	"strings"
)


/*
#Mumbling

This time no story, no theory. The examples below show you how to write function accum:

Examples:
accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
The parameter of accum is a string which includes only letters from a..z and A..Z.
*/
func Accum(s string) string {
	result := ""
	for i, char := range s {
		for j:=0;j<=i;j++ {
			if j == 0 {
				result += strings.ToUpper(string(char))
			}else{
				result += strings.ToLower(string(char))
			}			
		}
		if i<len(s) -1 {result += "-"}
	}
	return result
}

func Accum2(s string) string {
    parts := make([]string, len(s))
    
    for i := 0; i < len(s); i++ {
      parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
    }
    
    return strings.Join(parts, "-")
}

func Accum3(s string) string {
    words := make([]string, len(s))
    
    for i, c := range s {
        words[i] = strings.Title(strings.Repeat(strings.ToLower(string(c)), i+1))
    }
   
    return strings.Join(words, "-")
}