package kata8

import "testing"

func TestEachCons(t *testing.T) {
	arr := []int{3, 5, 8, 13}
	testCases := []struct {
		arr []int
		n   int
		out [][]int
	}{
		{arr, 1, [][]int{{3}, {5}, {8}, {13}}},
		{arr, 2, [][]int{{3, 5}, {5, 8}, {8, 13}}},
		{arr, 3, [][]int{{3, 5, 8}, {5, 8, 13}}},
		{[]int{}, 3, [][]int{}},
	}

	for _, test := range testCases {
        got := EachCons(test.arr, test.n)
        expected := test.out

        if !equalSlices(got, expected) {
            t.Errorf("Expected %v but got %v", expected, got)
        }
    }
	
	
}

func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalInnerSlices(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalInnerSlices(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
