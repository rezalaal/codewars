package kata8

import "testing"

func TestSmallestIntegerFinder(t *testing.T) {
	testCases := []struct{
		numbers []int
		out int
	}{
		{[]int{34, 15, 88, 2},     2},
		{[]int{34, -345, -1, 100}, -345},
	}

	for _, test :=range testCases {
		got := SmallestIntegerFinder(test.numbers)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = SmallestIntegerFinder1(test.numbers)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}