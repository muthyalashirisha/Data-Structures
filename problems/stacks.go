package problems

import (
	"container/list"
	"fmt"
)

func NextGreaterEle(arr []int) {
	n := len(arr)
	result := make([]int, len(arr))
	stack := list.New()
	for i := n - 1; i >= 0; i-- {
		val := arr[i]
		for stack.Len() > 0 && val >= stack.Back().Value.(int) {
			stack.Remove(stack.Back())
		}
		if stack.Len() > 0 {
			result[i] = stack.Back().Value.(int)
		} else {
			result[i] = -1
		}
		stack.PushBack(val)
	}
	fmt.Println(result)
}
func PrevGreaterEle(arr []int) {
	n := len(arr)
	result := make([]int, len(arr))
	stack := list.New()
	for i := 0; i < n; i++ {
		val := arr[i]
		for stack.Len() > 0 && val >= stack.Back().Value.(int) {
			stack.Remove(stack.Back())
		}
		if stack.Len() > 0 {
			result[i] = stack.Back().Value.(int)
		} else {
			result[i] = -1
		}
		stack.PushBack(val)
	}
	fmt.Println(result)
}
func NextLesserEle(arr []int) {
	n := len(arr)
	result := make([]int, len(arr))
	stack := list.New()
	for i := n - 1; i >= 0; i-- {
		val := arr[i]
		for stack.Len() > 0 && val <= stack.Back().Value.(int) {
			stack.Remove(stack.Back())
		}
		if stack.Len() > 0 {
			result[i] = stack.Back().Value.(int)
		} else {
			result[i] = -1
		}
		stack.PushBack(val)
	}
	fmt.Println(result)
}
func PrevLesserEle(arr []int) {
	n := len(arr)
	result := make([]int, len(arr))
	stack := list.New()
	for i := 0; i < n; i++ {
		val := arr[i]
		for stack.Len() > 0 && val <= stack.Back().Value.(int) {
			stack.Remove(stack.Back())
		}
		if stack.Len() > 0 {
			result[i] = stack.Back().Value.(int)
		} else {
			result[i] = -1
		}
		stack.PushBack(val)
	}
	fmt.Println(result)
}

var minEle int

var stack []int

func Constructors() []int {
	return []int{}
}
func StackOpe() {
	pushs(4)
	pushs(3)
	pushs(2)
	pushs(1)
	pushs(0)
	fmt.Println(stack)
	fmt.Println(getMin())
	fmt.Println(pop())
	fmt.Println(stack)
	pushs(0)
	pushs(1)
	pushs(3)
	pushs(0)
	fmt.Println(stack)
	fmt.Println(getMin())
	fmt.Println(stack, "1")
	fmt.Println(pop())
	fmt.Println(stack, "2")
	fmt.Println(getMin())
	fmt.Println(pop())
	fmt.Println(stack, "3")
	fmt.Println(pop())
	fmt.Println(stack, "4")
}
func pushs(x int) {
	if len(stack) <= 0 {
		stack = append(stack, x)
		minEle = x
		return
	}

	// s2 := []int{}
	// for len(stack) > 0 && stack[len(stack)-1] > x {
	// 	s2 = append(s2, stack[len(stack)-1])
	// 	stack = stack[:len(stack)-1]
	// }
	// for len(s2) > 0 {
	// 	stack = append(stack, s2[len(s2)-1])
	// 	s2 = s2[:len(s2)-1]
	// }
	// stack = append(stack, x)
	// minEle = stack[len(stack)-1]
	if minEle >= x {
		stack = append(stack, minEle)
		minEle = x
	}
	stack = append(stack, x)
}
func pop() int {
	if len(stack) <= 0 {
		return -1
	}
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if x == minEle && len(stack) > 0 {
		minEle = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return x
}
func getMin() int {
	if len(stack) <= 0 {
		return -1
	}
	return minEle
}

func CelebrityGraph(M [][]int, N int) int {
	indegree := make([]int, N)
	outdegree := make([]int, N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if M[i][j] == 1 {
				outdegree[i]++
				indegree[j]++
			}
		}
	}

	for i := 0; i < N; i++ {
		if indegree[i] == N-1 && outdegree[i] == 0 {
			return i
		}
	}
	return -1
}
func CelebrityStack(M [][]int, N int) int {
	stack := []int{}

	for i := 0; i < N; i++ {
		stack = append(stack, i)
	}

	for len(stack) > 1 {
		row := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		col := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if M[row][col] == 1 {
			stack = append(stack, col)
		} else {
			stack = append(stack, row)
		}
	}
	fmt.Println(stack)
	x := stack[len(stack)-1]
	fmt.Println(x)
	// if len(stack) == 1 {
	// 	return stack[0]
	// }

	for j := 0; j < N; j++ {
		if x != j && M[x][j] == 1 {
			return -1
		}

	}

	for i := 0; i < N; i++ {
		if x != i && M[i][x] == 0 {
			return -1
		}
	}

	return x
}
func SortStack(stack []int) []int {
	recurSortStack(&stack)
	fmt.Println(stack)
	res := []int{}
	reverseRecStack(&stack, &res)
	fmt.Println(res)
	return stack
}
func recurSortStack(stack *[]int) {
	if len(*stack) <= 0 {
		return
	}
	ele := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	recurSortStack(stack)
	insertEleSorted(stack, ele)
}
func insertEleSorted(stack *[]int, ele int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= ele {
		*stack = append(*stack, ele)
		return
	}
	temp := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	insertEleSorted(stack, ele)
	*stack = append(*stack, temp)
}
func reverseStack(arr []int) []int {
	res := []int{}
	reverseRecStack(&arr, &res)
	fmt.Println(arr, res)
	return arr
}
func reverseRecStack(stack *[]int, res *[]int) {
	if len(*stack) <= 0 {
		return
	}
	ele := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	*res = append(*res, ele)
	reverseRecStack(stack, res)
}
func SubArrayRanges(n []int) int64 {
	return sumSubarrayComp(n, func(a, b int) bool { return a < b }) - sumSubarrayComp(n, func(a, b int) bool { return a > b })
}

func sumSubarrayComp(n []int, comp func(int, int) bool) int64 {
	var res int64 = 0
	var s []int

	for i := 0; i <= len(n); i++ {
		fmt.Println(s)
		for len(s) > 0 && (i == len(n) || comp(n[s[len(s)-1]], n[i])) {
			// fmt.Println(comp(n[s[len(s)-1]], n[i]))
			j, k := s[len(s)-1], -1
			if len(s) >= 2 {
				k = s[len(s)-2]
			}
			fmt.Println(j, k, i, n[j])

			res += int64(i-j) * int64(j-k) * int64(n[j])
			s = s[:len(s)-1]
		}

		s = append(s, i)
	}

	return res
}
