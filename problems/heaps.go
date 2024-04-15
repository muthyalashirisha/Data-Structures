package problems

import (
	"container/heap"
	"fmt"
	"math"
)

// represents max heap based on priority queue
type MaxPriorityQueue []int

func (pq MaxPriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq MaxPriorityQueue) Len() int {
	return len(pq)
}

func (pq MaxPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxPriorityQueue) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *MaxPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func PriorityHeap(arr []int) {

	pq := &MaxPriorityQueue{}
	for _, v := range arr {
		*pq = append(*pq, v)
	}

	heap.Init(pq)
	heap.Push(pq, 23)
	//pop
	item := heap.Pop(pq)
	fmt.Println(item)
}

// MaxHeap represents a max-heap.
type MaxHeap struct {
	array []int
}

// NewMaxHeap initializes a new empty max-heap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Push adds an element to the max-heap.
func (h *MaxHeap) Push(x int) {
	h.array = append(h.array, x)
	h.heapifyUp()
}

// Pop removes and returns the maximum element from the max-heap.
func (h *MaxHeap) Pop() (int, error) {
	if h.isEmpty() {
		return 0, fmt.Errorf("heap is empty")
	}

	max := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.heapifyDown()

	return max, nil
}

// heapifyUp maintains the max-heap property by moving the last element up the heap.
func (h *MaxHeap) heapifyUp() {
	currentIndex := len(h.array) - 1

	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2

		if h.array[currentIndex] <= h.array[parentIndex] {
			break // Heap property is satisfied
		}

		// Swap the current element with its parent
		h.array[currentIndex], h.array[parentIndex] = h.array[parentIndex], h.array[currentIndex]

		// Move up the heap
		currentIndex = parentIndex
	}
}

// heapifyDown maintains the max-heap property by moving the root element down the heap.
func (h *MaxHeap) heapifyDown() {
	currentIndex := 0
	size := len(h.array)

	for {
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2

		largestIndex := currentIndex

		// Find the index of the largest element among the current node and its children
		if leftChildIndex < size && h.array[leftChildIndex] > h.array[largestIndex] {
			largestIndex = leftChildIndex
		}
		if rightChildIndex < size && h.array[rightChildIndex] > h.array[largestIndex] {
			largestIndex = rightChildIndex
		}

		// If the largest element is the current node, heap property is satisfied
		if largestIndex == currentIndex {
			break
		}

		// Swap the current element with the largest child
		h.array[currentIndex], h.array[largestIndex] = h.array[largestIndex], h.array[currentIndex]

		// Move down the heap
		currentIndex = largestIndex
	}
}

// isEmpty checks if the max-heap is empty.
func (h *MaxHeap) isEmpty() bool {
	return len(h.array) == 0
}

func Heap() {
	// Example usage
	maxHeap := NewMaxHeap()

	// Push elements into the max heap
	maxHeap.Push(3)
	maxHeap.Push(1)
	maxHeap.Push(4)
	maxHeap.Push(1)
	maxHeap.Push(5)
	maxHeap.Push(9)

	// Pop elements from the max heap
	for !maxHeap.isEmpty() {
		max, err := maxHeap.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Popped:", max)
	}
}

