package kata7

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		slice []string
		out   []int
	}{
		{[]string{"abode", "ABc", "xyzD"}, []int{4, 3, 1}},
		{[]string{"abide", "ABc", "xyz"}, []int{4, 3, 0}},
		{[]string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"}, []int{6, 5, 7}},
		{[]string{"encode", "abc", "xyzD", "ABmD"}, []int{1, 3, 1, 3}},
	}

	for _, test := range testCases {
		got := Solve(test.slice)
		want := test.out

		assertCorrectMessage(t, want, got)
	}
}