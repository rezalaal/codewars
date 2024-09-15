package kata8

import "testing"

func TestSetAlarm(t *testing.T) {
	testCases := []struct {
		employed bool
		vacation bool
		out      bool
	}{
		{true, true, false},
		{false, true, false},
		{false, false, false},
		{true, false, true},
	}

	for _, test := range testCases {
		got := SetAlarm(test.employed, test.vacation)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
