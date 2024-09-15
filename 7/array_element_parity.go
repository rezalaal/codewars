package kata7

import "sort"

/*
Array element parity
In this Kata, you will be given an array of integers whose elements have both a negative and a positive value,
except for one integer that is either only negative or only positive. Your task will be to find that integer.

Examples:

[1, -1, 2, -2, 3] => 3

3 has no matching negative appearance

[-3, 1, 2, 3, -1, -4, -2] => -4

-4 has no matching positive appearance

[1, -1, 2, -2, 3, 3] => 3

(the only-positive or only-negative integer may appear more than once)

Good luck!
*/
func ArrayElementParity(arr []int) int {
	
	for _, n := range arr {
		flag := false
		m := -n
		for _, number := range arr {
			if number == m {
				flag = true
			}
		}
		if !flag {
			return n
		}
	}
	return 0
}

func ArrayElementParity2(arr []int) int {
	hash := make(map[int]bool)
	ans := 0
	for _, entry := range arr {
		if _, value := hash[entry]; !value {
			hash[entry] = true
			ans += entry
		}
	}
	return ans
}

func ArrayElementParity3(arr []int) int {
	sort.Ints(arr)
  for i := 0; i < len(arr) / 2; i ++ {
    add := arr[i] + arr[len(arr) - i - 1]
    switch {
      case add > 0: return arr[len(arr) - i - 1]
      case add < 0: return arr[i]
    }
  }
  return arr[len(arr) / 2]
}