package problems

import (
	"container/heap"
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Permutations(arr []int) [][]int {
	visited := make([]int, len(arr))
	result := [][]int{}
	uniquePermutations := make(map[string]struct{})
	sort.Ints(arr)
	solve(arr, len(arr), []int{}, &result, visited, &uniquePermutations)
	return result
}
func solve(arr []int, n int, ans []int, result *[][]int, visited []int, up *map[string]struct{}) {
	if len(ans) == n {
		// Create a new slice and copy the values to avoid modifying the original one
		// tmp := make([]int, len(ans))
		// copy(tmp, ans)
		key := fmt.Sprintf("%v", ans)
		if _, exists := (*up)[key]; !exists {
			*result = append(*result, ans)
			(*up)[key] = struct{}{}
		}
		return
	}
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			ans = append(ans, arr[i])
			solve(arr, n, ans, result, visited, up)
			ans = ans[:len(ans)-1]
			visited[i] = 0
		}
	}
}

func Min_Spinklers1(gallery []int, n int) int {
	area := make([][]int, 0)
	for i := 0; i < n; i++ {
		if gallery[i] != -1 {
			left := i - gallery[i]
			right := i + gallery[i]
			if left < 0 {
				left = 0
			}
			if right > n-1 {
				right = n - 1
			}
			area = append(area, []int{left, right})
		}
	}

	m := len(area)

	var ans [][]int
	fmt.Println(area)
	for i := 0; i < m; i++ {
		// If the current interval does not overlap with the last interval
		if len(ans) == 0 || area[i][0] > ans[len(ans)-1][1] {
			ans = append(ans, []int{area[i][0], area[i][1]})
		} else { // If the current interval overlaps with the last interval
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], area[i][1])
		}
	}
	fmt.Println(ans)
	for i := 1; i < len(ans); i++ {
		if ans[i][0] != ans[i-1][1]+1 {
			return -1
		}
	}
	return len(ans)
}
func Min_Spinklers(gallery []int, n int) int {
	//[]int{1, 0, 4, 2, 0, 6, 2, 4}
	// Create a slice of intervals representing sprinkler ranges sorted by end position in descending order
	var sprinklerRanges [][]int
	for i, g := range gallery {
		if g != -1 {
			left := i - g
			right := i + g
			fmt.Println(left, right)
			if left < 0 {
				left = 0
			}
			if right > n-1 {
				right = n - 1
			}

			sprinklerRanges = append(sprinklerRanges, []int{left, right})
		}
	}

	// Sort sprinklerRanges by end position in descending order
	sort.Slice(sprinklerRanges, func(i, j int) bool {
		return sprinklerRanges[i][0] == sprinklerRanges[j][0] && sprinklerRanges[i][1] > sprinklerRanges[j][1] || sprinklerRanges[i][0] < sprinklerRanges[j][0]
	})

	start, index, answer := 0, 0, 0
	currentMax := -1

	for start < n {
		for index < len(sprinklerRanges) {
			if sprinklerRanges[index][0] > start {
				break
			}
			// fmt.Print("..", index)
			currentMax = max(currentMax, sprinklerRanges[index][1])
			index++
		}

		if currentMax < start {
			return -1
		}

		answer++
		start = currentMax + 1
	}

	return answer
}

// type pair struct {
// 	val, freq int
// }
// type maxHeap []pair

// func (h maxHeap) Less(i, j int) bool {
// 	return h[i].freq > h[j].freq || (h[i].freq == h[j].freq && h[i].val < h[j].val)
// }
// func (h maxHeap) Len() int {
// 	return len(h)
// }
// func (h maxHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }
// func (h *maxHeap) Push(x interface{}) {
// 	*h = append(*h, x.(pair))
// }
// func (h *maxHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	ele := old[n-1]
// 	*h = old[:n-1]
// 	return ele
// }
// func TopKFrequent(arr []int, N int, K int) {
// 	// var pq = []pair{}
// 	// heap.Init(pq)
// 	ans := [][]int{}
// 	mp := map[int]int{}
// 	for i := 0; i < N; i++ {
// 		mp[arr[i]]++

//			ans = append(ans, currList)
//		}
//		fmt.Println(ans)
//	}
type Element struct {
	num     int
	freq    int
	lessNum bool
}

