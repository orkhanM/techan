package techan

import (
	"fmt"
	"testing"
)

// Unit test for VWAP Indicator
func TestVwapIndicator(t *testing.T) {
	mockVolumes := mockIndicatorFromFloats([]float64{100, 200, 300, 400, 500})
	mockCtps := mockIndicatorFromFloats([]float64{10, 20, 30, 40, 50})

	vwap := NewVwapIndicator(mockVolumes, mockCtps)

	expectedValues := []float64{
		10,                                 // (100*10) / 100 = 10
		(100*10 + 200*20) / 300.0,          // 16.67
		(100*10 + 200*20 + 300*30) / 600.0, // 23.33
		(100*10 + 200*20 + 300*30 + 400*40) / 1000.0,          // 30.00
		(100*10 + 200*20 + 300*30 + 400*40 + 500*50) / 1500.0, // 36.67
	}

	for i, expected := range expectedValues {
		actual := vwap.Calculate(i).Float()

		// Debugging print statement
		fmt.Printf("Index %d: Expected VWAP = %.4f, Calculated VWAP = %.4f\n", i, expected, actual)

		if !almostEqual(actual, expected) { // Use the existing almostEqual function
			t.Errorf("VWAP at index %d: expected %.4f, got %.4f", i, expected, actual)
		}
	}
}
