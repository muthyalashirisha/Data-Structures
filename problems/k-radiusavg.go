package problems

import (
	"fmt"
	"math"
)

func GetAverages(nums []int, k int) []int {
	n := len(nums)
	if k == 0 {
		return nums
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	prefSum := make([]int, n+1)

	// Compute prefix sum
	for i := 0; i < n; i++ {
		prefSum[i+1] = prefSum[i] + nums[i]
	}
	fmt.Println(prefSum)
	// Calculate averages
	for i := k; i < n-k; i++ {
		totalSum := prefSum[i+k+1] - prefSum[i-k]
		ans[i] = totalSum / (2*k + 1)
	}

	return ans
}

func CountGoodSubstrings(s string) int {
	k := 3
	n := len(s)

	if n < k {
		return 0
	}

	count := 0
	freq := make(map[byte]int)
	ans := map[string]int{}

	// Initialize the frequency map for the first window
	for i := 0; i < k; i++ {

		freq[s[i]]++
	}

	// Check if the first window is a good substring
	if len(freq) == k {
		ans[s[0:k]]++
		count++
	}

	// Slide the window across the string
	for i := k; i < n; i++ {
		// Add the new character to the frequency map
		freq[s[i]]++
		// Remove the character that is no longer in the window
		freq[s[i-k]]--
		if freq[s[i-k]] == 0 {
			delete(freq, s[i-k])
		}

		// Check if the current window is a good substring
		if len(freq) == k {
			ans[s[i-k+1:i+1]]++
			count++
		}
	}
	fmt.Println(ans)

	return count
}
func NumberOfSubstrings(s string) int {
	count, left := 0, 0
	m := make(map[byte]int)
	m['a'], m['b'], m['c'] = 0, 0, 0

	for right := 0; right < len(s); right++ {
		m[s[right]]++

		for m['a'] > 0 && m['b'] > 0 && m['c'] > 0 {
			m[s[left]]--
			left++
		}
		// Increment count by the number of substrings ending at the current right index
		count += left
		fmt.Println(count, right)
	}

	return count
}
func LongestSubarray(nums []int, limit int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	l, r := 0, 1
	cur_mx := nums[0]
	cur_mn := nums[0]
	max_l := 1

	for l <= r && r < len(nums) {
		cur_mx = max(cur_mx, nums[r])
		cur_mn = min(cur_mn, nums[r])

		if cur_mx-cur_mn <= limit {
			max_l = max(max_l, r-l+1)
		} else {
			if nums[l] == cur_mx {
				cur_mx = maxSlice(nums[l+1 : r+1])
			}
			if nums[l] == cur_mn {
				cur_mn = minSlice(nums[l+1 : r+1])
			}
			l++
		}

		r++
	}

	return max_l
}
func maxSlice(slice []int) int {
	maxVal := math.MinInt64
	for _, num := range slice {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func minSlice(slice []int) int {
	minVal := math.MaxInt64
	for _, num := range slice {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}
func MaxVowels(s string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	ans := 0
	v := 0

	// Initial window
	// for i := 0; i < k; i++ {
	// 	if vowels[s[i]] {
	// 		v++
	// 	}
	// }
	// ans = v

	// Sliding window
	for i := 0; i < len(s); i++ {
		if i >= k && vowels[s[i-k]] {
			v--
		}
		if vowels[s[i]] {
			v++
		}
		if i >= k && v > ans {
			ans = v
		}
	}

	return ans
}
