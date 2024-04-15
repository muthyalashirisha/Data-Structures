package problems_test

import (
	"DS/problems"
	"fmt"
	"testing"
	"time"
)

func TestDistinctIslands(t *testing.T) {
	t.Parallel()

	// Set a deadline for the test to complete
	deadline := time.Now().Add(1 * time.Millisecond) // Example: 10 seconds

	tests := []struct {
		grid        [][]int
		expectedNum int
	}{
		// Test cases
		{
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			expectedNum: 1, // Expected number of distinct islands
		},
		// Add more test cases as needed
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			if time.Now().After(deadline) {
				t.Fatalf("Test took too long. Deadline exceeded.")
			}
			actualNum := problems.DistinctIslands(test.grid)
			if actualNum != test.expectedNum {
				t.Errorf("Test case %d: Expected %d, but got %d", i+1, test.expectedNum, actualNum)
			}
		})

	}
}
