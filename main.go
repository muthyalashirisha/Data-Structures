package main

import (
	"DS/problems"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"sync"
)

// func main() {
// 	// as := problems.SearchOccurence("abcabcabcabc", "abcabc")
// 	// fmt.Println(as)
// 	// A := "abcdef"
// 	// B := "xyzcba"
// 	// fmt.Println(problems.CelebrityStack([][]int{{1, 0, 1, 0}, {1, 1, 1, 1}, {0, 0, 0, 0}, {0, 1, 1, 1}}, 4)) // Output: true
// 	// queries := [][]int{{1, 3, 7}}
// 	// fmt.Println(problems.FindProductsOfElements(queries)) // Example usage
// 	res := problems.KmpAlgo("sadbutsad", "sad")
// 	fmt.Println(res, "1")
// 	res1 := problems.KmpAlgo1("abcabcabcabc", "abcabc")
// 	// problems.SearchZ("abcabcabcabc", "abcabc")
// 	fmt.Println(res1, "2")
// 	// // Sample input string
// 	// input := "this is an example for huffman encoding"

// 	// // Calculate frequency of each character
// 	// frequency := make(map[rune]int)
// 	// for _, char := range input {
// 	// 	frequency[char]++
// 	// }

// 	// // Build Huffman Tree
// 	// huffmanTree := problems.BuildHuffmanTree(frequency)

// 	// // Generate Huffman Codes
// 	// huffmanCodes := make(map[rune]string)
// 	// problems.GenerateHuffmanCodes(huffmanTree, "", huffmanCodes)

// 	// // Print Huffman Codes
// 	// fmt.Println("Character\tFrequency\tHuffman Code")
// 	// for char, freq := range frequency {
// 	// 	fmt.Printf("%c\t\t%d\t\t%s\n", char, freq, huffmanCodes[char])
// 	// }

// 	// trie := problems.NewTrie()

// 	// // Insert words into the trie
// 	// trie.Insert("apple")
// 	// trie.Insert("app")
// 	// trie.Insert("banana")

// 	// // Search for words
// 	// fmt.Println(trie.Search("apple"))   // true
// 	// fmt.Println(trie.Search("app"))     // true
// 	// fmt.Println(trie.Search("appl"))    // false
// 	// fmt.Println(trie.Search("banana"))  // true
// 	// fmt.Println(trie.Search("bananas")) // false

// 	// // Search for prefixes
// 	// fmt.Println(trie.StartsWith("app")) // true
// 	// fmt.Println(trie.StartsWith("ban")) // true
// 	// fmt.Println(trie.StartsWith("bat")) // false

// 	// ac := problems.NewAhoCorasick()
// 	// patterns := []string{"he", "she", "his", "hers"}
// 	// text := "ushers"

// 	// for _, pattern := range patterns {
// 	// 	ac.Add(pattern)
// 	// }
// 	// ac.Build()

//		// matches := ac.Search(text)
//		// for pattern, positions := range matches {
//		// 	fmt.Printf("Pattern '%s' found at positions: %v\n", pattern, positions)
//		// }
//	}
//
//	func main() {
//		musicPlayer := problems.NewMusicPlayer()
//		for i := 0; i < 10; i++ {
//			song := musicPlayer.GetRandomSong()
//			fmt.Println("Playing song:", song)
//			musicPlayer.InsertAgainToPlay()
//		}
//	}
func main() {
	// musicPlayer := problems.NewMusicPlayer()
	// for i := 0; i < 10; i++ {
	// 	song := musicPlayer.GetRandomSong()
	// 	if song == "" {
	// 		fmt.Println("No eligible song to play.")
	// 	} else {
	// 		fmt.Println("Playing song:", song)
	// 		musicPlayer.InsertAgainToPlay()
	// 	}
	// }
	// nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	// k := 3
	// result := problems.GetAverages(nums, k)
	// fmt.Println(result) // Output: [-1 2 3 4 5 6 7 8 9 -1]
	// s1 := "xyzzaz"
	// fmt.Println(problems.CountGoodSubstrings(s1)) // Output: 1

	// s2 := "abcabc"
	// fmt.Println(problems.NumberOfSubstrings(s2)) // Output: 4
	// s1 := "pbppppdlf"
	// k1 := 4
	// fmt.Println(problems.MaxVowels(s1, k1)) // Output: 3
	// fmt.Println(problems.CelebrityStack([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 3))
	// fmt.Println(problems.CelebrityStack([][]int{{1, 1, 1, 0},

	// 	{0, 0, 1, 0},

	// 	{0, 0, 0, 0},

	// 	{1, 1, 1, 0}}, 4))
	// fmt.Println(problems.SortStack([]int{4, 1, 2, 6, 5, 3, 0}))
	fmt.Println(problems.LongestPalindromeLengthKMP("abca", "cba"))
}

