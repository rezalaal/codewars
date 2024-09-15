package kata8

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	result := ""
	for i:=0;i<len(arr[0]);i++ {
		if i != len(arr[0])-1 {
			result += string(arr[0][i])	+ "***"
		}else{
			result += string(arr[0][i])	
		}
			
	}
	return result
}


func TwoSort1(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}

func TwoSort2(arr []string) string {
	sort.Strings(arr)
	return strings.Join(strings.Split(arr[0],""), "***")
}