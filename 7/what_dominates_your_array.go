package kata7


/*
What dominates your array?

	zero-indexed array arr consisting of n integers is given. The dominator of array arr is the value that occurs
	 in more than half of the elements of arr.

For example, consider array arr such that arr = [3,4,3,2,3,1,3,3]
The dominator of arr is 3 because it occurs in 5 out of 8 elements of arr and 5 is more than a half of 8.
Write a function dominator(arr) that, given a zero-indexed array arr consisting of n integers, returns the
dominator of arr. The function should return −1 if array does not have a dominator. All values in arr will be >=0.
*/
func Dominator(a []int) int {
	h := len(a) / 2	
	for i, number := range a {	
		count := 0	
		for j := i; j < len(a); j++ {
			if number == a[j] {
				count++
			}
			if count > h {
				return a[j]
			}
		}
	}
	return -1
}

func Dominator1(arr []int) int {
    if len(arr) == 0 {
        return -1 // Return -1 for empty array
    }

    countMap := make(map[int]int)
    halfLength := len(arr) / 2

    for _, num := range arr {
        countMap[num]++
        if countMap[num] > halfLength {
            return num // Return the dominator as soon as it's found
        }
    }

    return -1 // Return -1 if no dominator is found
}