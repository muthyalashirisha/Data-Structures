package problems

import "fmt"

func IsSubSet(arr1 []int, arr2 []int) bool {
	n := len(arr1)
	m := len(arr2)
	// if n < m {
	// 	return IsSubSet(arr2, arr1)
	// }
	// h := map[int]bool{}
	// for _, val := range arr1 {
	// 	h[val] = true
	// }
	// for _, val := range arr2 {
	// 	if !h[val] {
	// 		return false
	// 	}
	// }
	// return true
	h := make(map[int]int)
	if n < m {
		return false
	}
	for _, val := range arr1 {
		h[val]++
	}

	for _, val := range arr2 {
		if h[val] == 0 {
			return false
		}
		h[val]--
	}

	return true
}
func CountPairs(arr []int, n int, k int) int {
	// hm := map[int]interface{}{}
	// cnt := 0
	// for _, val := range arr {
	// 	if hm[k-val] != nil {
	// 		cnt++
	// 	}
	// 	hm[val] = true
	// }
	// return cnt
	seen := make(map[int]int)
	count := 0

	for _, num := range arr {
		b := k - num
		fmt.Println(b, num, seen[b], seen[num])
		count += seen[b]
		fmt.Println(count)
		seen[num]++
	}
	fmt.Println(seen)

	return count
}
func FirstRepeating(arr []int) int {
	// hm := map[int]interface{}{}
	// for index, val := range arr {
	// 	if hm[val] != nil {
	// 		return hm[val].(int)
	// 	}
	// 	hm[val] = index
	// }
	// return -1
	hm := map[int]int{}
	for _, val := range arr {
		hm[val] += 1
	}
	for index, val := range arr {
		if hm[val] > 1 {
			return index + 1
		}
	}
	return -1
}
func AlterantePosNeg(arr []int) {
	var pos []int
	var neg []int
	for _, val := range arr {
		if val >= 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
		}
	}
	i := 0
	j := 0
	k := 0
	for i < len(arr) {
		if j < len(pos) && k < len(neg) {
			if i%2 == 0 {
				arr[i] = pos[j]
				j++
			} else {
				arr[i] = neg[k]
				k++
			}
		} else if j < len(pos) {
			arr[i] = pos[j]
			j++
		} else if k < len(neg) {
			arr[i] = neg[k]
			k++
		}
		i++
	}
}
func NonRepeating(arr []int) {
	hm := map[int]int{}
	for _, value := range arr {
		hm[value]++
	}
	for _, value := range arr {
		if hm[value] < 2 {
			fmt.Println(value)
			break
		}
	}

	fmt.Println(0)
}
