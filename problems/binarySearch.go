package problems

import (
	"fmt"
	"sort"
)

func BinarySearch() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// n:=len(s)
	// var dp  [5][5]int
	ind := sort.Search(len(arr), func(j int) bool {
		return arr[j] >= 3
	})
	if ind < len(arr) && arr[ind] == 2 {
		fmt.Println("true")
	} else {
		fmt.Println(ind, len(arr))
	}
	// var s string
	// fmt.Printf("Pick an integer from 0 to 100.\n")
	// answer := sort.Search(100, func(i int) bool {
	// 	fmt.Printf("Is your number <= %d? ", i)
	// 	fmt.Scanf("%s", &s)
	// 	return s != "" && s[0] == 'y'
	// })
	// fmt.Printf("Your number is %d.\n", answer)
}
