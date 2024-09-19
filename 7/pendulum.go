package kata7

import "sort"

/*
The Poet And The Pendulum
Scenario
the rhythm of beautiful musical notes is drawing a Pendulum

# Beautiful musical notes are the Numbers

Task
Given an array/list [] of n integers , Arrange them in a way similar to the to-and-fro movement of a Pendulum

The Smallest element of the list of integers , must come in center position of array/list.

The Higher than smallest , goes to the right .
The Next higher number goes to the left of minimum number and So on , in a to-and-fro manner similar to that of a Pendulum.

!alt

Notes
Array/list size is at least 3 .

In Even array size , The minimum element should be moved to (n-1)/2 index (considering that indexes start from 0)

Repetition of numbers in the array/list could occur , So (duplications are included when Arranging).

Input >> Output Examples:
pendulum ([6, 6, 8 ,5 ,10]) ==> [10, 6, 5, 6, 8]
Explanation:
Since , 5 is the The Smallest element of the list of integers , came in The center position of array/list

The Higher than smallest is 6 goes to the right of 5 .

The Next higher number goes to the left of minimum number and So on .

Remember , Duplications are included when Arranging , Don't Delete Them .

pendulum ([-9, -2, -10, -6]) ==> [-6, -10, -9, -2]
Explanation:
Since , -10 is the The Smallest element of the list of integers , came in The center position of array/list

The Higher than smallest is -9 goes to the right of it .

The Next higher number goes to the left of -10 , and So on .

Remeber , In Even array size , The minimum element moved to (n-1)/2 index (considering that indexes start from 0) .

pendulum ([11, -16, -18, 13, -11, -12, 3, 18]) ==> [13, 3, -12, -18, -16, -11, 11, 18]
Explanation:
Since , -18 is the The Smallest element of the list of integers , came in The center position of array/list

The Higher than smallest is -16 goes to the right of it .

The Next higher number goes to the left of -18 , and So on .

Remember , In Even array size , The minimum element moved to (n-1)/2 index (considering that indexes start from 0) .
*/
func Pendulum(values []int) []int {
	if len(values) % 2 == 0 {
		p := MakeValley(values)
		res:=[]int{}
		for i:=len(p)-1;i>=0;i-- {
			res = append(res, p[i])
		}
		return res
	}

	return MakeValley(values)
}

func Pendulum1(values []int) (dst []int) {
   sort.Ints(values)
   for i, v := range values {
       if i % 2 == 0 {
           dst = append([]int{v}, dst...)
       } else {
           dst = append(dst, v)
       }
    }
    
    return
}

func Pendulum2(values []int) []int {
	sort.Ints(values)
	
	var newValues = make([]int, len(values))
	cur := (len(values)-1)/2
 
	direction:=1
	for i:=0; i<len(values); i++ {
	  newValues[cur] = values[i] 
	  cur += (i+1)*direction
	  direction = -direction
	}
	return newValues
}	