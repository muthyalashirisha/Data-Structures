package problems

import (
	"fmt"
	"strings"
)

func FruitsInBasket(fruits []int) int {
	left := 0
	right := 0
	globalMax := 0
	currentMax := 0
	currFruits := map[int]int{}
	n := len(fruits)
	for right = 0; right < n; right++ {
		currFruits[fruits[right]]++
		for len(currFruits) > 2 {
			currFruits[fruits[left]]--
			if currFruits[fruits[left]] == 0 {
				delete(currFruits, fruits[left])
			}
			left++
		}
		currentMax = right - left + 1
		globalMax = max(globalMax, currentMax)
	}
	fmt.Println(globalMax)
	return currentMax
}
func MaxSubarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	ans, sum := 0, 0
	mp[0] = -1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		r := sum % k
		if r < 0 {
			r += k
		}
		fmt.Println(r, sum, mp, i, ans)
		if _, ok := mp[r]; !ok {
			mp[r] = i
		} else {
			ans = max(ans, i-mp[r])
		}
	}
	fmt.Println(ans)
	return ans
}

func RemoveKdigits(s string, k int) {
	//10002996
	fmt.Println(recur(k, s))
	strings.TrimPrefix(s, "0")
	fmt.Println(removeloop(s, k))
	fmt.Println(removeUsingStack(s, k))
}
func recur(k int, s string) string {
	if k == 0 {
		return s
	}

	length := len(s)
	if k >= length {
		return ""
	}

	minIndex := 0
	for i := 1; i <= k; i++ {
		if s[i] < s[minIndex] {
			minIndex = i
		}
	}

	return string(s[minIndex]) + recur(k-minIndex, s[minIndex+1:])
}
func removeloop(s string, k int) string {
	stack := ""

	temp := s
	for len(temp) > 0 {
		fmt.Println(temp)
		if k == 0 {
			stack += temp
			break
		}
		if k > len(temp) {
			break
		}

		minIndex := 0
		for i := 1; i <= k; i++ {
			if temp[i] < temp[minIndex] {
				minIndex = i
			}
		}
		stack += string(temp[minIndex])
		temp = temp[minIndex+1:]
		k = k - minIndex
	}
	return stack
}
func removeUsingStack(s string, k int) string {
	stack := []byte{}

	if len(s) <= k {
		return "0"
	}
	if k == 0 {
		return s
	}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		for len(stack) > 0 && stack[len(stack)-1] > ch && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		if len(stack) > 0 || ch != '0' {
			stack = append(stack, ch)
		}
	}
	for len(stack) > 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	if len(stack) <= 0 {
		return "0"
	}
	res := ""
	for i := 0; i < len(stack); i++ {
		res += string(stack[i])
	}
	return res
}

func asteriods(arr []int) []int {
	return []int{}
}

//	func max(a, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	st := []byte{}
	for i := 0; i < len(num); i++ {
		for k > 0 && len(st) > 0 && st[len(st)-1] > num[i] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, num[i])
	}
	st = st[:len(st)-k]
	for len(st) > 1 && st[0] == '0' {
		st = st[1:]
	}
	return string(st)
}
