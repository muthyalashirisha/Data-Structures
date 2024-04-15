package problems

import "fmt"

func PossibleWays(n int) int {
	// cnt := count(n, -1)
	// fmt.Println(cnt)
	// // return cnt * cnt
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 1
	dp[0][1] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-1][1]
		dp[i][1] = dp[i-1][0]
	}
	fmt.Println(dp[n][0])
	// return dp[n][0]
	prev := []int{1, 1}
	curr := []int{0, 0}
	for i := 1; i <= n; i++ {
		curr[0] = prev[0] + prev[1]
		curr[1] = prev[0]
		copy(prev, curr)
	}
	fmt.Println(prev[0])
	return prev[0]
}

//0 for space 1  for building
func count(i int, building int) int {
	if i == 0 {
		return 1
	}
	if i < 0 {
		return 0
	}
	space := count(i-1, 0)
	build := 0
	if building == 0 || building == -1 {
		build = count(i-1, 1)
	}
	return space + build
}

// // Calculate the number of ways when a building is not constructed on the current plot
// waysWithoutBuilding := count(i-1, building)

// // Calculate the number of ways when a building is constructed on the current plot
// waysWithBuilding := 0

// // If the previous plot on the same side has no building, a building can be constructed
// if !building {
// 	waysWithBuilding += count(i-1, !building)
// }
// return waysWithoutBuilding + waysWithBuilding
