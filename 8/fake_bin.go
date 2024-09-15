package kata8

import (
	"strconv"
	"strings"
)

/*
#Fake Binary

Given a string of digits, you should replace any digit below 5 with '0' and any digit 5 and above with '1'. 
Return the resulting string.

Note: input will never be an empty string
*/
func FakeBin(x string) string {
	res := ""
	for _, value := range x {
		number, _ := strconv.Atoi(string(value))
		if number < 5 {
			res += "0"
		} else {
			res += "1"
		}
	}

	return res
}

func FakeBin2(x string) (result string) {
    for _, char := range x {
        if char < '5' {
            result += "0"
        } else {
            result += "1"
        }
    }
    return
}

func FakeBin3(x string) (r string) {

	res := strings.NewReplacer("1", "0", "2", "0", "3", "0", "4", "0", "5", "1", "6", "1", "7", "1", "8", "1", "9", "1")

	r = res.Replace(x)
	return
}

func FakeBin4(x string) string {
	x = strings.Replace(x, "1", "0", -1)
	x = strings.Replace(x, "2", "0", -1)
	x = strings.Replace(x, "3", "0", -1)
	x = strings.Replace(x, "4", "0", -1)
	x = strings.Replace(x, "5", "1", -1)
	x = strings.Replace(x, "6", "1", -1)
	x = strings.Replace(x, "7", "1", -1)
	x = strings.Replace(x, "8", "1", -1)
	x = strings.Replace(x, "9", "1", -1)
	  return x
}