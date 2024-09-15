package kata8


/*#I love you, a little , a lot, passionately ... not at all

Who remembers back to their time in the schoolyard, when girls would take a flower and tear its petals, saying each of the following phrases each time a petal was torn:

1. "I love you"
2. "a little"
3. "a lot"
4. "passionately"
5. "madly"
6. "not at all"
If there are more than 6 petals, you start over with "I love you" for 7 petals, "a little" for 8 petals and so on.

When the last petal was torn there were cries of excitement, dreams, surging thoughts and emotions.

Your goal in this kata is to determine which phrase the girls would say at the last petal for a flower of a given number of petals. The number of petals is always greater than 0.


*/
func HowMuchILoveYou(i int) string {
	switch (i-1) % 6 {
		case 0:	return "I love you"
		case 1:  return "a little"
		case 2:  return "a lot"
		case 3:  return "passionately"
		case 4:  return "madly"
		case 5:  return "not at all"
		default: return "I love you"
	}	 
}

func HowMuchILoveYou2(i int) string {
    return []string{"I love you","a little","a lot","passionately","madly","not at all"}[(i-1) % 6]
}

func HowMuchILoveYou3(i int) string {
	d := 6
	var arr = map[int]string{
		1: "I love you",
		2: "a little",
		3: "a lot",
		4: "passionately",
		5: "madly",
		0: "not at all",
	}
	return arr[i%d]
}