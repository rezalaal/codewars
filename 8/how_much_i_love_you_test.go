package kata8

import "testing"

func TestHowMuchILoveYou(t *testing.T) {
	testCases := []struct {
		i   int
		out string
	}{
		{7, "I love you"},
		{3, "a lot"},
		{6, "not at all"},
	}

	for _, test := range testCases {
		got := HowMuchILoveYou(test.i)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = HowMuchILoveYou2(test.i)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = HowMuchILoveYou3(test.i)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
