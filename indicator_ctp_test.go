package techan

import (
	"fmt"
	"math"
	"testing"

	"github.com/sdcoffey/big"
)

type mockIndicator struct {
	values []big.Decimal
}

func (m mockIndicator) Calculate(index int) big.Decimal {
	if index < 0 || index >= len(m.values) {
		return big.ZERO
	}
	return m.values[index]
}

func mockIndicatorFromFloats(values []float64) Indicator {
	bigValues := make([]big.Decimal, len(values))
	for i, v := range values {
		bigValues[i] = big.NewDecimal(v)
	}
	return mockIndicator{values: bigValues}
}

func TestCtpIndicator(t *testing.T) {
	mockHighs := mockIndicatorFromFloats([]float64{10, 12, 14, 16, 18})
	mockLows := mockIndicatorFromFloats([]float64{5, 6, 7, 8, 9})
	mockCloses := mockIndicatorFromFloats([]float64{8, 10, 12, 14, 16})

	ctp := NewCtpIndicator(mockHighs, mockLows, mockCloses)

	expectedValues := []float64{
		(10 + 5 + 8) / 3.0,  // 7.67
		(12 + 6 + 10) / 3.0, // 9.33
		(14 + 7 + 12) / 3.0, // 11.00
		(16 + 8 + 14) / 3.0, // 12.67
		(18 + 9 + 16) / 3.0, // 14.33
	}

	for i, expected := range expectedValues {
		actual := ctp.Calculate(i).Float()
		fmt.Printf("Index %d: Expected %.4f, Got %.4f\n", i, expected, actual)
		if !almostEqual(actual, expected) {
			t.Errorf("CTP at index %d: expected %.4f, got %.4f", i, expected, actual)
		}
	}
}

func almostEqual(a, b float64) bool {
	const epsilon = 0.00001
	return math.Abs(a-b) < epsilon
}
