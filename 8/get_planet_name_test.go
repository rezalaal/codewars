package kata8

import "testing"

func TestGetPlanetName(t *testing.T) {
	testCases := []struct {
		ID  int
		out string
	}{
		{2, "Venus"},
		{5, "Jupiter"},
		{3, "Earth"},
		{4, "Mars"},
		{8, "Neptune"},
		{1, "Mercury"},
	}

	for _, test := range testCases {
		got := GetPlanetName(test.ID)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
