package kata8

import "testing"

func TestInvert(t *testing.T) {
	testCases := []struct {
		arr []int
		out []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{-1, -2, -3, -4, -5}},
		{[]int{1, -2, 3, -4, 5}, []int{-1, 2, -3, 4, -5}},
		{nil, nil},
		{[]int{0}, []int{0}},
	}

	for _, test := range testCases {
		got := Invert(test.arr)
		want := test.out

		for i, n:=range got {
			if want[i] != n {
				t.Errorf("Expected %v but want %v", want[i], n)
			}
		}

		got = Invert1(test.arr)
		want = test.out

		for i, n:=range got {
			if want[i] != n {
				t.Errorf("Expected %v but want %v", want[i], n)
			}
		}
	}
}
