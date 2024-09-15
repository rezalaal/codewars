package kata7

import (
	"fmt"
	"strings"
)

/*
#Parts of a list

Write a function partlist that gives all the ways to divide a list (an array) of at least two elements into two non-empty parts.

Each two non empty parts will be in a pair (or an array for languages without tuples or a structin C - C: see Examples test Cases - )
Each part will be in a string
Elements of a pair must be in the same order as in the original array.
Examples of returns in different languages:
a = ["az", "toto", "picaro", "zone", "kiwi"] -->
[["az", "toto picaro zone kiwi"], ["az toto", "picaro zone kiwi"], ["az toto picaro", "zone kiwi"], ["az toto picaro zone", "kiwi"]] 
or
 a = {"az", "toto", "picaro", "zone", "kiwi"} -->
{{"az", "toto picaro zone kiwi"}, {"az toto", "picaro zone kiwi"}, {"az toto picaro", "zone kiwi"}, {"az toto picaro zone", "kiwi"}}
or
a = ["az", "toto", "picaro", "zone", "kiwi"] -->
[("az", "toto picaro zone kiwi"), ("az toto", "picaro zone kiwi"), ("az toto picaro", "zone kiwi"), ("az toto picaro zone", "kiwi")]
or 
a = [|"az", "toto", "picaro", "zone", "kiwi"|] -->
[("az", "toto picaro zone kiwi"), ("az toto", "picaro zone kiwi"), ("az toto picaro", "zone kiwi"), ("az toto picaro zone", "kiwi")]
or
a = ["az", "toto", "picaro", "zone", "kiwi"] -->
"(az, toto picaro zone kiwi)(az toto, picaro zone kiwi)(az toto picaro, zone kiwi)(az toto picaro zone, kiwi)"
Note
You can see other examples for each language in "Your test cases"
*/
func PartList(arr []string) string {
    var result strings.Builder

	for i := 1; i < len(arr); i++ {
		firstPart := strings.Join(arr[:i], " ")
		secondPart := strings.Join(arr[i:], " ")
		result.WriteString(fmt.Sprintf("(%s, %s)", firstPart, secondPart))
	}

	return result.String()
}

func PartList2(arr []string) (result string) {
	for idx := 1; idx < len(arr); idx++ {
		firstHalf := strings.Join(arr[:idx], " ")
		secondHalf := strings.Join(arr[idx:], " ")
		variations := fmt.Sprintf("%s, %s", firstHalf, secondHalf)
		result += fmt.Sprintf("(%s)", variations)
	}
	return
}

func PartList3(arr []string) string {
	var result []string
	for i := 1; i < len(arr); i++ {
		parts := fmt.Sprintf("(%s, %s)", strings.Join(arr[:i], " "), strings.Join(arr[i:], " "))
		result = append(result, parts)
	}
	return strings.Join(result, "")
}