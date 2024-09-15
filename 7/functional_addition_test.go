package kata7

import "testing"

func TestAdd(t *testing.T) {
	// Test adding 1
    addOne := Add(1)
    result := addOne(3)
    expected := 4
    if result != expected {
        t.Errorf("addOne(3) = %d; expected %d", result, expected)
    }

    // Test adding 5
    addFive := Add(5)
    result = addFive(10)
    expected = 15
    if result != expected {
        t.Errorf("addFive(10) = %d; expected %d", result, expected)
    }

    // Test adding 0
    addZero := Add(0)
    result = addZero(7)
    expected = 7
    if result != expected {
        t.Errorf("addZero(7) = %d; expected %d", result, expected)
    }

    // Test adding negative numbers
    addNegative := Add(-3)
    result = addNegative(5)
    expected = 2
    if result != expected {
        t.Errorf("addNegative(5) = %d; expected %d", result, expected)
    }

    // Test adding large numbers
    addLarge := Add(1000000)
    result = addLarge(5000000)
    expected = 6000000
    if result != expected {
        t.Errorf("addLarge(5000000) = %d; expected %d", result, expected)
    }
}