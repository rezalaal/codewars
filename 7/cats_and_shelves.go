package kata7


/*
Cats and shelves
Description
An infinite number of shelves are arranged one above the other in a staggered fashion.
The cat can jump either one or three shelves at a time:
from shelf i to shelf i+1 or i+3 (the cat cannot climb on the shelf directly above its head),
according to the illustration:

                 ┌────────┐
                 │-6------│
                 └────────┘
┌────────┐
│------5-│
└────────┘  ┌─────► OK!
            │    ┌────────┐
            │    │-4------│
            │    └────────┘
┌────────┐  │
│------3-│  │
BANG!────┘  ├─────► OK!
  ▲  |\_/|  │    ┌────────┐
  │ ("^-^)  │    │-2------│
  │ )   (   │    └────────┘
┌─┴─┴───┴┬──┘
│------1-│
└────────┘
Input
Start and finish shelf numbers (always positive integers, finish no smaller than start)

Task
Find the minimum number of jumps to go from start to finish

Example
Start 1, finish 5, then answer is 2 (1 => 4 => 5 or 1 => 2 => 5)

Inspirers
https://i.ibb.co/BymvZtL/Inspirers.jpg
*/
func Cats(start, finish int) int {
	count := 0
	for finish - start >= 3 {
		start +=  3		
		count ++
	}
	for finish - start > 2 {
		start += 2		
		count++
	}
	for finish - start > 0 {
		start ++		
		count++
	}
	return count;
}

func Cats2(start, finish int) int {
	a := finish - start
	return a / 3 + a % 3
}

const max_jumps = 3

func Cats3(start, finish int) int {
  shelves := finish - start  
  return (shelves/max_jumps) + shelves % max_jumps;
}

func Cats4(start, finish int) (result int) {
	for start < finish {
	  if start + 3 <= finish {
		start += 3
	  } else if start + 1 <= finish {
		start += 1
	  }
	  
	  result++
	}
	
	return result
}

func Cats5(start, finish int) int {
	//Mew - 1 or 3 jumps not 2
	diff := finish - start
	modDiff := diff % 3
	divDiff := diff / 3
	
	return divDiff + modDiff
}

func Cats6(start, finish int) int {
	switch finish-start {  
	  case 0,1,2: return finish - start 
	  case 3: return 1
	}
	return 1 + Cats(start+3, finish)
}