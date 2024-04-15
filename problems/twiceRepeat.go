package problems

import (
	"fmt"
	"math"
	"sort"
)

func TwoRepeating(arr []int) {
	ans := []int{}
	for _, v := range arr {
		fmt.Println(arr)
		if arr[int(math.Abs(float64(v)))] < 0 {
			fmt.Println(int(math.Abs(float64(v))))
			ans = append(ans, int(math.Abs(float64(v))))
		}
		arr[int(math.Abs(float64(v)))] = -1 * arr[int(math.Abs(float64(v)))]
	}
	fmt.Println(arr)
	fmt.Println(ans)
}

// func TwoRepeating1(arr []int) {
// 	var res []int
// 	for i, a := 0, 0; i < len(arr); i++ {
// 		if arr[abs(arr[i])] < 0 {
// 			res = append(res, abs(arr[i]))
// 			a++
// 		}
// 		arr[abs(arr[i])] = -1 * arr[abs(arr[i])]
// 	}
// 	fmt.Println(res)
// }

//	func abs(x int) int {
//		if x < 0 {
//			return -x
//		}
//		return x
//	}
func FindDuplicates(nums []int) []int {
	tortoise := nums[0]
	hare := nums[0]

	// Phase 1: Find the intersection point
	for {
		fmt.Println(nums[tortoise], nums[nums[hare]])
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	// Phase 2: Find the entry point of the cycle
	hare = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	// At this point, 'hare' and 'tortoise' point to the start of the cycle

	// Phase 3: Find the second duplicate
	hare = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	fmt.Println(tortoise, hare)
	return []int{tortoise, hare}
}

// func FindLargestZeroSumSubmatrix(mat [][]int) [][]int {
// 	rows := len(mat)
// 	if rows == 0 {
// 		return [][]int{}
// 	}
// 	cols := len(mat[0])

// 	// Compute the prefix sum matrix
// 	prefixSum := make([][]int, rows)
// 	for i := range prefixSum {
// 		prefixSum[i] = make([]int, cols)
// 		copy(prefixSum[i], mat[i])
// 		for j := 1; j < cols; j++ {
// 			prefixSum[i][j] += prefixSum[i][j-1]
// 		}
// 	}
// 	fmt.Println(prefixSum)
// 	maxArea := 0
// 	var result [][]int

// 	// Iterate over all possible pairs of columns
// 	for start := 0; start < cols; start++ {
// 		for end := start; end < cols; end++ {
// 			// Calculate column sums using prefix sum
// 			colSum := make([]int, rows)
// 			for i := 0; i < rows; i++ {
// 				if start > 0 {
// 					colSum[i] = prefixSum[i][end] - prefixSum[i][start-1]
// 				} else {
// 					colSum[i] = prefixSum[i][end]
// 				}
// 			}

// 			// Apply Kadane's algorithm for 1D array
// 			colIndexMap := make(map[int]int)
// 			colIndexMap[0] = -1
// 			currentSum := 0
// 			startRow := -1
// 			endRow := -1

// 			for i := 0; i < rows; i++ {
// 				currentSum += colSum[i]
// 				if index, ok := colIndexMap[currentSum]; ok {
// 					if i-index > endRow-startRow {
// 						startRow = index
// 						endRow = i
// 					}
// 				} else {
// 					colIndexMap[currentSum] = i
// 				}
// 			}

// 			// Update maximum area
// 			if startRow != -1 && endRow != -1 {
// 				currentArea := (end - start + 1) * (endRow - startRow + 1)
// 				if currentArea > maxArea {
// 					maxArea = currentArea
// 					// Update the result matrix
// 					result = make([][]int, endRow-startRow+1)
// 					for i := range result {
// 						result[i] = make([]int, end-start+1)
// 						copy(result[i], mat[startRow+i][start:end+1])
// 					}
// 				}
// 			}
// 		}
// 	}

//		return result
//	}
func kadane(v []int) (int, int) {
	sum := 0
	maxi := -1e9
	l, r := 0, -1

	s := 0
	for i := 0; i < len(v); i++ {
		sum += v[i]
		if float64(sum) > maxi {
			maxi = float64(sum)
			l = s
			r = i
		}

		if sum < 0 {
			sum = 0
			s = i + 1
		}
	}

	return l, r
}

func FindLargestZeroSumSubmatrix(a [][]int) [][]int {
	m := len(a)
	n := len(a[0])
	fmt.Println(a[0])
	matrix := a
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	rows, cols := len(matrix), len(matrix[0])

	// Calculate the prefix sum matrix
	prefixSum := make([][]int, rows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			fmt.Println(matrix[i-1][j-1], prefixSum[i-1][j], prefixSum[i][j-1], prefixSum[i-1][j-1], i, j)
			prefixSum[i][j] = matrix[i-1][j-1] + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
			fmt.Println(prefixSum)
		}
	}

	left, right, up, down := 0, 0, 0, 0

	for i := 0; i < n; i++ {
		arr := make([]int, m)

		for j := i; j < n; j++ {
			for k := 0; k < m; k++ {
				arr[k] += a[k][j]
			}
			sumMap := make(map[int]int)
			sumMap[0] = -1

			l, r := 0, 0
			sum := 0
			fmt.Println(arr, "arr")
			for k := 0; k < m; k++ {
				sum += arr[k]

				if val, ok := sumMap[sum]; ok {
					fmt.Println(k, val, r, l, "....")
					if (k - val) > (r - l) {
						l = val + 1
						r = k + 1
					}
					fmt.Println(l, r, "l", "r")
				} else {
					sumMap[sum] = k
				}
				fmt.Println(l, r, sum, sumMap)
			}
			fmt.Println((j-i+1)*(r-l), (right-left)*(down-up), i, j, r, l, right, left, up, down)
			if (j-i+1)*(r-l) > (right-left)*(down-up) {
				up = l
				down = r
				left = i
				right = j + 1
			}
		}
	}

	result := make([][]int, 0)

	for i := up; i < down; i++ {
		arr := make([]int, 0)
		for j := left; j < right; j++ {
			arr = append(arr, a[i][j])
		}
		result = append(result, arr)
	}

	return result
}

