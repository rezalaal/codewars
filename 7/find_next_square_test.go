package kata7

import "testing"

func TestFindNextSquare(t *testing.T) {
	testCases := []struct {
		sq  int64
		out int64
	}{
		{int64(121), int64(144)},
		{int64(625), int64(676)},
		{int64(319225), int64(320356)},
		{int64(15241383936), int64(15241630849)},
		{int64(155), int64(-1)},
	}

	for _, test := range testCases {
		got := FindNextSquare(test.sq)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = FindNextSquare1(test.sq)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = FindNextSquare2(test.sq)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = FindNextSquare3(test.sq)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = FindNextSquare4(test.sq)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = FindNextSquare5(test.sq)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
