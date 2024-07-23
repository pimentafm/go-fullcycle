package tax

import "testing"

// func TestCalculateTax(t *testing.T) {
// 	amount := 500.0
// 	expected := 5.0

// 	result := CalculateTax(amount)

// 	if result != expected {
// 		t.Errorf("expected %v but got %v", expected, result)
// 	}
// }

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{0, 0},
		{500, 5},
		{1000, 10},
		{1500, 10},
	}

	for _, test := range tests {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("expected %v but got %v", test.expected, result)
		}
	}
}
