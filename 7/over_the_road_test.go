package kata7

import "testing"

func TestOverTheRoad(t *testing.T) {
	testCases := []struct {
		address int
		n       int
		out     int
	}{
		{1, 3, 6},
		{3, 3, 4},
		{2, 3, 5},
		{3, 5, 8},
		{7, 11, 16},
		{23633656673, 310027696726, 596421736780},
	}

	for _, test := range testCases {
		got := OverTheRoad(test.address, test.n)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = OverTheRoad1(test.address, test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = OverTheRoad2(test.address, test.n)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
