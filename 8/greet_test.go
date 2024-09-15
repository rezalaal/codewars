package kata8

import "testing"

func TestGreet(t *testing.T) {
	got := Greet()
	want := "hello world!"

	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
	}
}
