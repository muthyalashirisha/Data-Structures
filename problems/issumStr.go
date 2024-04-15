package problems

import (
	"fmt"
	"strconv"
	"strings"
)

func IsSumStr(s string) int {
	for i, _ := range s {
		for j := i + 1; j < len(s); j++ {
			if sol(s, 0, i, j) {
				return 1
			}
		}
	}
	return 0
}

// def is_valid_sum(i, j, k):
//             nonlocal input_string, string_length

//             if k == string_length:
//                 return True
//             if k > string_length:
//                 return False

//             num1 = int(input_string[i:j])
//             num2 = int(input_string[j:k])
//             sum_str = str(num1 + num2)

//             if not input_string[k:].startswith(sum_str):
//                 return False

//             return is_valid_sum(j, k, k + len(sum_str))

func sol(s string, i, j ,k int) bool {
	n := len(s)
	if k == n {
		return true
	}
	if k > n {
		return false
	}
	fmt.Println(i,j,k)
	num1, _ := strconv.Atoi(s[i:j])
	num2, _ := strconv.Atoi(s[j:k])
	sum := strconv.Itoa(num1 + num2)
	if strings.Contains(s[j:], sum) {
		return false
	}
	return sol(s, j, k, k+len(sum))
}
