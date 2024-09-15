package kata8

import "testing"

func TestCountPositivesSumNegatives(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	got := CountPositivesSumNegatives(arr)
	expected := []int{10, -65}

	for i, n := range got {
		if expected[i] != n {
			t.Errorf("Expected %v but got %v", expected[i], n)
		}
	}
}