type maxHeap []Element

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq || (h[i].freq == h[j].freq && h[i].lessNum)
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TopKFrequent(nums []int, k int) [][]int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	maxHeap := &maxHeap{}
	heap.Init(maxHeap)

	var result [][]int

	for i, num := range nums {
		heap.Push(maxHeap, Element{num, frequencyMap[num], num < nums[i]})

		// Create a copy of the heap and extract top K elements
		copyHeap := *maxHeap
		var topK []int
		for j := 0; j < k && j < len(copyHeap); j++ {
			topK = append(topK, heap.Pop(&copyHeap).(Element).num)
		}

		result = append(result, topK)
	}

	return result
}
func KTop(array []int, size, k int) [][]int {
	result := make([][]int, 0)
	currentTop := make([]int, k)   // Array to store the current top K elements
	frequency := make(map[int]int) // Map to store the frequency of elements in currentTop

	for index := 0; index < size; index++ {
		// Update the frequency of the current element in the array
		frequency[array[index]]++

		// Set the last element of currentTop as the current element from the array
		currentTop[k-1] = array[index]

		// Find the correct position for the current element in the currentTop array
		i := k - 2
		for i >= 0 {
			// Compare frequencies and values for sorting
			if frequency[currentTop[i]] < frequency[currentTop[i+1]] ||
				(frequency[currentTop[i]] == frequency[currentTop[i+1]] && currentTop[i] > currentTop[i+1]) {
				currentTop[i], currentTop[i+1] = currentTop[i+1], currentTop[i]
			} else {
				break
			}
			i--
		}

		// Extract the top K elements and append them to the result list
		i = 0
		currentResult := make([]int, 0)
		for i < k && currentTop[i] != 0 {
			currentResult = append(currentResult, currentTop[i])
			i++
		}

		// Append the current result to the result list
		result = append(result, currentResult)
	}

	return result
}
func AddOperators(num string, target int) []string {
	result := []string{}
	recursiveCall(0, "", 0, 0, num, target, &result)
	return result
}

func recursiveCall(i int, sumPath string, sum, prev int, num string, target int, result *[]string) {
	if i == len(num) {
		if sum == target {
			*result = append(*result, sumPath)
		}
		return
	}

	for j := i; j < len(num); j++ {
		if j > i && num[i] == '0' {
			// Avoid leading zeros in the number
			break
		}

		number, _ := strconv.Atoi(num[i : j+1])
		// Convert the substring to an integer
		tempPath := num[i : j+1]
		// Temporary substring representing the path we have traversed so far

		fmt.Println(sumPath)
		if i == 0 {
			// If we are on the first index of 'num', start a new path
			recursiveCall(j+1, tempPath, number, number, num, target, result)
		} else {
			// Addition operation
			recursiveCall(j+1, sumPath+"+"+tempPath, sum+number, number, num, target, result)
			// Subtraction operation
			recursiveCall(j+1, sumPath+"-"+tempPath, sum-number, -number, num, target, result)
			// Multiplication operation
			recursiveCall(j+1, sumPath+"*"+tempPath, sum-prev+(prev*number), prev*number, num, target, result)
			// While doing the multiplication operation, we remove the previous operation and update it with the multiplication operation
			// This is done to follow the BODMAS rules for correct precedence
		}
	}
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func PrintPaths(root *TreeNode, sum int) [][]int {
	ans := [][]int{}
	printPathsHelper(root, sum, 0, []int{}, &ans)
	return ans
}
func printPathsHelper(root *TreeNode, sum int, currSum int, ds []int, ans *[][]int) {
	if root == nil {
		return
	}
	ds = append(ds, root.data)
	currSum += root.data
	if sum == currSum {
		*ans = append(*ans, ds)
	}
	printPathsHelper(root.left, sum, currSum, ds, ans)
	printPathsHelper(root.left, sum, currSum, ds, ans)
	ds = ds[:len(ds)-1]
}
func IsItTree(n, m int, edges [][]int) bool {
	adjList := make([][]int, n)

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make([]int, n)
	res := dfs(0, visited, adjList, -1)
	if !res {
		return false
	}
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			return false
		}
	}
	return true
}
func dfs(node int, visited []int, adjList [][]int, parent int) bool {
	visited[node] = 1
	for _, adjNode := range adjList[node] {
		if visited[adjNode] == 0 {
			if !dfs(adjNode, visited, adjList, node) {
				return false
			}
		} else if parent != adjNode {
			return false
		}
	}
	return true
}

type MaxHeapT struct {
	items   []int
	itemMap map[int]int // Mapping of item to its index in items
}

func (h *MaxHeapT) Len() int {
	return len(h.items)
}

func (h *MaxHeapT) Less(i, j int) bool {
	return h.items[i] > h.items[j]
}

