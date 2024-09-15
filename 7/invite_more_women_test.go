package kata7

import "testing"

func TestInviteMoreWomen(t *testing.T) {
	testCases := []struct {
		L   []int
		out bool
	}{
		{[]int{1, -1, 1}, true},
		{[]int{1, 1, 1}, true},
		{[]int{-1, -1, -1}, false},
		{[]int{1, -1}, false},
	}

	for _, test := range testCases {
		got := InviteMoreWomen(test.L)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
