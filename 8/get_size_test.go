package kata8

import "testing"

func TestGetSize(t *testing.T) {
	testCases := []struct {
		w   int
		h   int
		d   int
		out [2]int
	}{
		{4, 2, 6, [2]int{88, 48}},
		{1, 1, 1, [2]int{6, 1}},
		{1, 2, 1, [2]int{10, 2}},
		{1, 2, 2, [2]int{16, 4}},
		{10, 10, 10, [2]int{600, 1000}},
	}

	for _, test := range testCases {
		got := GetSize(test.w, test.h, test.d)
		want := test.out

		for i, item := range got {
			if want[i] != item {
				t.Errorf("Expected %v but got %v", want[i], item)
			}
		}
	}
}