func (h *MaxHeapT) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.itemMap[h.items[i]] = i
	h.itemMap[h.items[j]] = j
}

func (h *MaxHeapT) Push(x interface{}) {
	item := x.(int)
	h.items = append(h.items, item)
	h.itemMap[item] = len(h.items) - 1
}

func (h *MaxHeapT) Pop() interface{} {
	n := len(h.items)
	item := h.items[n-1]
	h.items = h.items[0 : n-1]
	delete(h.itemMap, item)
	return item
}

func (h *MaxHeapT) Remove(item int) {
	index, exists := h.itemMap[item]
	if !exists {
		return // Item not found in the heap
	}

	// Swap the item with the last element and then pop
	h.Swap(index, len(h.items)-1)
	heap.Pop(h)
}

type Pair struct {
	first  string
	second int
}

func sieve(primeNumbers []int) {
	for i := 2; i < 10001; i++ {
		if primeNumbers[i] == 0 {
			primeNumbers[i] = i
			for j := i * i; j < 10001; j += i {
				if primeNumbers[j] == 0 {
					primeNumbers[j] = i
				}
			}
		}
	}
}

func PrimePath(n1, n2 int) int {
	if n1 == n2 {
		return 0
	}

	primeNumbers := make([]int, 10001)
	sieve(primeNumbers)
	// fmt.Println(primeNumbers)
	s1 := strconv.Itoa(n1)
	s2 := strconv.Itoa(n2)
	st := make(map[int]bool)
	q := list.New()
	q.PushBack(Pair{s1, 0})
	st[n1] = true

	fmt.Println(Pair{s1, 0})
	for q.Len() > 0 {
		p := q.Front().Value.(Pair)
		q.Remove(q.Front())

		if p.first == s2 {
			return p.second
		}

		for j := 0; j < len(p.first); j++ {
			ch := p.first[j]
			ind, _ := strconv.Atoi(string(ch))
			for i := 0; i < 10; i++ {
				if j == 0 && i == 0 {
					continue
				}
				if ind != i {
					s := p.first
					s = s[:j] + string(byte(i+48)) + s[j+1:]
					x, _ := strconv.Atoi(s)

					if primeNumbers[x] == x && !st[x] {
						st[x] = true
						q.PushBack(Pair{s, p.second + 1})
					}
				}
			}
		}
	}
	return -1
}

func LCSof3(A, B, C string, n1, n2, n3 int) int {
	dp := make([][][]int, n1)
	for i := 0; i < n1; i++ {
		dp[i] = make([][]int, n2)
		for j := 0; j < n2; j++ {
			dp[j] = make([][]int, n3)
			for k := 0; k < n3; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	return helperLcsOf3(A, B, C, n1-1, n2-1, n3-1, dp)
}
func helperLcsOf3(A, B, C string, i, j, k int, dp [][][]int) int {
	if i < 0 || j < 0 || k < 0 {
		return 0
	}
	if dp[i][j][k] != -1 {
		return dp[i][j][k]
	}
	if A[i] == B[j] && B[j] == C[k] {
		return 1 + helperLcsOf3(A, B, C, i-1, j-1, k-1, dp)
	} else {
		return helperLcsOf3(A, B, C, i-1, j, k, dp) + helperLcsOf3(A, B, C, i, j-1, k, dp) + helperLcsOf3(A, B, C, i, j, k-1, dp)
	}
}

func CheckPanagram(s string) bool {
	s = strings.ToLower(s)
	arr := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			arr[s[i]] = struct{}{}
		}
	}

	return len(arr) == 26
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z'
}
func FirstUniqChar(s string) int {

	// mp := map[byte]int{}
	// for i := 0; i < len(s); i++ {
	// 	mp[s[i]]++
	// }
	// for i := 0; i < len(s); i++ {
	// 	if mp[s[i]] == 1 {
	// 		return i
	// 	}
	// }
	// return -1
	for i := 0; i < len(s); i++ {
		ind1 := strings.Index(s, string(s[i]))
		ind2 := strings.LastIndex(s, string(s[i]))
		if ind1 == ind2 {
			return i
		}
	}
	return -1
}
func repeatedCharacter(s string) byte {
	mp := map[byte]int{}
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
		if mp[s[i]] > 1 {
			return s[i]
		}
	}
	// math.Ceil(float64(piles[i]) / float64(k))
	strings.TrimPrefix(s, " ")
	sort.Sort(sort.StringSlice{s})
	return s[0]
}
