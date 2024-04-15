package problems

import (
	"fmt"
	"strconv"
)

func SumOfDependencies(V int, adjList map[int][]int) int {
	cnt := 0
	// for i := 0; i < V; i++ {
	// 	n := len(adjList[i])
	// 	cnt += n
	// }
	for _, val := range adjList {
		cnt += len(val)
	}
	return cnt
}
func CountX(L, R, X int) {
	cnt := 0
	for i := L + 1; i < R; i++ {
		cnt += countOccr(i, X)
	}
	fmt.Println(cnt)
}

func countOccr(i, x int) int {
	searchSpace := strconv.Itoa(i)
	searchStr := strconv.Itoa(x)
	cnt := 0
	for _, v := range searchSpace {
		if string(v) == searchStr {
			cnt++
		}
	}
	return cnt
}
