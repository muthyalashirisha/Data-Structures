package problems

import (
	"fmt"
	"sort"
)

// type MinHeap []int

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// func kthSmallestMinHeap(arr []int, k int) int {
// 	if k > 0 && k <= len(arr) {
// 		minHeap := make(MinHeap, k)
// 		copy(minHeap, arr[:k])
// 		heap.Init(&minHeap)

// 		for i := k; i < len(arr); i++ {
// 			if arr[i] > minHeap[0] {
// 				heap.Pop(&minHeap)
// 				heap.Push(&minHeap, arr[i])
// 			}
// 		}

// 		return minHeap[0]
// 	}

// 	return -1 // If K is out of bounds
// }

func CountBinaryStrings(N int) int {
	const MOD = 1000000007

	// Initialize the arrays to store counts of strings ending in 0 and 1
	endInZero := make([]int, N+1)
	endInOne := make([]int, N+1)

	// Base cases
	endInZero[1] = 1
	endInOne[1] = 1
	// Iterate from length 2 to N
	for i := 2; i <= N; i++ {
		// Count of strings ending in 0 is the sum of strings ending in 0 and 1 from the previous length
		endInZero[i] = (endInZero[i-1] + endInOne[i-1]) % MOD
		// Count of strings ending in 1 is the count of strings ending in 0 from the previous length
		endInOne[i] = endInZero[i-1] % MOD
	}

	// The total count is the sum of counts of strings ending in 0 and 1 for the given length
	totalCount := (endInZero[N] + endInOne[N]) % MOD

	return totalCount
}

func GetMinDiff(arr []int, n, k int) int {
	// Check if the array has only one element
	if n == 1 {
		return 0
	}

	// Sort the array
	sort.Ints(arr)

	// Initialize result
	ans := arr[n-1] - arr[0]

	// Handle the corner elements
	small := arr[0] + k
	big := arr[n-1] - k
	if small > big {
		small, big = big, small
	}

	// Traverse the middle elements and update result
	for i := 1; i < n-1; i++ {
		subtract := arr[i] - k
		add := arr[i] + k

		// If both subtraction and addition do not change the difference
		if subtract >= small || add <= big {
			continue
		}

		// Update either subtraction or addition based on which minimizes the difference
		fmt.Println(big, subtract, add, small)
		if big-subtract <= add-small {
			small = subtract
		} else {
			big = add
		}
	}

	return min(ans, big-small)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		fmt.Println(leftMax[i-1])
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	fmt.Println(leftMax)

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	fmt.Println(rightMax)
	trappedWater := 0
	for i := 0; i < n; i++ {
		fmt.Println(trappedWater)
		trappedWater += min(leftMax[i], rightMax[i]) - height[i]
	}

	return trappedWater
}
