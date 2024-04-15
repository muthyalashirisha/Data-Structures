package problems

import (
	"container/list"
	"fmt"
	"math"
)

func SmallestSubString(S string) int {
	cnt0 := 0
	cnt1 := 0
	cnt2 := 0
	left := 0
	i := 0
	length := math.MaxInt
	for i = 0; i < len(S); i++ {
		if S[i] == '0' {
			cnt0++
		}
		if S[i] == '1' {
			cnt1++
		}
		if S[i] == '2' {
			cnt2++
		}
		for cnt0 >= 1 && cnt1 >= 1 && cnt2 >= 1 {
			length = min(length, i-left+1)
			if S[left] == '0' {
				cnt0--
			}
			if S[left] == '1' {
				cnt1--
			}
			if S[left] == '2' {
				cnt2--
			}
			left++
		}
	}

	if length == math.MaxInt {
		return -1
	} else {
		return length
	}
}

// Function to find maximum of each subarray of size k.
func Max_of_subarrays(arr []int, n int, k int) []int {
	left := 0
	result := []int{}
	numbers := []int{}
	right := 0
	for right < n {
		numbers = append(numbers, arr[right])
		right++
		if right-left+1 > k {
			result = append(result, maxK(numbers))
			numbers = numbers[1:]
			left++
		}
	}
	return result
	// fmt.Println(result)
}

func Max_of_subarrays2(arr []int, n int, k int) []int {
	var result []int
	deque := list.New()

	for i := 0; i < n; i++ {
		// Remove elements that are outside the current window
		for deque.Len() > 0 && deque.Front().Value.(int) < i-k+1 {
			deque.Remove(deque.Front())
			// fmt.Println(k)
		}

		// Remove elements that are smaller than the current element
		for deque.Len() > 0 && arr[deque.Back().Value.(int)] < arr[i] {
			k := deque.Remove(deque.Back())
			fmt.Println(k)
		}

		// Add the current index to the deque
		deque.PushBack(i)

		// Add the maximum element for the current window to the result
		if i >= k-1 {
			fmt.Println(deque.Front().Value.(int), arr[deque.Front().Value.(int)], i)
			result = append(result, arr[deque.Front().Value.(int)])
		}
	}

	return result
}

func maxK(numbers []int) int {
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
	}
	return max
}

func SmallestWindow(s string, p string) string {
	left := 0
	right := 0
	minLen := math.MaxInt32
	index := -1
	cnt := 0
	mp := map[byte]interface{}{}
	ms := map[byte]interface{}{}
	for i := 0; i < len(p); i++ {
		if val, ok := mp[s[i]]; ok {
			mp[p[right]] = val.(int) + 1
		} else {
			mp[p[i]] = 1
		}
	}

	for right < len(s) {
		if val, ok := ms[s[right]]; ok {
			ms[s[right]] = val.(int) + 1
		} else {
			ms[s[right]] = 1
		}
		if mp[s[right]] == ms[s[right]] {
			cnt++
			// fmt.Print(cnt)
		}
		// fmt.Println(ms)
		if cnt == len(p) {
			fmt.Println(cnt)
			for ms[s[left]].(int) > mp[s[left]].(int) || mp[s[left]] == 0 {
				if ms[s[left]].(int) > mp[s[left]].(int) {
					ms[s[left]] = ms[s[left]].(int) - 1
				}
				left++
			}
			l := right - left + 1
			// length = int(math.Max(float64(l), float64(length)))
			if minLen > l {
				minLen = l
				index = left
			}
		}
		right++
	}
	if index == -1 {
		return ""
	}
	return s[index:right]
}
