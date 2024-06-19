package problems

import (
	"fmt"
	"sort"
)

func FindProductsOfElements(queries [][]int) []int {
	cnt1 := func(num int) int {
		res := 0
		for i := 0; i < 32; i++ {
			if (num >> i) == 0 {
				break
			}
			cur := num % (1 << (i + 1))
			res += (num - cur) / 2
			if cur >= (1 << i) {
				res += cur + 1 - (1 << i)
			}
		}
		return res
	}

	acc0 := func(num int) int {
		res := 0
		for i := 0; i < 32; i++ {
			if (num >> i) == 0 {
				break
			}
			cur := num % (1 << (i + 1))
			res += (num - cur) / 2 * i
			if cur >= (1 << i) {
				res += (cur + 1 - (1 << i)) * i
			}
		}
		return res
	}

	count := func(bound int) int {
		fmt.Println(bound)
		target := sort.Search(bound, func(i int) bool {
			c1 := cnt1(i)
			fmt.Println(c1, "c1")
			return c1 >= bound
		})
		fmt.Println(target, "tar")
		rest := bound - cnt1(target-1)
		cnt := acc0(target - 1)
		fmt.Println(cnt, "cnt")

		for i := 0; i < 32; i++ {
			if target&(1<<i) != 0 {
				cnt += i
				rest--
				if rest == 0 {
					break
				}
			}
		}

		return cnt
	}

	var result []int
	for _, query := range queries {
		low, high, mod := query[0], query[1], query[2]
		ch := count(high + 1)
		cl := count(low)
		p := pow(2, ch-cl, mod)
		fmt.Println(p, ch, cl)
		result = append(result, p)
	}
	return result
}

func pow(x, y, mod int) int {
	fmt.Println(x, y, mod)
	if y == 0 {
		return 1
	}
	res := 1
	for y > 0 {
		if y&1 == 1 {
			res = (res * x) % mod
		}
		y >>= 1
		x = (x * x) % mod
	}
	fmt.Println(res)
	return res
}
