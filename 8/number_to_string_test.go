package kata8

import "testing"

func TestNumberToString(t *testing.T) {
	testCases := []struct {
		n        int
		out string
	}{
		{67,    "67"},
		{79585, "79585"},
		{3,     "3"},
		{-1,    "-1"},
	}

	for _, test := range testCases {
		got := NumberToString(test.n)
		if got != test.out {
			t.Errorf("For input %v, expected %v but got %v", test.n, test.out, got)
		}
		got = NumberToString2(test.n)
		if got != test.out {
			t.Errorf("For input %v, expected %v but got %v", test.n, test.out, got)
		}
		got = NumberToString3(test.n)
		if got != test.out {
			t.Errorf("For input %v, expected %v but got %v", test.n, test.out, got)
		}
		got = NumberToString4(test.n)
		if got != test.out {
			t.Errorf("For input %v, expected %v but got %v", test.n, test.out, got)
		}
		got = NumberToString5(test.n)
		if got != test.out {
			t.Errorf("For input %v, expected %v but got %v", test.n, test.out, got)
		}
	}
}
