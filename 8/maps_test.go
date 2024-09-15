package kata8

import "testing"

func TestMaps(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{2, 4, 6}},
		{input: []int{4, 1, 1, 1, 4}, expected: []int{8, 2, 2, 2, 8}},
		{input: []int{2, 2, 2, 2, 2, 2}, expected: []int{4, 4, 4, 4, 4, 4}},
	}

	for _, test := range tests {
		got := Maps(test.input)
		for j, v := range got {
			if v != test.expected[j] {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, got)
			}
		}
		got = Maps2(test.input)
		for j, v := range got {
			if v != test.expected[j] {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, got)
			}
		}
	}
} 