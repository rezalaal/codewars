package kata7

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		str    string
		ending string
		out    bool
	}{
		{"", "", true},
		{" ", "", true},
		{"abc", "c", true},
		{"banana", "ana", true},
		{"a", "z", false},
		{"", "t", false},
		{"$a = $b + 1", "+1", false},
		{"    ", "   ", true},
		{"1oo", "100", false},
	}
	for _, test := range testCases {
		got := Solution(test.str, test.ending)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Solution1(test.str, test.ending)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Solution2(test.str, test.ending)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Solution3(test.str, test.ending)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
