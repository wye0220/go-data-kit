package slice

import (
	"testing"
)

// TestSumInt tests the Sum function for slices of int.
func TestSumInt(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := 15
	if result := Sum(data); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestSumFloat64 tests the Sum function for slices of float64.
func TestSumFloat64(t *testing.T) {
	data := []float64{1.1, 2.2, 3.3}
	expected := 6.6
	if result := Sum(data); result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

// TestSumComplex64 tests the Sum function for slices of complex64.
func TestSumComplex64(t *testing.T) {
	data := []complex64{1 + 2i, 3 + 4i}
	expected := complex64(4 + 6i)
	if result := Sum(data); result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestSumEmptySlice tests the Sum function for an empty slice.
func TestSumEmptySlice(t *testing.T) {
	var data []int
	expected := 0
	if result := Sum(data); result != expected {
		t.Errorf("Expected %d for an empty slice, got %d", expected, result)
	}
}

// You can add more tests for other types like uint, int8, etc. as needed.
