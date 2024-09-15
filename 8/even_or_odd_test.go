package kata8

import "testing"

func TestEvenOrOdd(t *testing.T) {
	testCases := []struct {
		number int
		out    string
	}{
		{1, "Odd"},
		{2, "Even"},
		{-1, "Odd"},
		{-2, "Even"},
		{0, "Even"},
		{1, "Odd"},
	}
	for _, test := range testCases {
		got := EvenOrOdd(test.number)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = EvenOrOdd2(test.number)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = EvenOrOdd3(test.number)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = EvenOrOdd4(test.number)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
