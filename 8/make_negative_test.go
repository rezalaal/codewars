package kata8

import (
	"testing"
)

func TestMakeNegative(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{42, -42},
		{-42, -42},
		{0, 0},
	}

	for _, test := range testCases {
		got := MakeNegative(test.n)		
		want := test.out
		
		if got != want {
			t.Errorf("Expect %v but got %v", want, got)
		}

		got = MakeNegative1(test.n)		
		want = test.out
		
		if got != want {
			t.Errorf("Expect %v but got %v", want, got)
		}
	}
}
