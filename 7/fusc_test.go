package kata7

import "testing"

func TestFusc(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{0, 0},
		{1, 1},		
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 3},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 4},
		{10, 3},
		{11, 5},
		{12, 2},
		{13, 5},
		{14, 3},
		{15, 4},
		{16, 1},
		{17, 5},
		{18, 4},
		{19, 7},
		{20, 3},
		{85, 21},
	}

	for _, test := range testCases {
		got := Fusc(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
