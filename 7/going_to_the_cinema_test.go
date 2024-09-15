package kata7

import "testing"

func TestMovie(t *testing.T) {
	testCases := []struct {
		card   int
		ticket int
		perc   float64
		out    int
	}{
		{500, 15, 0.9, 43},
		{0, 10, 0.95, 2},
	}

	for _, test := range testCases {
		got := Movie(test.card, test.ticket, test.perc)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
