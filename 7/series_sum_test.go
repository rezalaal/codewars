package kata7

import "testing"

func TestSeriesSum(t *testing.T) {
	testCases :=[]struct{
		n int
		out string
	}{
		{1, "1.00"}, //"n = 1"
		{2, "1.25"}, //"n = 2"
		{3, "1.39"}, //"n = 3"
		{4, "1.49"}, //"n = 4"
		{5, "1.57"}, //"n = 4"
	}

	for _, test := range testCases {
		got := SeriesSum(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}