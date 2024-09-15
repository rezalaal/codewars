package kata8

import "testing"


func TestCalculateYears(t *testing.T) {
	testCases := []struct {
		t int
		output [3]int
	}{
		{t:1, output:[3]int{1, 15, 15}},
		{t:2, output:[3]int{2, 24, 24}},
		{t:10, output:[3]int{10, 56, 64}},
	}

	for _, test :=range testCases {
		got := CalculateYears(test.t)
		expected := test.output

		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}

		got = CalculateYears1(test.t)
		expected = test.output

		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}
	
}