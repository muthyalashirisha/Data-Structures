package problems

import "fmt"

func SearchSlidingWindow(pat, txt string) []int {
	if len(pat) > len(txt) {
		return []int{-1}
	}
	start := 0
	end := 0
	windowSize := len(pat)
	ans := []int{}
	str := ""
	for end < len(txt) {
		str += string(txt[end])
		if len(str) > windowSize {
			str = str[1:]
			start++
		}
		if str == pat {
			ans = append(ans, start+1)
		}
		end++
	}
	// fmt.Println(ans)
	if len(ans) <= 0 {
		return []int{-1}
	}
	return ans
}

func Search(pat, txt string) []int {
	var indices []int
	n, m := len(txt), len(pat)
	lps := buildLPS(pat)
	fmt.Println(lps)
	i, j := 0, 0
	for i < n {
		if pat[j] == txt[i] {
			i++
			j++
		}
		fmt.Println(i, j)
		if j == m {
			indices = append(indices, i-j+1)
			j = lps[j-1]
		} else if i < n && pat[j] != txt[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	fmt.Println(indices)
	return indices
}

func buildLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length, i := 0, 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}
