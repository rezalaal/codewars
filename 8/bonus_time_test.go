package kata8

import "testing"

func TestBonusTime(t *testing.T) {
	testCases := []struct {
		salary int
		bonus  bool
		out    string
	}{
		{100, false, "£100"},
		{9, false, "£9"},
		{55000, false, "£55000"},
		{100, true, "£1000"},
		{14000, true, "£140000"},
	}

	for _, test := range testCases {
		got := BonusTime(test.salary, test.bonus)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
