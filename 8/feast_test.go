package kata8

import "testing"

func TestFeast(t *testing.T) {
	testCases := []struct {
		beast string
		dish  string
		out   bool
	}{
		{"great blue heron", "garlic naan", true}, 
		{"chickadee", "chocolate cake", true},     
		{"brown bear", "bear claw", false},        
	}

	for _, test := range testCases {
		got := Feast(test.beast, test.dish)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
