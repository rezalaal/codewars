package kata8

import "testing"

func TestSubtractSum(t *testing.T) {
	testCases := []struct {
		n   int
		out string
	}{
		{10, "pineapple"},
		{325, "apple"},
	}
	for _, test := range testCases {
		got := SubtractSum(test.n)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