type Node struct {
	data int64
	next *Node
}

type TNode struct {
	data  int64
	left  *TNode
	right *TNode
}

var head *Node = nil
var head1 *Node = nil

func (node *Node) init(data int64) {
	node.data = data
	node.next = nil
}

func (node *TNode) init(data int64) {
	node.data = data
	node.left = nil
	node.right = nil
}

type MyStack struct {
	queue *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue.PushBack(x)
	// Rotate the queue to bring the newly added element to the front
	for i := 0; i < this.queue.Len()-1; i++ {
		this.queue.PushBack(this.queue.Front().Value)
		this.queue.Remove(this.queue.Front())
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.queue.Len() == 0 {
		return -1 // You can choose a different value or handle this case as appropriate.
	}
	val := this.queue.Front().Value.(int)
	this.queue.Remove(this.queue.Front())
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.queue.Len() == 0 {
		return -1 // You can choose a different value or handle this case as appropriate.
	}
	return this.queue.Front().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}

var data = "This is the base data for this program, containing some random data for testing human comprehension of a golang program"

func main1() {
	// ws := strings.Fields(data)
	// uw := make(map[string]bool)
	// cc := 0
	// for _, w := range ws {
	// 	uw[w] = true
	// 	cc += len(w)
	// }
	// fmt.Printf("ws: %d\n", len(ws))
	// fmt.Printf("cc: %d\n", cc)
	// fmt.Printf("uw: %d\n", len(uw))
	// fmt.Printf("awl: %.2f\n", float64(cc)/float64(len(ws)))
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			// fmt.Println(i)
			ch <- i
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(ch)
	res := []int{}
	for c := range ch {
		// fmt.Println(c)
		res = append(res, c)
	}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

// 	// MyValue1 := "Welcome to GeeksforGeeks"
// 	// MyValue2 := "Welcome!\nGeeksforGeeks"
// 	// MyValue3 := `Hello!GeeksforGeeks`
// 	// MyValue4 := `Hello!\nGeeksforGeeks`
// 	// fmt.Println("String 1:", MyValue1)
// 	// fmt.Println("String 2:", MyValue2)
// 	// fmt.Println("String 2:", MyValue3)
// 	// fmt.Println("String 2:", MyValue4)
// 	// ans := make([]int, 26)
// 	// mp := map[string]int{}
// 	// q := "aabbccddaa"
// 	// for i := 0; i < len(q); i++ {
// 	// 	ans[q[i]-'a']++
// 	// 	mp[string(q[i])]++
// 	// }
// 	// for i := 0; i < 26; i++ {
// 	// 	if ans[i] != 0 {
// 	// 		fmt.Println(string(i+'a'), ans[i])
// 	// 	}
// 	// }
// 	// // res:=[][]interface{}
// 	// fmt.Println(ans)
// 	// fmt.Println(mp)
// 	// Your MyStack object will be instantiated and called as such:
// 	// obj := Constructor()
// 	// fmt.Println(obj)
// 	// obj.Push(3)
// 	// fmt.Println(obj)
// 	// param_2 := obj.Pop()
// 	// param_3 := obj.Top()
// 	// param_4 := obj.Empty()
// 	// fmt.Println(param_2, param_3, param_4)

// 	// linkedList()
// 	// stack()
// 	// queue()
// 	//binaryTree()
// 	// n := 2
// 	// k := 3
// 	// minimumString := findString(n, k)
// 	// fmt.Println("Minimum Length String:", minimumString)
// 	//bfs
// 	// bfs()
// 	// problems.AVLTreeMain()
// 	// fmt.Println("Enter the number of rows:")
// 	// var n int
// 	// fmt.Scanln(&n)
// 	// data := problems.NthRowOfPascalT(5)
// 	// fmt.Println(data)
// 	// V := 4 // Number of nodes
// 	//E := 3 // Number of edges

// 	// Example graph represented as an adjacency list
// 	// graph := [][]int{
// 	// 	{0, 1},
// 	// 	{1, 2},
// 	// 	{2, 0},
// 	// }

// 	// if problems.HasCycle(graph, V) {
// 	// 	fmt.Println("The graph contains a cycle.")
// 	// } else {
// 	// 	fmt.Println("The graph does not contain a cycle.")
// 	// }
// 	// res := problems.Nby2Time([]int{2, 2, 3, 3, 3})
// 	// fmt.Println(res)
// 	// adjList := map[int][]int{}
// 	// adjList[1] = []int{2}
// 	// adjList[2] = []int{3}
// 	// adjList[3] = []int{2}
// 	// adjList[4] = []int{2}

// 	// res := problems.SumOfDependencies(5, adjList)

// 	// res := problems.RearrangeArray([]int{-1, 1})
// 	// var str = map[string]int{"str1": 5}
// 	// res := str["str"]
// 	// data := bson.A{bson.M{"k": "v", "k1": "v1"}, bson.M{"k2": "v", "k3": "v1"}, bson.M{"k4": "v", "k5": "v1"}}
// 	// data = bson.A{data[res]}
// 	// fmt.Println(data)
// 	// str["str"] = res + 1
// 	// data = bson.A{bson.M{"k": "v", "k1": "v1"}, bson.M{"k2": "v", "k3": "v1"}, bson.M{"k4": "v", "k5": "v1"}}
// 	// res = str["str"]
// 	// data = bson.A{data[res]}
// 	// fmt.Println(data)
// 	// fmt.Println(problems.IsSumStr("1111112223"))7

// 	// getMinDiff([]int{1, 8, 10, 6, 4, 6, 9, 1}, 8, 7)

// 	// problems.CountX(18, 81, 9)
// 	// problems.ListsDemo()
// 	// fmt.Println(problems.IsSubSet([]int{1, 2, 1, 3, 4, 5}, []int{1, 6, 1, 2, 3, 4, 5}))
// 	// fmt.Println(problems.CountPairs([]int{1, 5, 5, 5, 5, 7}, 6, 10))
// 	// fmt.Println(problems.FirstRepeating([]int{1, 5, 3, 4, 3, 5}))
// 	// problems.AlterantePosNeg([]int{1, 5, 3, -4, 3, 5, -1})
// 	// problems.NonRepeating([]int{-1, 5, 3, -4, 1, 5, -1})
// 	// fmt.Println(problems.SumZero([]int{4, 2, -1, -3, 6}, 5))
// 	// problems.GetMaxGold(4, 4, [][]int{{1, 3, 5, 1}, {2, 2, 4, 1}, {5, 0, 2, 3}, {0, 6, 1, 2}})
// 	// problems.GetMaxGold(5, 3, [][]int{{0, 4, 4}, {3, 2, 2}, {3, 2, 2}, {0, 0, 2}, {3, 2, 0}})
// 	// N := 5
// 	// result := problems.CountBinaryStrings(N)
// 	// fmt.Println(result)
// 	// arr := []int{1, 3, 37, 40}
// 	// n := len(arr)
// 	// k := 2

// 	// result := problems.GetMinDiff(arr, n, k)
// 	// fmt.Println(result)

// 	// arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
// 	// arr := []int{8, 0, 0, 0, 8}
// 	// result := problems.Trap(arr)
// 	// fmt.Println(result)
// 	// problems.LongestComPrefix([]string{"geeks", "gees", "geeksforgeeks"}, 3)
// 	// problems.MinRepeats("abcd", "cdabcdab")
// 	// arrival := []int{900, 940, 950, 1100, 1500, 1800}
// 	// departure := []int{910, 1200, 1120, 1130, 1900, 2000}

// 	// result := problems.MinPlatforms(arrival, departure)
// 	// fmt.Println("Minimum number of platforms required:", result)
// 	// fmt.Println(problems.PeakElementBS([]int{2, 4, 7, 7, 4}))
// 	// fmt.Println(problems.MaxProduct([]int64{6, -3, -10, 0, 2}))
// 	// problems.MinMissing([]int{1, 2, 3, 4, 5, 6})
// 	// fmt.Println(problems.Sum3([]int{1, 2, 4, 3, 6}, 10))
// 	// fmt.Println(problems.StockBuySell([]int{822, 754, 1689, 305, 214, 782, 1463, 1432, 1382, 1734, 948, 231, 210, 1676, 877, 670, 1384, 725, 1370, 412, 724, 371, 928, 358, 533, 1031, 850, 1555, 1064, 845, 1692, 514, 630, 1333, 1640, 315, 1639, 1792, 1779, 1325, 1619, 1400, 378, 145, 913, 901, 1652, 530, 1259, 880, 303, 685, 1586}))
// 	// fmt.Println(problems.CountWays([]int{1, 2, 3}, 4))
// 	// problems.Candy([]int{0, 0, 0})
// 	// problems.FindDuplicates([]int{1, 2, 1, 3, 4, 3})
// 	// Example usage
// 	// matrix := [][]int{
// 	// 	{1, 2, 3, 4},
// 	// 	{-3, -2, -1, 6},
// 	// 	{-7, -6, -3, -16},
// 	// 	{1, 7, 5, 8},
// 	// }

// 	// result := problems.FindLargestZeroSumSubmatrix(matrix)
// 	// // Display the result
// 	// fmt.Println("Largest Zero Sum Submatrix:")
// 	// for _, row := range result {
// 	// 	fmt.Println(row)
// 	// }
// 	// problems.AntiDiagonal(matrix)
// 	// fmt.Println(problems.Match("uvm*lxusn?r",
// 	// 	"uvmewlfxusnhr"))
// 	// fmt.Println(problems.MaxMeetings([]int{1, 3, 0, 5, 8, 5}, []int{2, 4, 6, 7, 9, 9}, 6))
// 	// problems.LargestSumWithK([]float64{1, -3, 0, 5, 8, 5}, 6, 4)
// 	// problems.LargestSumDP([]float64{1, 2, -3, 4, 5, -6}, 6, 4)
// 	// problems.MaxSum("abcde", []rune{'c'}, []int{-1000}, 1)
// 	// problems.KthLargestSum([]int{2, 6, 4, 1}, 4, 3)
// 	// problems.KthLargestSum2([]int{2, 6, 4, 1}, 4, 3)
// 	// problems.MaxSumWOAdj([]int{3, 2, 7, 10}, 4)
// 	// fmt.Println(problems.SmallestSubString("121222"))
// 	// fmt.Println(problems.Max_of_subarrays2([]int{1, 2, 3, 1, 4, 5, 2, 3, 6}, 9, 4))
// 	// fmt.Println(problems.SmallestWindow("this is a test string", "t stri"))
// 	// fmt.Println(problems.SingleElement([]int{1, 1, 10, 1}))
// 	// problems.Pattern(4)
// 	// problems.PossibleWays(3)
// 	// problems.ConvertMintoMax([]int{1, 2, 3, 4})
// 	// problems.MergeTwoMaxHeaps([]int{10, 5, 6, 2}, []int{12, 7, 9}, 4, 3)
// 	// problems.HeightHeap([]int{10, 5, 6, 2}, 9)
// 	// problems.Search("ggggg", "gggggforggggg")
// 	// problems.KthLargeEle([]int{3, 2, 1, 5, 6, 4}, 2)
// 	// problems.MaxSubarraySum([]int{-2, 2, -5, 12, -11, -1, 7}, 3)
// 	// problems.FruitsInBasket([]int{1, 3, 1, 3})
// 	// problems.RemoveKdigits("1002661", 3)
// 	// problems.NextGreaterEle([]int{2, 5, 3, 9})
// 	// problems.PrevGreaterEle([]int{2, 5, 3, 9})
// 	// problems.NextLesserEle([]int{2, 5, 3, 9})
// 	// problems.PrevLesserEle([]int{2, 5, 3, 9})
// 	// problems.StackOpe()
// 	// fmt.Println(problems.CelebrityGraph([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 3))
// fmt.Println(problems.CelebrityStack([][]int{{0, 1, 0}, {1, 1, 0}, {1, 1, 0}}, 3))
// 	// fmt.Println(problems.Permutations([]int{1, 2, 1}))
// 	// fmt.Println(problems.SortStack([]int{1, 2, 1}))
// 	// fmt.Println(problems.DistSubSeq("fff"))
// 	// fmt.Println(problems.BinaryString(4))
// 	// fmt.Println(problems.GenerateParanthesis(4))
// 	// fmt.Println(problems.SubSeqWithSumK([]int{1, 2, 3, 1, 1, 1}, 3))
// 	// fmt.Println(problems.CombinationSum([]int{2, 3, 5}, 8))
// 	// fmt.Println(problems.CombinationSum3(3, 7))
// 	// fmt.Println(problems.Min_Spinklers([]int{-1, 0, 1, 2, 0, 4, 2, 4}, 8))
// 	// fmt.Println(problems.MergeNewIntervals([][]int{{1, 3}, {6, 9}}, []int{-1, 0}))
// 	// fmt.Println(problems.TopKFrequent([]int{5, 2, 1, 3, 2}, 4))
// 	// fmt.Println(problems.KTop([]int{5, 2, 1, 3, 2}, 5, 4))
// 	// fmt.Println(problems.AddOperators("123", 6))
// 	// result := problems.PrimePath(1033, 8179)
// 	// fmt.Println("Shortest path length:", result)
// 	result := problems.DistinctIslands([][]int{{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {1, 1, 0, 1, 1}})
// 	fmt.Println(result)
// 	// mp := map[string]interface{}{}
// 	// mp["str"] = func() int { return 3 }
// 	// fmt.Println(mp)
// 	// fmt.Println()
// 	// fmt.Println("Shortest path length:", result)
// 	// problems.BinarySearch()
// 	// m	ath.MaxInt32
// 	// chars := []string{"a", "b", "c"}
// 	// sort.Sort(sort.Reverse(sort.StringSlice(chars)))
// }

// }

func bfs() {
	var n int
	fmt.Println("Enter the number of nodes:")
	fmt.Scanln(&n)
	var m int
	fmt.Println("Enter the number of edges:")
	fmt.Scanln(&m)
	var adjList = map[int][]int{}
	for i := 0; i < m; i++ {
		var u int
		var v int
		fmt.Println("Enter the edge:")
		fmt.Scanln(&u)
		fmt.Scanln(&v)
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}
	var visited = map[int]bool{}
	var queue []int
	queue = append(queue, 2)
	var bfsList []int
	visited[2] = true
	for len(queue) != 0 {
		var node int = queue[0]
		queue = queue[1:]
		bfsList = append(bfsList, node)
		for _, neighbour := range adjList[node] {
			if !visited[neighbour] {
				queue = append(queue, neighbour)
				visited[neighbour] = true
			}
		}
	}
	fmt.Println("BFS traversal:")
	for i := 1; i <= n; i++ {
		fmt.Println(i, ":", bfsList[i-1])
	}
}

func findString(n, k int) string {
	ans := ""
	st := make(map[string]bool)

	for i := 0; i < n; i++ {
		ans += "0"
	}
	st[ans] = true

	var temp string

	for i := k - 1; i >= 0; i-- {
		temp = ans[len(ans)-n+1:] + strconv.Itoa(i)

		if _, exists := st[temp]; !exists {
			ans += strconv.Itoa(i)
			i = k
			st[temp] = true
		}
	}

	return ans
}

func linkedList() {
	fmt.Println("Select option from below:")
	fmt.Println("1. Create a linked list")
	fmt.Println("2. Insert a node at the beginning")
	fmt.Println("3. Insert a node at the end")
	fmt.Println("4. Insert a node at the specific position")
	fmt.Println("5. Insert a node after a specific node")
	fmt.Println("6. Delete a node from the beginning")
	fmt.Println("7. Delete a node from the end")
	fmt.Println("8. Delete a node from the specific position")
	fmt.Println("9. Delete a node after a specific node")
	fmt.Println("10. Display the linked list")
	fmt.Println("11. Length of the linked list")
	fmt.Println("12. Search a node in the linked list")
	fmt.Println("13. Search a node in the linked list using recursion")
	fmt.Println("14. Reverse a linked list")
	fmt.Print("15. Intersection of two linked lists")
	fmt.Println("16. Exit")
	for {
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Enter the data for the first node:")
			var data int64
			fmt.Scanln(&data)
			createLL(data)
		case 2:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			insertAtBeginning(data)
		case 3:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			insertAtEnd(data)
		case 4:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			fmt.Println("Enter the position for the new node:")
			var position int64
			fmt.Scanln(&position)
			insertAtPosition(data, position)
		case 5:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			fmt.Println("Enter the data of the node after which the new node is to be inserted:")
			var nodeData int64
			fmt.Scanln(&nodeData)
			insertAfterNode(data, nodeData)
		case 6:
			deleteFromBeginning()
		case 7:
			deleteFromEnd()
		case 8:
			fmt.Println("Enter the position of the node to be deleted:")
			var position int64
			fmt.Scanln(&position)
			deleteFromPosition(position)
		case 9:
			fmt.Println("Enter the data of the node after which the new node is to be deleted:")
			var nodeData int64
			fmt.Scanln(&nodeData)
			deleteAfterNode(nodeData)
		case 10:
			displayLL()
		case 11:
			fmt.Println("Length of the linked list is:", lengthLL())
		case 12:
			fmt.Println("Searching for a node in the linked list")
			fmt.Println("Enter the data of the node to be searched:")
			var nodeData int64
			fmt.Scanln(&nodeData)
			searchNode(nodeData)
		case 13:
			fmt.Println("search with recursion")
			fmt.Println("Enter the data of the node to be searched:")
			var nodeData int64
			fmt.Scanln(&nodeData)
			fmt.Println(searchNodeRecursion(head, nodeData))
		case 14:
			fmt.Println("reverse a linked list")
			reverseLL()
		case 15:
			fmt.Println("intersection of two linked lists")
			IOLinkedList(head, head1)
		case 16:
			return
		case 17:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
			fmt.Println("Select option from below:")
		}
		fmt.Println("Select option from below:")
	}
}

func createLL(data int64) {
	var node Node
	node.init(data)
	head = &node
	fmt.Println("Linked list created successfully with single node")
}

func insertAtBeginning(data int64) {
	var node Node
	node.init(data)
	node.next = head
	head = &node
}

func insertAtEnd(data int64) {
	var node Node
	node.init(data)
	if head == nil {
		head = &node
		return
	}

	var temp *Node
	temp = head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &node
}

func insertAtPosition(data int64, position int64) {
	var node Node

	node.init(data)
	if position == 1 {
		node.next = head
		head = &node
		return
	}
	if head == nil {
		head = &node
		return
	}

	var temp *Node
	temp = head
	var nodeCount int64 = 0
	for temp.next != nil && nodeCount < position-2 {
		temp = temp.next
		nodeCount++
	}
	node.next = temp.next
	temp.next = &node
	return
}

func insertAfterNode(data int64, nodeData int64) {
	var node Node

	node.init(data)
	if head == nil {
		fmt.Println("Linked list is empty")
	}

	var temp *Node
	temp = head
	for temp.next != nil && temp.data != nodeData {
		temp = temp.next
	}

	if temp.data != nodeData {
		fmt.Println("Node not found")
		return
	} else {
		node.next = temp.next
		temp.next = &node
	}

}

func deleteFromBeginning() {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println("Node deleted:", head.data)
	head = head.next
}

func deleteFromEnd() {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	var temp *Node
	temp = head
	for temp.next.next != nil {
		temp = temp.next
	}
	fmt.Println("Node deleted:", temp.next.data)
	temp.next = nil
}

func deleteFromPosition(position int64) {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	if position == 1 {
		head = head.next
		return
	}

	var temp *Node
	temp = head
	var nodeCount int64 = 0
	for temp.next != nil && nodeCount < position-2 {
		temp = temp.next
		nodeCount++
	}
	if temp.next == nil {
		fmt.Println("Position not found")
		return
	}
	fmt.Println("Node deleted:", temp.next.data)
	temp.next = temp.next.next
}

func deleteAfterNode(nodeData int64) {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	var temp *Node
	temp = head
	for temp.next != nil && temp.data != nodeData {
		temp = temp.next
		fmt.Println(temp.data)
	}
	if temp.next == nil {
		fmt.Println("Node not found")
		return
	}
	fmt.Println("Node deleted:", temp.next.data)
	temp.next = temp.next.next
}

func displayLL() {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	var temp *Node
	temp = head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func displayLL1(head3 *Node) {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	var temp *Node
	temp = head3
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func lengthLL() int64 {
	if head == nil {
		fmt.Println("Linked list is empty")
		return 0
	}
	var temp *Node
	temp = head
	var nodeCount int64 = 0
	for temp != nil {
		temp = temp.next
		nodeCount++
	}
	return nodeCount
}

func searchNode(data int64) {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	var temp *Node
	temp = head
	for temp.next != nil && temp.data != data {
		temp = temp.next
	}
	if temp.data == data {
		fmt.Println("Node with data", data, "found")
		return
	}
	fmt.Println("Node not found")
	return
}

// search using recursion
func searchNodeRecursion(head *Node, data int64) bool {
	if head == nil {
		return false
	}
	var temp *Node
	temp = head
	if temp.data == data {
		return true
	}
	temp = temp.next
	return searchNodeRecursion(temp, data)
}

func reverseLL() {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	var prev *Node
	var next *Node
	var current *Node
	current = head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	head = prev
}
func mergeResult(node1 *Node, node2 *Node) *Node {
	head := &Node{}
	temp := head
	temp1 := node1
	temp2 := node2
	for temp1 != nil && temp2 != nil {
		if temp1.data < temp2.data {
			temp.next = &Node{data: temp1.data}
			temp1 = temp1.next
		} else {
			temp.next = &Node{data: temp2.data}
			temp2 = temp2.next
		}
		temp = temp.next
	}
	for temp1 != nil {
		temp.next = &Node{data: temp1.data}
		temp1 = temp1.next
		temp = temp.next
	}
	for temp2 != nil {
		temp.next = &Node{data: temp2.data}
		temp2 = temp2.next
		temp = temp.next
	}
	head = head.next
	t := head.next
	for t.next != nil {
		fmt.Println(t.data)
		t = t.next
	}

	var prev *Node
	var next *Node
	var current *Node
	current = head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	head = prev
	return head
}

func IOLinkedList(head *Node, head1 *Node) {
	var arr map[int64]bool = make(map[int64]bool)

	var temp *Node
	temp = head
	for temp != nil {
		arr[temp.data] = false
		temp = temp.next
	}

	var temp1 *Node
	temp1 = head1
	var intersection *Node = nil
	var head2 *Node = nil
	for temp1 != nil {
		for i := 0; i < len(arr); i++ {
			if arr[temp1.data] == false {
				var node Node
				node.init(temp1.data)
				if intersection == nil {
					intersection = &node
					head2 = intersection
				} else {
					intersection.next = &node
					intersection = intersection.next
				}
				arr[temp1.data] = true
			}
		}
		temp1 = temp1.next
	}
	displayLL1(head2)
}

func stack() {
	fmt.Println("Select option from below:")
	fmt.Println("1. Push")
	fmt.Println("2. Pop")
	fmt.Println("3. Display")
	fmt.Println("4. Exit")
	for {
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			pushArray(data)
		case 2:
			popArray()
		case 3:
			displayStackArray()
		case 4:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			pushLinkedList(data)
		case 5:
			popLinkedList()
		case 6:
			displayStackLinkedList()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
			fmt.Println("Select option from below:")
		}
		fmt.Println("Select option from below:")
	}
}

var topArray = -1
var topLL *Node = nil
var stackArray [9]int64

func pushArray(data int64) {
	fmt.Println("pushing data into array")
	if topArray == len(stackArray)-1 {
		fmt.Println("Stack overflow")
		return
	}
	topArray++
	stackArray[topArray] = data
	return
}

func popArray() {
	fmt.Println("popping data from array")
	if topArray == -1 {
		fmt.Println("Stack underflow")
		return
	}

	fmt.Println("Data popped:", stackArray[topArray])
	topArray--
	return
}

func displayStackArray() {
	fmt.Println("Displaying stack")
	if topArray == -1 {
		fmt.Println("Stack is empty")
		return
	}
	for i := topArray; i >= 0; i-- {
		fmt.Println(stackArray[i])
	}
}

func pushLinkedList(data int64) {
	fmt.Println("pushing data into linked list")
	var node Node
	node.init(data)
	node.next = topLL
	topLL = &node
}

func popLinkedList() {
	fmt.Println("popping data from linked list")
	if topLL == nil {
		fmt.Println("Stack underflow")
		return
	}
	fmt.Println("Data popped:", topLL.data)
	topLL = topLL.next
}

func displayStackLinkedList() {
	fmt.Println("Displaying stack")
	if topLL == nil {
		fmt.Println("Stack is empty")
		return
	}
	var temp *Node
	temp = topLL
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func queue() {
	fmt.Println("Select option from below:")
	fmt.Println("1. Insert")
	fmt.Println("2. Delete")
	fmt.Println("3. Display")
	fmt.Println("4. Exit")
	for {
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			insertArrayQ(data)
		case 2:
			deleteArrayQ()
		case 3:
			displayArrayQ()
		case 4:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			insertLinkedListQ(data)
		case 5:
			deleteLinkedListQ()
		case 6:
			displayLinkedListQ()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
			fmt.Println("Select option from below:")
		}
		fmt.Println("Select option from below:")
	}
}

var arrayQ [9]int64
var frontArrayQ = -1
var rearArrayQ = -1

func insertArrayQ(data int64) {
	fmt.Println("inserting data into array")
	if rearArrayQ == len(arrayQ)-1 {
		fmt.Println("Queue overflow")
		return
	}

	if frontArrayQ == -1 {
		frontArrayQ = 0
	}
	rearArrayQ++
	arrayQ[rearArrayQ] = data
	return
}

func deleteArrayQ() {
	fmt.Println("deleting data from array")
	if frontArrayQ == -1 {
		fmt.Println("Queue underflow")
		return
	}

	fmt.Println("Data deleted:", arrayQ[frontArrayQ])
	if frontArrayQ == rearArrayQ {
		frontArrayQ = -1
		rearArrayQ = -1
		return
	}
	frontArrayQ++
	return
}

func displayArrayQ() {
	fmt.Println("Displaying queue")
	if frontArrayQ == -1 {
		fmt.Println("Queue is empty")
		return
	}
	for i := frontArrayQ; i <= rearArrayQ; i++ {
		fmt.Println(arrayQ[i])
	}
}

var frontLLQ *Node = nil
var rearLLQ *Node = nil

func insertLinkedListQ(data int64) {
	fmt.Println("inserting data into linked list")
	var node Node
	node.init(data)
	if rearLLQ == nil {
		frontLLQ = &node
		rearLLQ = &node
		return
	}
	rearLLQ.next = &node
	rearLLQ = &node
}

func deleteLinkedListQ() {
	fmt.Println("deleting data from linked list")
	if frontLLQ == nil {
		fmt.Println("Queue underflow")
		return
	}
	fmt.Println("Data deleted:", frontLLQ.data)
	if frontLLQ == rearLLQ {
		frontLLQ = nil
		rearLLQ = nil
		return
	}
	frontLLQ = frontLLQ.next
}

func displayLinkedListQ() {
	fmt.Println("Displaying queue")
	if frontLLQ == nil {
		fmt.Println("Queue is empty")
		return
	}
	var temp *Node
	temp = frontLLQ
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func binaryTree() {
	fmt.Println("Select option from below:")
	fmt.Println("1. Create a binary tree")
	fmt.Println("2. Insert a node")
	fmt.Println("3. Delete a node")
	// fmt.Println("4. Display the binary tree")
	fmt.Println("4. Preorder traversal")
	fmt.Println("5. Inorder traversal")
	fmt.Println("6. Postorder traversal")
	fmt.Println("7. Exit")
	for {
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Enter the data for the root node:")
			var data int64
			fmt.Scanln(&data)
			createBT(data)
		case 2:
			fmt.Println("Enter the data for the new node:")
			var data int64
			fmt.Scanln(&data)
			insertBT(data)
		case 3:
			fmt.Println("Enter the data for the node to be deleted:")
			var data int64
			fmt.Scanln(&data)
			deleteBT(data)
		case 4:
			fmt.Println("Preorder traversal")
			preorderTraversal(root)
		case 5:
			fmt.Println("Inorder traversal")
			inorderTraversal(root)
		case 6:
			fmt.Println("Postorder traversal")
			postorderTraversal(root)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
			fmt.Println("Select option from below:")
		}
		fmt.Println("Select option from below:")
	}
}

var root *TNode = nil

func createBT(data int64) {
	var node TNode
	node.init(data)
	root = &node
	fmt.Println("Binary tree created successfully with root node")
}

func insertBT(data int64) {
	var node TNode
	node.init(data)
	if root == nil {
		root = &node
		return
	}

	var temp *TNode
	temp = root
	var parent *TNode
	for temp != nil {
		parent = temp
		if data < temp.data {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	if data < parent.data {
		parent.left = &node
	} else {
		parent.right = &node
	}
}

func deleteBT(data int64) {
	var temp *TNode
	temp = root
	var parent *TNode

	//searching for the node to be deleted
	for temp != nil && temp.data != data {
		parent = temp
		if data < temp.data {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	//if node not found
	if temp == nil {
		fmt.Println("Node not found")
		return
	}

	//if node is a leaf node
	//then delete the node by making the parent's left or right pointer nil whichever points to the node
	if temp.left == nil && temp.right == nil {
		if temp == root {
			root = nil
			return
		}
		if parent.left == temp {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}

	//if node has only one child
	//then delete the node by making the parent's left or right pointer point to the node's child
	if temp.left == nil || temp.right == nil {
		//if node is the root node
		//then make the root point to the node's child
		if temp == root {
			if temp.left == nil {
				root = temp.right
			} else {
				root = temp.left
			}
			return
		}
		//if node is not the root node
		//then make the parent's left or right pointer point to the node's child
		//whichever points to the node
		//if node is the left child of the parent then make the parent's left pointer point to the node's child
		//if node is the right child of the parent then make the parent's right pointer point to the node's child
		//if node has left child then make the parent's left pointer point to the node's left child
		//if node has right child then make the parent's left pointer point to the node's right child
		if parent.left == temp {
			if temp.left == nil {
				parent.left = temp.right
			} else {
				parent.left = temp.left
			}
		} else {
			if temp.left == nil {
				parent.right = temp.right
			} else {
				parent.right = temp.left
			}
		}
		return
	}

	//if node has two children
	//then find the inorder successor of the node
	//replace the node's data with the inorder successor's data
	//delete the inorder successor
	if temp.left != nil && temp.right != nil {
		var successor *TNode
		successor = temp.right
		for successor.left != nil {
			successor = successor.left
		}
		var data int64 = successor.data
		deleteBT(successor.data)
		temp.data = data
	}
}

func preorderTraversal(root *TNode) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	preorderTraversal(root.left)
	preorderTraversal(root.right)
}

func inorderTraversal(root *TNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.left)
	fmt.Println(root.data)
	inorderTraversal(root.right)
}

func postorderTraversal(root *TNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.left)
	postorderTraversal(root.right)
	fmt.Println(root.data)
}
func getMinDiff(arr []int, n int, k int) {
	sort.IntSlice.Sort(arr)
	ans := arr[n-1] - arr[0]
	for i := 0; i < n; i++ {
		if arr[i] < k {
			continue
		}
		tempMax := int(math.Max(float64(arr[i-1]+k), float64(arr[n-1]-k)))
		teampMin := int(math.Min(float64(arr[0]+k), float64(arr[i]-k)))
		ans = int(math.Min(float64(ans), math.Abs(float64(tempMax-teampMin))))
	}
	fmt.Println(ans)
}
