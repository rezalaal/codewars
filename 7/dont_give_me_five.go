package kata7

import (
	"strconv"
	"strings"
)

/*
#Don't give me five!

Don't give me five!
In this kata you get the start number and the end number of a region and should return the count of all numbers except numbers with a 5 in it. The start and the end number are both inclusive!

Examples:

1,9 -> 1,2,3,4,6,7,8,9 -> Result 8
4,17 -> 4,6,7,8,9,10,11,12,13,14,16,17 -> Result 12
The result may contain fives. ;-)
The start number will always be smaller than the end number. Both numbers can be also negative!

I'm very curious for your solutions and the way you solve it. Maybe someone of you will find an easy pure mathematics solution.

Have fun coding it and please don't forget to vote and rank this kata! :-)

I have also created other katas. Take a look if you enjoyed this kata!
*/
func DontGiveMeFive(start int, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if !strings.Contains(strconv.Itoa(i),"5") {
			count++
		}
	}
	return count
}


func DontGiveMeFive1(start, end int) int {
    n := end - start + 1
    for i := start; i<=end; i++ {
        d := i
        if i < 0 { d = -i }
        for d > 0 {
            if d % 10 == 5 {
                n--
                break
            }
            d /= 10
        }
    }
    return n
}

func DontGiveMeFive2(start int, end int) int {
	var cnt int
	for i := start; i<=end;i++ {
	  tmp := i
	  if tmp < 0 {
		tmp = -tmp
	  }
	  
	  if countainFive(tmp) {
		continue
	  }
	  cnt++
	}
	
	return cnt
}

func countainFive(n int) bool {
	for n != 0 {
	  if val := n % 10; val == 5 {
		return true
	  }
	  n /= 10
	}
	
	return false
}

func DontGiveMeFive3(start int, end int) (res int) {
	for i := start; i <= end; i++ {
	  for _, num := range (strconv.Itoa(i)) {
		if num == '5' {
		  res += 1
		  break
		}
	  }
	}
	return end - start - res + 1
}
