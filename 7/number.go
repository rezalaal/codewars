package kata7

/*
#Number of People in the Bus

There is a bus moving in the city which takes and drops some people at each bus stop.

You are provided with a list (or array) of integer pairs. Elements of each pair represent the number of people that get on the bus (the first item) and the number of people that get off the bus (the second item) at a bus stop.

Your task is to return the number of people who are still on the bus after the last bus stop (after the last array). Even though it is the last bus stop, the bus might not be empty and some people might still be inside the bus, they are probably sleeping there :D

Take a look on the test cases.

Please keep in mind that the test cases ensure that the number of people in the bus is always >= 0. So the returned integer can't be negative.

The second value in the first pair in the array is 0, since the bus is empty in the first bus stop.
*/
func Number(busStops [][2]int) int {	
	input := 0
	output := 0
	for _, stop :=range busStops {
		input += stop[0]
		output += stop[1]
	}
	return input - output
}

func Number2(busStops [][2]int) (inBus int) {
	for _, stopInfo := range busStops {
	  inBus += stopInfo[0] - stopInfo[1]
	}  
	return // Good Luck!
}

func Number3(busStops [][2]int) int {
	peopleInTheBus := 0;
	
	for _, busStop := range busStops {
	  peopleInTheBus += busStop[0] - busStop[1]
	}
	
	return peopleInTheBus
}