type MinHeap struct {
	array []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Push(x int) {
	h.array = append(h.array, x)
	h.heapifyUp()
}

func (h *MinHeap) pop() (int, error) {
	if h.isEmpty() {
		return 0, fmt.Errorf("heap is empty")
	}

	min := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.heapifyDown()
	return min, nil
}

func (h *MinHeap) isEmpty() bool {
	return len(h.array) == 0
}

func (h *MinHeap) heapifyUp() {
	currIndex := len(h.array) - 1
	for currIndex > 0 {
		parentIndex := (currIndex - 1) / 2

		if h.array[currIndex] >= h.array[parentIndex] {
			break
		}
		h.array[currIndex], h.array[parentIndex] = h.array[parentIndex], h.array[currIndex]
		currIndex = parentIndex
	}
}

func (h *MinHeap) heapifyDown() {
	currentIndex := 0
	size := len(h.array)

	for {
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2

		smallestIndex := currentIndex
		if leftChildIndex < size && h.array[smallestIndex] > h.array[leftChildIndex] {
			smallestIndex = leftChildIndex
		}

		if rightChildIndex < size && h.array[smallestIndex] > h.array[rightChildIndex] {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == currentIndex {
			break
		}

		h.array[smallestIndex], h.array[currentIndex] = h.array[currentIndex], h.array[smallestIndex]

		currentIndex = smallestIndex
	}
}

func isMinHeap(arr []int) bool {
	n := len(arr)

	for i := 0; i <= (n-2)/2; i++ {
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild < n && arr[leftchild] < arr[i] {
			return false
		}
		if rightchild < n && arr[rightchild] < arr[i] {
			return false
		}
	}
	return true
}
func HeightHeap(arr []int, n int) {
	h := int(math.Floor(math.Log2(float64(n + 1))))
	fmt.Println(h)
}
func isMaxHeap(arr []int) bool {
	n := len(arr)

	for i := 0; i <= (n-2)/2; i++ {
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild < n && arr[leftchild] > arr[i] {
			return false
		}
		if rightchild < n && arr[rightchild] > arr[i] {
			return false
		}
	}
	return true
}

func ConvertMintoMax(arr []int) {
	n := len(arr)
	// reverseArray(arr, n)
	fmt.Println(arr)
	//heapify
	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(arr, n, i)
	}
	fmt.Println(arr)
}

func reverseArray(arr []int, n int) {
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func heapifyDown(arr []int, n int, currIndex int) {

	for {
		lIndex := 2*currIndex + 1
		rIndex := 2*currIndex + 2

		largest := currIndex
		if lIndex < n && arr[lIndex] > arr[largest] {
			largest = lIndex
		}
		if rIndex < n && arr[rIndex] > arr[largest] {
			largest = rIndex
		}

		if currIndex == largest {
			break
		}

		arr[largest], arr[currIndex] = arr[currIndex], arr[largest]

		currIndex = largest
	}
}

func MergeTwoMaxHeaps(a []int, b []int, n, m int) []int {
	if n > m {
		return MergeTwoMaxHeaps(b, a, m, n)
	}
	for i := 0; i < n; i++ {
		b = push(a[i], b)
	}
	fmt.Println(b)
	return b
}

func MergeTwoMaxHeaps1(a []int, b []int, n, m int) []int {
	arr := append(a, b...)
	for i := (n+m-1)/2 - 1; i >= 0; i-- {
		heapify(arr, i, n+m)
	}
	fmt.Println(b)
	return b
}

func heapify(arr []int, currIndex int, n int) {
	for {
		lChild := 2*currIndex + 1
		rChild := 2*currIndex + 2

		largest := currIndex
		if lChild < n && arr[largest] < arr[lChild] {
			largest = lChild
		}
		if rChild < n && arr[largest] < arr[rChild] {
			largest = rChild
		}
		if largest == currIndex {
			break
		}
		arr[largest], arr[currIndex] = arr[currIndex], arr[largest]
		currIndex = largest
	}
}

func push(ele int, b []int) []int {
	res := append(b, ele)
	currIndex := len(b) - 1

	for currIndex > 0 {
		parentIndex := (currIndex - 1) / 2

		if res[parentIndex] >= res[currIndex] {
			return res
		}

		res[parentIndex], res[currIndex] = res[currIndex], res[parentIndex]

		currIndex = parentIndex
	}
	return res
}

type MinHeap1 struct {
	harr     []int
	capacity int
	heapSize int
}

func NewMinHeap1(capacity int) *MinHeap1 {
	return &MinHeap1{
		harr:     make([]int, capacity),
		capacity: capacity,
		heapSize: 0,
	}
}

func (minHeap *MinHeap1) parent(i int) int {
	return (i - 1) / 2
}

func (minHeap *MinHeap1) left(i int) int {
	return 2*i + 1
}

func (minHeap *MinHeap1) right(i int) int {
	return 2*i + 2
}

func (minHeap *MinHeap1) extractMin() int {
	if minHeap.heapSize <= 0 {
		return -1 // or some sentinel value indicating an empty heap
	}
	if minHeap.heapSize == 1 {
		minHeap.heapSize--
		return minHeap.harr[0]
	}

	root := minHeap.harr[0]
	minHeap.harr[0] = minHeap.harr[minHeap.heapSize-1]
	minHeap.heapSize--
	minHeap.MinHeapify(0)

	return root
}

func (minHeap *MinHeap1) insertKey(k int) {
	if minHeap.heapSize == minHeap.capacity {
		fmt.Println("Overflow: Could not insert key.")
		return
	}

	minHeap.heapSize++
	i := minHeap.heapSize - 1
	minHeap.harr[i] = k

	for i != 0 && minHeap.harr[minHeap.parent(i)] > minHeap.harr[i] {
		minHeap.harr[i], minHeap.harr[minHeap.parent(i)] = minHeap.harr[minHeap.parent(i)], minHeap.harr[i]
		i = minHeap.parent(i)
	}
}

func (minHeap *MinHeap1) deleteKey(i int) {
	if i < 0 || i >= minHeap.heapSize {
		fmt.Println("Index out of range: Could not delete key.")
		return
	}

	minHeap.decreaseKey(i, -1<<31) // Set the value to the smallest integer
	minHeap.extractMin()
}

func (minHeap *MinHeap1) decreaseKey(i, new_val int) {
	minHeap.harr[i] = new_val
	for i != 0 && minHeap.harr[minHeap.parent(i)] > minHeap.harr[i] {
		minHeap.harr[i], minHeap.harr[minHeap.parent(i)] = minHeap.harr[minHeap.parent(i)], minHeap.harr[i]
		i = minHeap.parent(i)
	}
}

func (minHeap *MinHeap1) MinHeapify(i int) {
	l := minHeap.left(i)
	r := minHeap.right(i)
	smallest := i

	if l < minHeap.heapSize && minHeap.harr[l] < minHeap.harr[i] {
		smallest = l
	}
	if r < minHeap.heapSize && minHeap.harr[r] < minHeap.harr[smallest] {
		smallest = r
	}
	if smallest != i {
		minHeap.harr[i], minHeap.harr[smallest] = minHeap.harr[smallest], minHeap.harr[i]
		minHeap.MinHeapify(smallest)
	}
}
