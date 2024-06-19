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

func SearchOccurence(s, p string) int {
	n := len(s)
	m := len(p)
	lps := computePrefixFunction(p)
	q := 0
	cnt := 0
	fmt.Println(lps)
	for i := 0; i < n; i++ {
		for q > 0 && s[i] != p[q] {
			q = lps[q-1]
		}
		if s[i] == p[q] {
			q++
		}
		if q == m {
			cnt++
			q = lps[m-1]
		}
	}
	return cnt
}

func computePrefixFunction(p string) []int {
	n := len(p)
	lps := make([]int, n)
	lps[0] = 0
	q := 0
	for i := 1; i < n; i++ {
		for q > 0 && p[q] != p[i] {
			q = lps[q-1]
		}
		if p[q] == p[i] {
			q++
		}
		lps[i] = q
	}
	return lps
}

func KmpAlgo(s string, p string) int {
	m := len(p)
	n := len(s)
	fail := failureLinks1(p)
	q := 0
	i := 0
	fmt.Println(fail)
	cnt := 0
	for i < n {
		if s[i] == p[q] {
			i++
			q++
			if q == m {
				return i - q
			}
		} else {
			if q >= 1 {
				q = fail[q-1]
			} else {
				i++
			}
		}
	}
	return cnt
}
func failureLinks(p string) []int {
	m := len(p)
	fail := make([]int, m)
	fail[0] = 0
	x := 0
	for j := 1; j < m; j++ {
		fail[j] = x
		for p[x] != p[j] {
			if x == 0 {
				x = -1
				break
			} else {
				x = fail[x]
			}
		}
		x++
	}
	return fail
}
func KmpAlgo1(s string, p string) int {
	m := len(p)
	n := len(s)
	fail := failureLinks1(p)
	q := 0
	i := 0
	cnt := 0
	fmt.Println(fail)
	for i < n {
		if s[i] == p[q] {
			i++
			q++
			if q == m {
				cnt++
				q = fail[q-1] // reset q to continue searching for next match
			}
		} else {
			if q >= 1 {
				q = fail[q-1] // correct the array access here
			} else {
				i++
			}
		}
	}
	return cnt
}

func failureLinks1(p string) []int {
	m := len(p)
	fail := make([]int, m)
	fail[0] = 0
	x := 0
	for j := 1; j < m; j++ {
		for x > 0 && p[x] != p[j] {
			x = fail[x-1]
		}
		if p[x] == p[j] {
			x++
		}
		fail[j] = x
	}
	return fail
}

func ZAlgo(s string) []int {
	n := len(s)
	l := 0
	r := 0
	z := make([]int, n)
	for i := 1; i < n; i++ {
		if i < r {
			z[i] = min(r-i, z[i-l])
			fmt.Println("into")
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		fmt.Println(z[i], i, i+z[i], l, r, z)
		if i+z[i] > r {
			l = i
			r = i + z[i]
		}
	}
	return z
}
func SearchZ(s, p string) {
	z := ZAlgo("Geeks#Geeks For Geeks")
	fmt.Println(z)
}
