package problems

import "math"

func minFallingPathSum(matrix [][]int) int {
	minSum := math.MaxInt
	n := len(matrix)
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix); j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
			} else {
				dp[i][j] = 1e9 + 7
			}
		}
	}

	for row := 1; row < len(matrix[0]); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			currMin := math.MaxInt
			// left
			if col > 0 {
				min := dp[row+1][col-1]
				if min < currMin {
					currMin = min
				}
			}
			// down
			min := dp[row+1][col]
			if min < currMin {
				currMin = min
			}
			// right
			if col < n-1 {
				min := dp[row+1][col+1]
				if min < currMin {
					currMin = min
				}
			}
			dp[row][col] = matrix[row][col] + currMin
		}
	}
	for i := 0; i < n; i++ {
		if dp[n-1][i] < minSum {
			minSum = dp[n-1][i]
		}
	}
	return minSum
}
