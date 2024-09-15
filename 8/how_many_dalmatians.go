package kata8

/*
#101 Dalmatians - squash the bugs, not the dogs!

Your friend has been out shopping for puppies (what a time to be alive!)... He arrives back with multiple dogs, and you simply do not know how to respond!

By repairing the function provided, you will find out exactly how you should respond, depending on the number of dogs he has.

The number of dogs will always be a number and there will always be at least 1 dog.

Good luck!

	dogs : [string("Hardly any", "More than a handful!", "Woah that's a lot of dogs!" "101 DALMATIONS!!!")

	if number <= 10; dogs[0]
	if number =< 50:, dogs[1]
	if number = 101 : dogs[3]
	dogs[2]
*/
func HowManyDalmatians(number int) string {

	dogs := []string{"Hardly any", "More than a handful!", "Woah that's a lot of dogs!", "101 DALMATIONS!!!"}

	if number <= 10 {
		return dogs[0]
	} else if number <= 50 {
		return dogs[1]
	} else if number == 101 {
		return dogs[3]
	} else {
		return dogs[2]
	}
}

// Clever solution

func HowManyDalmatians1(number int) string {
	switch {
	case number==101:
	  return "101 DALMATIONS!!!"
	case number>50:
	  return "Woah that's a lot of dogs!"
	case number>10:
	  return "More than a handful!"
	default:
	   return "Hardly any"
	}
}

