package kata7

import "testing"

func TestNbYear(t *testing.T) {
	testCases := []struct {
		p0      int
		percent float64
		aug     int
		p       int
		out     int
	}{
		{1000, 2, 50, 1213, 3},
		{1500, 5, 100, 5000, 15},
		{1500000, 2.5, 10000, 2000000, 10},
		{1500000, 0.25, 1000, 2000000, 94},
		{1500000, 0.25, -1000, 2000000, 151},
	}

	for _, test := range testCases {
		got := NbYear(test.p0, test.percent, test.aug, test.p)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = NbYear2(test.p0, test.percent, test.aug, test.p)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = NbYear3(test.p0, test.percent, test.aug, test.p)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
