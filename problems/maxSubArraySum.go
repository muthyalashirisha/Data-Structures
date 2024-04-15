package problems

import (
	"fmt"
	"math"
)

func MaxSumSubarray(nums []int) int {
	// maxSum := 0
	// n := len(nums)
	// arr := [][]int{}
	// for i := 0; i < n; i++ {
	// 	for j := i; j < n; j++ {
	// 		sum := 0
	// 		arr1 := []int{}
	// 		for k := i; k <= j; k++ {
	// 			sum += nums[k]
	// 			arr1 = append(arr1, nums[k])
	// 		}
	// 		arr = append(arr, arr1)
	// 		maxSum = max(sum, maxSum)
	// 	}
	// }
	// fmt.Println("maxsum", maxSum)
	// return arr

	// maxSum := nums[0]
	// n := len(nums)
	// for i := 0; i < n; i++ {
	// 	sum := 0
	// 	for j := i; j < n; j++ {
	// 		sum += nums[j]
	// 		if maxSum < sum {
	// 			maxSum = sum
	// 		}
	// 	}
	// }
	// return maxSum
	n := len(nums)
	sum := 0
	fmt.Println("arr", nums)
	maxsum := math.MinInt
	left := 0
	right := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum > maxsum {
			maxsum = sum
			right = i
		}
		if sum < 0 {
			sum = 0
			left = i
		}
	}
	fmt.Println(left, right)
	return maxsum
}
func RearrangeArray(nums []int) []int {
	pos := []int{}
	neg := []int{}
	res := []int{}
	for _, val := range nums {
		if val >= 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
		}
	}
	i := 0

	for i < len(pos) {
		res = append(res, pos[i])
		res = append(res, neg[i])
		i++
	}
	return res
}
