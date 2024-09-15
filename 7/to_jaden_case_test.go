package kata7


import "testing"

func TestToJadenCase(t *testing.T) {
	testCases := []struct{
		str string
		out string
	}{		
		{"most trees are blue",                                                                     "Most Trees Are Blue"},
		{"All the rules in this world were made by someone no smarter than you. So make your own.", "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."},
		{"When I die. then you will realize",                                                       "When I Die. Then You Will Realize"},
		{"Jonah Hill is a genius",                                                                  "Jonah Hill Is A Genius"},
		{"Dying is mainstream",                                                                     "Dying Is Mainstream"},
	}

	for _, test := range testCases {
		got := ToJadenCase(test.str)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = ToJadenCase1(test.str)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}