func AntiDiagonal(matrix [][]int) {
	result := []int{}
	// k := 1
	for i := 0; i < len(matrix); i++ {
		for j := i; j >= 0; j-- {
			result = append(result, matrix[i-j][j])
		}
	}
	for i := 1; i < len(matrix); i++ {
		k := i
		for j := len(matrix) - 1; j >= i; j-- {
			result = append(result, matrix[k][j])
			k++
		}
	}
	fmt.Println(result)
}
func Match(pattern, wild string) bool {
	// l := 0
	// r := 0
	// starIndex := -1
	// matchIndex := -1
	// for r < len(pattern) {
	// 	if l < len(wild) && (wild[l] == pattern[r] || wild[l] == '?') {
	// 		l++
	// 		r++
	// 	} else {
	// 		if l < len(wild) && wild[l] == '*' {
	// 			starIndex = l
	// 			matchIndex = r
	// 			l++
	// 		} else if starIndex != -1 {
	// 			l = starIndex + 1
	// 			r = matchIndex + 1
	// 			matchIndex = r
	// 		} else {
	// 			return false
	// 		}
	// 	}
	// }
	// for l < len(wild) && wild[l] == '*' {
	// 	l++
	// }
	// if l == len(wild) {
	// 	return true
	// }
	// return false
	m := len(wild)
	n := len(pattern)

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	// Handle patterns with * at the beginning
	for j := 1; j <= n && pattern[j-1] == '*'; j++ {
		dp[0][j] = dp[0][j-1]
	}
	fmt.Println(dp)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if wild[i-1] == pattern[j-1] || wild[i-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if wild[i-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)

	return dp[m][n]
}

type Meeting struct {
	Start int
	End   int
	Index int
}

func MaxMeetings(start, end []int, n int) int {
	meetings := []Meeting{}

	for i := 0; i < n; i++ {
		meeting := Meeting{
			Start: start[i],
			End:   end[i],
			Index: i,
		}
		meetings = append(meetings, meeting)
	}

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].End < meetings[j].End {
			return true
		} else if meetings[i].End == meetings[j].End {
			if meetings[i].Index < meetings[j].Index {
				return true
			}
		}
		return false
	})

	cnt := 1
	endTime := meetings[0].End
	for i := 1; i < n; i++ {
		if meetings[i].Start > endTime {
			endTime = meetings[i].End
			cnt++
		}
	}
	return cnt
}
