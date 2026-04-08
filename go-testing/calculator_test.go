package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 3)
	expected := 8
	
	if result != expected {
		t.Errorf("Add(5, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 6},
		{4, 5, 20},
		{0, 10, 0},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", 
				test.a, test.b, result, test.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Divide(10, 2) = %f; want 5.0", result)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
}