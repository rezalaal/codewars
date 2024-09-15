package kata7

import "testing"

func TestUnluckyDays(t *testing.T) {
	testCases := []struct {
		year int
		out  int
	}{
		{2015, 3},
		{1986, 1},
	}

	for _, test := range testCases {
		got := UnluckyDays(test.year)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
