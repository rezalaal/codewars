package kata7

import "testing"

func TestFizzBuzzCuckooClock(t *testing.T) {
	testCases := []struct {
		time string
		out  string
	}{
		{"13:34", "tick"},
		{"21:00", "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{"11:15", "Fizz Buzz"},
		{"03:03", "Fizz"},
		{"14:30", "Cuckoo"},
		{"08:55", "Buzz"},
		{"00:00", "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{"12:00", "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
	}

	for _, test := range testCases {
		got := FizzBuzzCuckooClock(test.time)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
