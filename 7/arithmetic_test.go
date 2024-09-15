package kata7

import "testing"

func TestArithmetic(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		operator string
		res      int
	}{
		{8, 2, "add", 10},
		{8, 2, "subtract", 6},
		{8, 2, "multiply", 16},
		{8, 2, "divide", 4},
	}

	for _, test := range testCases {
		got := Arithmetic(test.a, test.b, test.operator)
		want := test.res
		assertCorrectMessage(t, want, got)

		got = Arithmetic2(test.a, test.b, test.operator)
		want = test.res
		assertCorrectMessage(t, want, got)

		got = Arithmetic3(test.a, test.b, test.operator)
		want = test.res
		assertCorrectMessage(t, want, got)
	}
}
