package problems

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func GetMaxGold(n, m int, mine [][]int) int {
	rows := len(mine)
	cols := len(mine[0])

	// Create a 2D slice to store the maximum gold collected at each cell
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	fmt.Println(dp)

	// Initialize the last column of dp with the values from the last column of the mine
	for i := 0; i < rows; i++ {
		dp[i][cols-1] = mine[i][cols-1]
	}
	fmt.Println(dp)
	// Iterate through the columns from right to left
	for j := cols - 2; j >= 0; j-- {
		for i := 0; i < rows; i++ {
			// Consider three possible moves: right, up-right, down-right
			moveRight := dp[i][j+1]
			moveUpRight := 0
			if i > 0 {
				moveUpRight = dp[i-1][j+1]
			}
			moveDownRight := 0
			if i < rows-1 {
				moveDownRight = dp[i+1][j+1]
			}

			// Update the current cell with the maximum value of the three moves plus the gold in the current cell
			dp[i][j] = mine[i][j] + max3(moveRight, moveUpRight, moveDownRight)
			fmt.Println(dp, "inner")
		}
		fmt.Println(dp, "outer")
	}
	fmt.Println(dp)
	// Find the maximum value in the first column of dp
	maxGold := dp[0][0]
	for i := 1; i < rows; i++ {
		maxGold = max(maxGold, dp[i][0])
	}
	fmt.Println(maxGold)
	return maxGold
}

func GetMaxGold1(n, m int, arr [][]int) {
	maxSum := math.MinInt
	row, col := 0, 0
	for i := 0; i < n; i++ {
		if arr[i][0] >= maxSum {
			maxSum = arr[i][0]
			row = i
			col = 0
		}
	}
	i, j := row, col
	for {
		var upRight, right, downRight = math.MinInt, math.MinInt, math.MinInt
		found := false
		fmt.Println(i, j)
		if i-1 >= 0 && i-1 < n && j+1 < m {
			upRight = arr[i-1][j+1]
			found = true
		}
		if i >= 0 && i < n && j+1 < m {
			right = arr[i][j+1]
			found = true
		}
		if i+1 >= 0 && i+1 < n && j+1 < m {
			downRight = arr[i+1][j+1]
			found = true
		}
		if !found {
			fmt.Println(maxSum)
			return
		}
		max := max3(upRight, right, downRight)
		if upRight == max {
			i = i - 1
			j = j + 1
		} else if right == max {
			j = j + 1
		} else if downRight == max {
			i = i + 1
			j = j + 1
		}
		maxSum += max
	}
}
func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func LongestComPrefix(arr []string, n int) {
	// minLen := math.MaxInt
	// shtstr := ""
	// for i := 0; i < n; i++ {
	// 	if len(arr[i]) < minLen {
	// 		shtstr = arr[i]
	// 		minLen = len(arr[i])
	// 	}
	// }
	// var str = ""
	// for _, ch := range shtstr {
	// 	str = str + string(ch)
	// 	for i := 0; i < n; i++ {
	// 		if strings.Contains(arr[i], str) {
	// 			continue
	// 		} else {
	// 			if str == "" {
	// 				fmt.Println(-1)
	// 			} else {
	// 				fmt.Println(str[:len(str)-1])
	// 			}
	// 			return
	// 		}
	// 	}
	// }
	// fmt.Println(str)
	// return
	prefix := arr[0]

	for i := 1; i < n; i++ {
		fmt.Println(prefix)
		for !strings.HasPrefix(arr[i], prefix) {
			if len(prefix) > 0 {
				prefix = prefix[:len(prefix)-1]
			} else {
				break
			}
		}
	}

	if len(prefix) == 0 {
		fmt.Print(-1)
		return
	}
	fmt.Println(prefix)
}

func MinRepeats(a, b string) {
	for _, c := range a {
		if !strings.ContainsRune(b, c) {
			fmt.Println(-1)
			return
		}
	}
	for _, c := range b {
		if !strings.ContainsRune(a, c) {
			fmt.Println(-1)
			return
		}
	}
	var str string = ""
	cnt := 0
	for {
		str += a
		cnt++
		if strings.Contains(str, b) {
			fmt.Println(cnt)
			return
		}
	}
}
func findRelativeRanks(score []int) []string {
	type pair struct {
		val int
		ind int
	}
	scores := []pair{}
	for i, v := range score {
		scores = append(scores, pair{v, i})
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].val > scores[j].val
	})
	ans := make([]string, len(scores))
	for i := 0; i < len(scores); i++ {
		if i == 0 {
			ans[scores[0].ind] = "Gold Medal"
		} else if i == 1 {
			ans[scores[1].ind] = "Silver Medal"
		} else if i == 2 {
			ans[scores[2].ind] = "Bronze Medal"
		} else {
			ans[scores[i].ind] = strconv.Itoa(scores[i].val)
		}
	}
	return ans
}
