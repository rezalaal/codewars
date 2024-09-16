package kata7

import "testing"

func TestCats(t *testing.T) {
	testCases := []struct {
		start  int
		finish int
		out    int
	}{
		{1, 5, 2},
		{1, 1, 0},
		{2, 5, 1},
		{2, 4, 2},
	}

	for _, test := range testCases {
		got := Cats(test.start, test.finish)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
