package problems

import (
	"fmt"
	"math"
)

func Candy(arr []int) {
	n := len(arr)
	ans := make([]int, n)
	//forward
	for i := 0; i < n; i++ {
		if i == 0 {
			ans[0] = 1
		} else if arr[i] > arr[i-1] {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = 1
		}
	}
	//backward
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			ans[i] = ans[i+1] + 1
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += ans[i]
	}
	fmt.Println(ans)
	fmt.Println(sum)
}

func LargestSum(a []float64, n int, k int) {
	var maxSum float64 = math.MinInt64

	sum := 0.0

	index := 0
	start := 0
	end := 0
	for i := 0; i < n; i++ {
		sum += a[i]
		if maxSum < sum {
			maxSum = sum
			start = index
			end = i
		}
		if sum < 0 {
			sum = 0
			index = i + 1
		}
	}
	fmt.Println(maxSum, start, end)
}

func MaxSum(s string, char []rune, b []int, n int) {
	var maxSum int = math.MinInt64
	var mp = map[rune]interface{}{}
	for i := 0; i < n; i++ {
		mp[char[i]] = b[i]
	}
	sum := 0
	index := 0
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		if mp[rune(s[i])] != nil {
			sum += mp[rune(s[i])].(int)
		} else {
			sum += int(s[i])
		}
		if maxSum < sum {
			maxSum = sum
			start = index
			end = i
		}
		if sum < 0 {
			sum = 0
			index = i + 1
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		ans += string(s[i])
	}
	fmt.Println(maxSum, start, end, ans)
}

func LargestSumDP(a []float64, n int, k int) {
	var dp = make([]float64, n)
	dp[0] = a[0]
	var ans float64 = math.MinInt64
	for i := 1; i < n; i++ {
		dp[i] = math.Max(a[i], dp[i-1]+a[i])
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	fmt.Println(ans)
}

func LargestSumWithK(a []float64, n int, k int) {

	var maxPreSum = make([]float64, n)
	maxPreSum[0] = a[0]
	sum := a[0]
	for i := 1; i < n; i++ {
		sum = math.Max(a[i], sum+a[i])
		maxPreSum[i] = sum
	}

	sum = 0
	for i := 0; i < k; i++ {
		sum += a[i]
	}

	// Use the concept of sliding window
	result := sum
	for i := k; i < n; i++ {
		// Compute sum of k elements ending
		// with a[i].
		sum = sum + a[i] - a[i-k]

		// Update result if required
		result = math.Max(result, sum)

		// Include maximum sum till [i-k] also
		// if it increases overall max.
		result = math.Max(result, sum+maxPreSum[i-k])
	}
	fmt.Println(result)
}
