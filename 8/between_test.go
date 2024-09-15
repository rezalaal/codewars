package kata8

import "testing"

func TestBetween(t *testing.T) {
	testCases := []struct {
		a      int
		b      int
		result []int
	}{
		{1, 4, []int{1, 2, 3, 4}},
		{-2, 2, []int{-2, -1, 0, 1, 2}},
	}

	for _, test := range testCases {
		got := Between(test.a, test.b)
		want := test.result

		for i, n :=range got {
			if want[i] != n {
				t.Errorf("expected %v but got %v", want[i], n)
			}
		}
	}
}
