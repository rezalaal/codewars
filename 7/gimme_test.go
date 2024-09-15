package kata7

import "testing"

func TestGimme(t *testing.T) {
	testCases := []struct {
		array [3]int
		out   int
	}{
		{[3]int{2, 3, 1}, 0},
		{[3]int{5, 10, 14}, 1},
		{[3]int{1, 3, 4}, 1},
		{[3]int{15, 10, 14}, 2},
		{[3]int{-4, -23, 4}, 0},
		{[3]int{-15, -10, 14}, 1},
	}

	for _, test := range testCases {
		got := Gimme(test.array)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Gimme1(test.array)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Gimme2(test.array)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Gimme3(test.array)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Gimme4(test.array)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
