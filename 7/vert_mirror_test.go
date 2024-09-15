package kata7

import (
	"testing"
)

func TestVertMirror(t *testing.T) {
	input := "abcd\nefgh\nijkl\nmnop"
	expected := "dcba\nhgfe\nlkji\nponm"
	result := VertMirror(input)
	if result != expected {
		t.Errorf("VertMirror(%q) = %q; want %q", input, result, expected)
	}
}

func TestHorMirror(t *testing.T) {
	input := "abcd\nefgh\nijkl\nmnop"
	expected := "mnop\nijkl\nefgh\nabcd"
	result := HorMirror(input)
	if result != expected {
		t.Errorf("HorMirror(%q) = %q; want %q", input, result, expected)
	}
}

func TestOper(t *testing.T) {
	input := "abcd\nefgh\nijkl\nmnop"

	// Test with VertMirror
	expectedVert := "dcba\nhgfe\nlkji\nponm"
	resultVert := Oper(VertMirror, input)
	if resultVert != expectedVert {
		t.Errorf("Oper(VertMirror, %q) = %q; want %q", input, resultVert, expectedVert)
	}

	// Test with HorMirror
	expectedHor := "mnop\nijkl\nefgh\nabcd"
	resultHor := Oper(HorMirror, input)
	if resultHor != expectedHor {
		t.Errorf("Oper(HorMirror, %q) = %q; want %q", input, resultHor, expectedHor)
	}

	input = "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu"

	// Test with VertMirror
	expectedVert = "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw"
	resultVert = Oper(VertMirror, input)
	if resultVert != expectedVert {
		t.Errorf("Oper(VertMirror, %q) = %q; want %q", input, resultVert, expectedVert)
	}

	input = "lVHt\nJVhv\nCSbg\nyeCt"
	// Test with HorMirror
	expectedHor = "yeCt\nCSbg\nJVhv\nlVHt"
	resultHor = Oper(HorMirror, input)
	if resultHor != expectedHor {
		t.Errorf("Oper(HorMirror, %q) = %q; want %q", input, resultHor, expectedHor)
	}
}
