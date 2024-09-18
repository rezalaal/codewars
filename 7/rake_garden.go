package kata7

import (
	"regexp"
	"strings"
)

/*
Help Suzuki rake his garden!
Help Suzuki rake his garden!

The monastery has a magnificent Zen garden made of white gravel and rocks and it is raked diligently everyday
by the monks. Suzuki having a keen eye is always on the lookout for anything creeping into the garden that must be
removed during the daily raking such as insects or moss.

You will be given a string representing the garden such as:

garden = 'gravel gravel gravel gravel gravel gravel gravel gravel gravel rock slug ant gravel gravel snail
rock gravel gravel gravel gravel gravel gravel gravel slug gravel ant gravel gravel gravel gravel rock slug
gravel gravel gravel gravel gravel snail gravel gravel rock gravel snail slug gravel gravel spider gravel gravel
gravel gravel gravel gravel gravel gravel moss gravel gravel gravel snail gravel gravel gravel ant gravel gravel
moss gravel gravel gravel gravel snail gravel gravel gravel gravel slug gravel rock gravel gravel rock gravel gravel
gravel gravel snail gravel gravel rock gravel gravel gravel gravel gravel spider gravel rock gravel gravel'
Rake out any items that are not a rock or gravel and replace them with gravel such that:

garden = 'slug spider rock gravel gravel gravel gravel gravel gravel gravel'
Returns a string with all items except a rock or gravel replaced with gravel:

garden = 'gravel gravel rock gravel gravel gravel gravel gravel gravel gravel'
*/
func RakeGarden(garden string) string {
	items := strings.Split(garden, " ")
	result := []string{}
	for _, item := range items {
		if item != "gravel" && item != "rock" {
			result = append(result, "gravel")
		}else{
			result = append(result, item)
		}
	}
	return strings.Join(result, " ")
}

func RakeGarden1(garden string) string {
	return regexp.MustCompile(`\S+`).ReplaceAllStringFunc(garden, func(s string) string { if s == "rock" { return "rock" }; return "gravel" })
}