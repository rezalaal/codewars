package kata7

import "testing"

func TestGetMiddle(t *testing.T) {
	testCases := []struct {
		s   string
		out string
	}{
		{"test", "es"},
		{"testing", "t"},
		{"middle", "dd"},
		{"A", "A"},
	}

	for _, test := range testCases {
		got := GetMiddle(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = GetMiddle1(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = GetMiddle2(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = GetMiddle3(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
