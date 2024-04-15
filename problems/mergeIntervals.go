package problems

import (
	"sort"
)

func MergeNewIntervals(arr [][]int, new []int) [][]int {
	// ([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	n := len(arr)
	ans := [][]int{}

	for i := 0; i < n; i++ {
		if new[1] < arr[i][0] { //if new interval befor current interval
			ans = append(ans, new)
			for j := i; j < n; j++ {
				ans = append(ans, arr[i])
			}
			return ans
		} else if new[0] > arr[i][1] { //if new interval after the current interval
			ans = append(ans, arr[i])
		} else {
			new[0] = min(arr[i][0], new[0])
			new[1] = max(arr[i][1], new[1])
		}
	}
	ans = append(ans, new)
	return ans
}
func Merge(intervals [][]int) [][]int {
	ans := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] < intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]
		if len(ans) <= 0 || start > ans[len(ans)-1][1] {
			ans = append(ans, interval)
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], end)
		}
	}
	return ans
}
