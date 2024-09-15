package kata7

import "testing"

func TestNumber(t *testing.T) {
	testCases := []struct {
		busStops [][2]int
		out      int
	}{
		{[][2]int{{10, 0}, {3, 5}, {5, 8}}, 5},
		{[][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}, 17},
		{[][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}, 21},
		{[][2]int{{0, 0}}, 0},
	}

	for _, test := range testCases {
		got := Number(test.busStops)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Number2(test.busStops)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Number3(test.busStops)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
