package kata8

import "testing"

func TestMergeArrays(t *testing.T) {
	testCases := []struct {
		arr1 []int
		arr2 []int
		out  []int
	}{
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{1, 3, 5, 7, 9}, []int{10, 8, 6, 4, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 3, 5, 7, 9, 11, 12}, []int{1, 2, 3, 4, 5, 10, 12}, []int{1, 2, 3, 4, 5, 7, 9, 10, 11, 12}},
	}
	for _, test := range testCases {
		got := MergeArrays(test.arr1, test.arr2)
		want := test.out
		for i, item := range got {
			if want[i] != item {
				t.Errorf("Expected %v but got %v", want[i], item)
			}
		}

		got = MergeArrays1(test.arr1, test.arr2)
		want = test.out
		for i, item := range got {
			if want[i] != item {
				t.Errorf("Expected %v but got %v", want[i], item)
			}
		}

		got = MergeArrays2(test.arr1, test.arr2)
		want = test.out
		for i, item := range got {
			if want[i] != item {
				t.Errorf("Expected %v but got %v", want[i], item)
			}
		}
	}
}
