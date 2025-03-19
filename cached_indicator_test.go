package techan

import (
	"testing"

	"github.com/sdcoffey/big"
)

// Mock implementation of cachedIndicator for testing
type mockCachedIndicator struct {
	data resultCache
}

func (m *mockCachedIndicator) cache() resultCache {
	return m.data
}

func (m *mockCachedIndicator) setCache(cache resultCache) {
	m.data = cache
}

// Implement Indicator interface by adding a dummy Calculate method
func (m *mockCachedIndicator) Calculate(index int) big.Decimal {
	if index < 0 || index >= len(m.data) || m.data[index] == nil {
		return big.ZERO
	}
	return *m.data[index]
}

// Test cacheResult function
func TestCacheResult(t *testing.T) {
	mockIndicator := &mockCachedIndicator{data: make(resultCache, 3)}
	val := big.NewDecimal(10)

	cacheResult(mockIndicator, 1, val)

	if mockIndicator.cache()[1] == nil || *mockIndicator.cache()[1] != val {
		t.Errorf("Expected cache index 1 to be %v, got %v", val, mockIndicator.cache()[1])
	}
}

// Test expandResultCache function
func TestExpandResultCache(t *testing.T) {
	mockIndicator := &mockCachedIndicator{data: make(resultCache, 2)}

	expandResultCache(mockIndicator, 5)

	if len(mockIndicator.cache()) != 5 {
		t.Errorf("Expected cache size 5, got %d", len(mockIndicator.cache()))
	}
}
