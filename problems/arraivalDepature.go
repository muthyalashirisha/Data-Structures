package problems

import (
	"fmt"
	"math"
	"sort"
)

func MinPlatforms1(arrival []int, departure []int) int {
	n := len(arrival)

	// Create an array of events where each event has a time and a type (0 for arrival, 1 for departure)
	events := make([][2]int, 2*n)
	for i := 0; i < n; i++ {
		events[i] = [2]int{arrival[i], 0}
		events[i+n] = [2]int{departure[i], 1}
	}

	// Sort events based on time, and in case of a tie, prioritize arrivals before departures
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	fmt.Println(events)
	platformsInUse := 0
	maxPlatforms := 0

	for _, event := range events {
		if event[1] == 0 { // Arrival
			platformsInUse++
		} else { // Departure
			platformsInUse--
		}

		// Update the maximum number of platforms
		if platformsInUse > maxPlatforms {
			maxPlatforms = platformsInUse
		}
	}

	return maxPlatforms
}
func MinPlatforms(arrival []int, departure []int) int {
	sort.Ints(arrival)
	sort.Ints(departure)
	p1 := 0
	p2 := 0
	platFromInUse := 0
	maxPlatforms := -1
	for p1 < len(arrival) && p2 < len(departure) {
		if arrival[p1] == departure[p2] {
			platFromInUse++
			p1++
		} else if arrival[p1] < departure[p2] {
			platFromInUse++
			p1++
		} else {
			platFromInUse--
			p2++
		}
		if platFromInUse > maxPlatforms {
			maxPlatforms = platFromInUse
		}
	}
	return maxPlatforms
}
func PeakElement(arr []int) int {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			return i
		}
	}
	return 0
}
func PeakElementBS(arr []int) int {
	low := 0
	high := len(arr) - 1
	for low < high {
		mid := (low + high) / 2
		if arr[mid] <= arr[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return high
}
func MaxProduct(arr []int64) int64 {
	n := len(arr)

	// Initialize variables to store the current maximum and minimum product
	maxProduct := int64(arr[0])
	minProduct := int64(arr[0])
	result := int64(arr[0])

	for i := 1; i < n; i++ {
		// If the current element is negative, swap maxProduct and minProduct
		if arr[i] < 0 {
			temp := maxProduct
			maxProduct = minProduct
			minProduct = temp
		}
		fmt.Println(maxProduct, minProduct)
		// Update maxProduct and minProduct for the current element
		fmt.Println(arr[i], maxProduct*arr[i], minProduct*arr[i])
		maxProduct = int64(math.Max(float64(arr[i]), float64(maxProduct*arr[i])))
		minProduct = int64(math.Min(float64(arr[i]), float64(minProduct*arr[i])))
		fmt.Println(maxProduct, minProduct)
		// Update the overall result
		result = int64(math.Max(float64(result), float64(maxProduct)))
	}

	return result
}
func MinMissing(arr []int) {
	mp := map[int]bool{}
	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = true
	}
	i := 1
	for {
		if !mp[i] {
			fmt.Println(i)
			return
		}
		i++
	}
}
func Sum3(arr []int, target int) (int, [][]int) {
	ans := [][]int{}
	cnt := 0
	n := len(arr)
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		j := i + 1
		k := n - 1
		if i != 0 && arr[i] != arr[i-1] {
			continue
		}
		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if sum == target {
				ans = append(ans, []int{arr[i], arr[j], arr[k]})
				cnt++
				j++
				k--
				for j < k && arr[j] != arr[j-1] {
					j++
				}
				for j > k && arr[k] != arr[k+1] {
					k--
				}
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return cnt, ans
}
func StockBuySell(arr []int) [][]int {
	ans := [][]int{}
	n := len(arr)
	i := 0
	for i < n {
		if i == n-1 {
			return [][]int{}
		}
		if arr[i] < arr[i+1] {
			break
		}
		i++
	}

	j := 1
	for j < len(arr) {
		fmt.Println(ans)
		if j != n-1 && arr[j] > arr[j+1] {
			ans = append(ans, []int{i, j})
			i = j + 1
			for {
				if i == n-1 {
					return ans
				}
				if arr[i] > arr[i+1] {
					i++
				} else {
					break
				}
			}
			j = i + 1
		} else {
			j++
		}
	}
	ans = append(ans, []int{i, j - 1})
	return ans
}
func CountWays(coins []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1

	for _, coin := range coins {
		fmt.Println(coin)
		for i := coin; i <= sum; i++ {
			dp[i] += dp[i-coin]
			fmt.Println(dp, "....", i)
		}
		fmt.Println("next loop")
	}

	return dp[sum]
}
