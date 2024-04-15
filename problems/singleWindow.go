package problems

import "fmt"

func SingleElement(arr []int) int {
	cnt := 0
	result := 0
	// size := 32
	x := 0
	for i := 0; i < 32; i++ {
		cnt = 0
		x = 1 << i
		for j := 0; j < len(arr); j++ {
			if arr[j]&x != 0 {
				cnt++
			}
		}
		if cnt%3 != 0 {
			result |= x
		}
	}
	return result
}

func Pattern(n int) {
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			// if i == 0 || i == 2*n-2 {
			// 	fmt.Print(n)
			// } else if j == 0 || j == 2*n-2 {
			// 	fmt.Print(n)
			// } else {
			// 	fmt.Print(n - 1)
			// }
			fmt.Print(n - i)
		}
		fmt.Println()
	}
}
