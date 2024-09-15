package kata8

import "testing"

func TestQuarterOf(t *testing.T) {
	testCases := []struct {
		month int
		out   int
	}{
		{3, 1},
		{8, 3},
		{11, 4},
	}
	for _, test := range testCases {
		got := QuarterOf(test.month)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}

		got = QuarterOf2(test.month)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}

		got = QuarterOf3(test.month)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}

		got = QuarterOf4(test.month)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}

		got = QuarterOf5(test.month)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}
	}
}
