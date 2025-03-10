// trimmean.go
// So many different downloads I had to make. Please email if it doesnt work because I have other versions of this.
// Create package to trim the mean correctly. 
package trimmean

import (
	"fmt"
	"sort"
)

// TrimmedMean computes the trimmed mean
// If only one degree of trimming is provided, it's symmetric.
// If two arguments are provided, the trimming is asymmetric.
// had to be useful for both symmetric and asymmetric data
func TrimmedMean(numbers []float64, lowTrim, highTrim float64) (float64, error) {
	if lowTrim < 0 || highTrim < 0 || lowTrim+highTrim >= 1 {
		return 0, fmt.Errorf("invalid trimming proportions")
	}

	// Sort the numbers in order increasingly
	sort.Float64s(numbers)

	// Calculate the number to trim by using the length of the set 
	n := len(numbers)

	var lowIndex, highIndex int

	// If only one degree of trimming is provided, apply symmetric trimming
	if highTrim == 0 {
		highTrim = lowTrim // Make trimming symmetric if only one argument is provided
	}

	lowIndex = int(float64(n) * lowTrim)
	highIndex = int(float64(n) * highTrim)

	// trimming the low and high ends of the arrray
	trimmedNumbers := numbers[lowIndex : n-highIndex]

	// Compute the mean of the numbers left behind from trimming the array 
	var sum float64
	for _, num := range trimmedNumbers {
		sum += num
	}
	mean := sum / float64(len(trimmedNumbers))

	return mean, nil
}
