package kata8

import "testing"

func TestSeatsInTheater(t *testing.T) {
	testCases := []struct {
		nCols int
		nRows int
		col   int
		row   int
		out   int
	}{
		{16, 11, 5, 3, 96},
		{1, 1, 1, 1, 0},
		{13, 6, 8, 3, 18},
		{60, 100, 60, 1, 99},
		{1000, 1000, 1000, 1000, 0},
	}

	for _, test := range testCases {
		got := SeatsInTheater(test.nCols, test.nRows, test.col, test.row)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
