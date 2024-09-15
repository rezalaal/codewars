package kata8

import "testing"

func TestGrow(t *testing.T) {
	testCases := []struct{
		input []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{4, 1, 1, 1, 4}, 16},
		{[]int{2, 2, 2, 2, 2, 2}, 64},
	}
	
	for _, test :=range testCases {
		got := Grow(test.input)

		if got != test.expected {
			t.Errorf("ToAlternatingCase: expected %v but got %v", test.expected, got)
		}
	}
}