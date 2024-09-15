package kata7

import "testing"

func TestGetSum(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		out int
	}{
		{0, 1, 1},
		{1, 2, 3},
		{5, -1, 14},
		{505, 4, 127759},
		{321, 123, 44178},
		{0, -1, -1},
		{-50, 0, -1275},
		{-1, -5, -15},
		{-5, -5, -5},
		{-505, 4, -127755},
		{-321, 123, -44055},
		{0, 0, 0},
		{-5, -1, -15},
		{5, 1, 15},
		{-17, -17, -17},
		{17, 17, 17},
	}

	for _, test := range testCases {
		got := GetSum(test.a, test.b)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = GetSum1(test.a, test.b)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = GetSum2(test.a, test.b)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
