package kata8

import "testing"

func TestSummation(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{1, 1},
		{8, 36},
		{22, 253},
		{100, 5050},
		{213, 22791},
	}

	for _, test :=range testCases {
		got := Summation(test.n)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = Summation1(test.n)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
