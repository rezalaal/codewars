package kata8

/*
#Cat years, Dog years

Kata Task
I have a cat and a dog.

I got them at the same time as kitten/puppy. That was humanYears years ago.

Return their respective ages now as [humanYears,catYears,dogYears]

NOTES:

humanYears >= 1
humanYears are whole numbers only
Cat Years
15 cat years for first year
+9 cat years for second year
+4 cat years for each year after that
Dog Years
15 dog years for first year
+9 dog years for second year
+5 dog years for each year after that
References

http://www.catster.com/cats-101/calculate-cat-age-in-cat-years
http://www.slate.com/articles/news_and_politics/explainer/2009/05/a_dogs_life.html
If you liked this Kata there is another related one here
*/
func CalculateYears(years int) (result [3]int) {

	catYears := 0
	dogYears := 0
	for i := 1; i <= years; i++ {
		if i == 1 {
			catYears += 15
			dogYears += 15
		} else if i == 2 {
			catYears += 9
			dogYears += 9
		} else {
			catYears += 4
			dogYears += 5
		}

	}
	result[0] = years
	result[1] = catYears
	result[2] = dogYears
	return result
}

func CalculateYears1(years int) (result [3]int) {
	switch years {
	  case 1: result = [3]int{1, 15, 15}
	  case 2: result = [3]int{2, 24, 24}
	  default: result = [3]int{years, 24 + 4 * (years - 2), 24 + 5 * (years - 2)}   
	}
	return 
}