package kata7

import "testing"

func TestPrevMultOfThree(t *testing.T) {
	testCases := []struct {
		n   int
		out interface{}
	}{
		{1, nil},
		{25, nil},
		{36, 36},
		{1244, 12},
		{952406, 9},
	}

	for _, test := range testCases {
		got := PrevMultOfThree(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
