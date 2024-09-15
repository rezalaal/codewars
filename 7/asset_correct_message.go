package kata7

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t *testing.T, want any, got any) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expect %v but got %v", want, got)
	}
}
