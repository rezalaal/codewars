package kata8

import "testing"

func TestGreeting(t *testing.T) {
	testCases := []struct{
		name string
		out string
	}{
		{"Ryan","Hello, Ryan how are you doing today?"},
	}
	
	for _, test := range testCases {
		got := Greeting(test.name)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}