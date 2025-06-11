package main

import "testing"

func TestSum(t *testing.T) {
	result, _ := sum(3, 2)
	if result != 5 {
		t.Errorf("Expected 5, got %f", result)
	}
}

func TestSubtract(t *testing.T) {
	result, _ := subtract(10, 4)
	if result != 6 {
		t.Errorf("Expected 6, got %f", result)
	}
}

func TestMultiply(t *testing.T) {
	result, _ := multiply(3, 3)
	if result != 9 {
		t.Errorf("Expected 9, got %f", result)
	}
}

func TestDivide(t *testing.T) {
	result, _ := divide(8, 2)
	if result != 4 {
		t.Errorf("Expected 4, got %f", result)
	}

	_, err := divide(5, 0)
	if err == nil {
		t.Error("Expected an error for division by zero")
	}
}
