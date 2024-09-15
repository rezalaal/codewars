package kata8

import "testing"

func TestGetVolumeOfCuboid(t *testing.T) {
	testCases := []struct {
		length float64
		width  float64
		height float64
		out    float64
	}{
		{1.0, 2.0, 2.0, 4.0},
		{6.3, 2.0, 5.0, 63.0},
	}

	for _, test := range testCases {
		got := GetVolumeOfCuboid(test.length, test.width, test.height)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
