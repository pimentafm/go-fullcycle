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

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0, 30000.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
	})
}
