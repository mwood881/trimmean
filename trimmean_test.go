// trimmean_test.go

package trimmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	// Test symmetric trimming
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mean, err := TrimmedMean(numbers, 0.1, 0.1)
	if err != nil || mean != 5 {
		t.Errorf("expected mean 5, got %v", mean)
	}

	// Test asymmetric trimming
	mean, err = TrimmedMean(numbers, 0.1, 0.2)
	if err != nil || mean != 5.5 {
		t.Errorf("expected mean 5.5, got %v", mean)
	}

	// Test invalid trimming proportions (sum exceeds 1)
	mean, err = TrimmedMean(numbers, 0.6, 0.6)
	if err == nil {
		t.Errorf("expected error for invalid trimming proportions, got mean %v", mean)
	}
}
