package kata8

/*
Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter; and month 11 (November), is part of the fourth quarter.

Constraint:

1 <= month <= 12*/
func QuarterOf(month int) int {
	if month >= 1 && month <= 3 {
		return 1
	} else if month >= 4 && month <= 6 {
		return 2
	} else if month >= 7 && month <= 9 {
		return 3
	} else if month >= 10 && month <= 12 {
		return 4
	} else {
		panic("invalid month")
	}
}

func QuarterOf2(month int) int {
	switch month {
		case 1,2,3:
			return 1
		case 4,5,6:
			return 2
		case 7,8,9:
			return 3
		case 10,11,12:
			return 4
	}

	return 0
}

func QuarterOf3(month int) int {
	quarters := [12]int{1,1,1,2,2,2,3,3,3,4,4,4}

	return quarters[month-1]
}

func QuarterOf4(month int) int {
	var quarters map[int]int = map[int]int{
	  1:1,
	  2:1,
	  3:1,
	  4:2,
	  5:2,
	  6:2,
	  7:3,
	  8:3,
	  9:3,
	  10:4,
	  11:4,
	  12:4,
  }
	return quarters[month]
}

func QuarterOf5(month int) int {
	return (month + 2) / 3
}