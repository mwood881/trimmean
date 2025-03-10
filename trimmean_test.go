// trimmean_test.go
// use as bencmark testing
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
// Had to test asymmetric as well as it has to work for both assymetric and symmetric
	// Test asymmetric trimming
	mean, err = TrimmedMean(numbers, 0.1, 0.2)
	if err != nil || mean != 5.5 {
		t.Errorf("expected mean 5.5, got %v", mean)
	}

	// Test is more than 1 so it is out of the trimming range
	mean, err = TrimmedMean(numbers, 0.6, 0.6)
	if err == nil {
		t.Errorf("expected error for invalid trimming proportions, got mean %v", mean)
	}
}
// ALL tests passed yayyy!
