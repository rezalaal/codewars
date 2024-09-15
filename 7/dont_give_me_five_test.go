package kata7

import "testing"

func TestDontGiveMeFive(t *testing.T) {
	testCases := []struct {
		start int
		end   int
		out   int
	}{
		{1, 9, 8},
		{4, 17, 12},
	}

	for _, test := range testCases {
		got := DontGiveMeFive(test.start, test.end)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = DontGiveMeFive1(test.start, test.end)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = DontGiveMeFive2(test.start, test.end)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = DontGiveMeFive3(test.start, test.end)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
