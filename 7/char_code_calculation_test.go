package kata7

import "testing"

func TestCalc(t *testing.T) {
	testCases := []struct {
		s   string
		out int
	}{
		{"abcdef", 6},
		{"ifkhchlhfd", 6},
		{"aaaaaddddr", 30},
		{"jfmgklf8hglbe", 6},
		{"jaam", 12},
	}

	for _, test := range testCases {
		got := Calc(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
