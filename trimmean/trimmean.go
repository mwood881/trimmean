// trimmean.go

package trimmean

import (
	"fmt"
	"sort"
)

// TrimmedMean computes the trimmed mean from a slice of numbers.
// If only one degree of trimming is provided, it's symmetric.
// If two arguments are provided, the trimming is asymmetric.
func TrimmedMean(numbers []float64, lowTrim, highTrim float64) (float64, error) {
	if lowTrim < 0 || highTrim < 0 || lowTrim+highTrim >= 1 {
		return 0, fmt.Errorf("invalid trimming proportions")
	}

	// Sort the numbers in ascending order
	sort.Float64s(numbers)

	// Calculate the number of elements to trim
	n := len(numbers)

	var lowIndex, highIndex int

	// If only one degree of trimming is provided, apply symmetric trimming
	if highTrim == 0 {
		highTrim = lowTrim // Make trimming symmetric if only one argument is provided
	}

	lowIndex = int(float64(n) * lowTrim)
	highIndex = int(float64(n) * highTrim)

	// Slice the array, trimming the low and high ends
	trimmedNumbers := numbers[lowIndex : n-highIndex]

	// Compute the mean of the remaining numbers
	var sum float64
	for _, num := range trimmedNumbers {
		sum += num
	}
	mean := sum / float64(len(trimmedNumbers))

	return mean, nil
}
