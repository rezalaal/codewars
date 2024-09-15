package kata8

import "testing"

func TestCountSheeps(t *testing.T) {
	arr1 := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	got := CountSheeps(arr1)
	expected := 17

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
