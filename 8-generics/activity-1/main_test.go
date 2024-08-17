package main

import "testing"

func TestFindMinGeneric(t *testing.T) {
	// Test case for float64
	floatNums := []float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}
	expectedFloatMin := 1.1

	if result := findMinGeneric(floatNums); result != expectedFloatMin {
		t.Errorf("findMinGeneric(floatnums) = %v; want %v", floatNums, expectedFloatMin)
	}

	// Test case for empty slice (int)
	emptyIntNums := []int{}
	expectedEmptyIntMin := 0
	if result := findMinGeneric(emptyIntNums); result != int(expectedEmptyIntMin) {
		t.Errorf("findMinGeneric(floatnums) = %v, want %v", emptyIntNums, expectedEmptyIntMin)
	}

	// Test case for empty slice (float64)
	emptyFloatNums := []float64{}
	expectedEmptyFloatMin := 0.0
	if result := findMinGeneric(emptyFloatNums); result != expectedEmptyFloatMin {
		t.Errorf("findMinGeneric(floatnums) = %v, want %v", emptyFloatNums, expectedEmptyFloatMin)
	}
}
