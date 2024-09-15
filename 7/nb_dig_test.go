package kata7

import "testing"

func TestNbDig(t *testing.T) {
	testCases := []struct{
		n int
		d int
		out int
	}{
		{10,1, 4},
		{550, 5, 213},
		{5750, 0, 4700},
	}

	for _, test := range testCases {
		got := NbDig(test.n, test.d)
		want := test.out

		assertCorrectMessage(t, want , got)

		got = NbDig2(test.n, test.d)
		want = test.out

		assertCorrectMessage(t, want , got)
	}
}