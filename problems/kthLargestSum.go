package problems

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func KthLargest(arr []int, n int, k int) {
	list := []int{}
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			list = append(list, sum)
		}
	}

	sort.Ints(list)
	fmt.Println(list, list[len(list)-k])
}

// KthLargestHeap is a min-heap implementation for Kth largest elements
type KthLargestHeap []int

func (h KthLargestHeap) Len() int           { return len(h) }
func (h KthLargestHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h KthLargestHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push pushes an element onto the heap
func (h *KthLargestHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop pops the minimum element from the heap
func (h *KthLargestHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// kthLargestSum calculates the Kth largest sum of contiguous subarrays
func KthLargestSum(arr []int, N, K int) {
	// Array to store prefix sums
	sum := make([]int, N+1)
	sum[0] = 0
	sum[1] = arr[0]

	for i := 2; i <= N; i++ {
		sum[i] = sum[i-1] + arr[i-1]
	}
	fmt.Println(sum)
	// Min heap
	var Q KthLargestHeap

	heap.Init(&Q)

	// Loop to calculate the contiguous subarray sum position-wise
	for i := 1; i <= N; i++ {
		// Loop to traverse all positions that form contiguous subarray
		for j := i; j <= N; j++ {
			// Calculates the contiguous subarray sum from j to i index
			x := sum[j] - sum[i-1]

			// If heap has less than K elements, then simply push it
			if Q.Len() < K {
				heap.Push(&Q, x)
			} else {
				// If the min heap has equal to K elements, then check
				// if the smallest Kth element is smaller than x, then insert
				if Q[0] < x {
					heap.Pop(&Q)
					heap.Push(&Q, x)
				}
			}
		}
	}

	fmt.Println(heap.Pop(&Q).(int))
	// The top element will be the Kth largest element
	// return heap.Pop(&Q).(int)
}

func KthLargestSum2(arr []int, N, K int) {
	// Array to store prefix sums
	presum := make([]int, N+1)
	presum[0] = 0
	presum[1] = arr[0]

	for i := 2; i <= N; i++ {
		presum[i] = presum[i-1] + arr[i-1]
	}
	fmt.Println(presum)

	sum := []int{}
	// Loop to calculate the contiguous subarray sum position-wise
	for i := 1; i <= N; i++ {
		// Loop to traverse all positions that form contiguous subarray
		for j := i; j <= N; j++ {
			// Calculates the contiguous subarray sum from j to i index
			sum = append(sum, presum[j]-presum[i-1])
		}
	}
	sort.Ints(sum)
	fmt.Println(sum[len(sum)-K])
}

func MaxSumWOAdj(arr []int, n int) {
	//memoization
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	fmt.Println(sum(n-1, arr, dp))

	//tabulation
	dp1 := make([]int, n)
	dp1[0] = arr[0]
	dp1[1] = int(math.Max(float64(arr[0]), float64(arr[1])))
	for i := 2; i < n; i++ {
		not_pick := dp1[i-1]
		pick := arr[i] + dp1[i-2]
		dp1[i] = int(math.Max(float64(pick), float64(not_pick)))
	}
	fmt.Println(dp1, dp[n-1])
}

func sum(i int, arr []int, dp []int) int {
	if i == 0 {
		return arr[0]
	}
	if i < 0 {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}
	not_pick := sum(i-1, arr, dp)
	pick := arr[i] + sum(i-2, arr, dp)
	dp[i] = int(math.Max(float64(pick), float64(not_pick)))
	return dp[i]
}

type KthLargestEle []int

func (k KthLargestEle) Less(i, j int) bool {
	return k[i] > k[j]
}
func (k KthLargestEle) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
func (k KthLargestEle) Len() int {
	return len(k)
}
func (k *KthLargestEle) Push(x interface{}) {
	*k = append(*k, x.(int))
}
func (k *KthLargestEle) Pop() interface{} {
	old := *k
	n := len(old)
	x := old[n-1]
	*k = old[:n-1]
	return x
}

// kth largest element
func KthLargeEle(arr []int, k int) int {
	if len(arr) < k {
		return -1
	}
	var PQ KthLargestEle
	heap.Init(&PQ)
	for i := 0; i < len(arr); i++ {
		heap.Push(&PQ, arr[i])
	}
	for i := 1; i < k; i++ {
		heap.Pop(&PQ)
	}
	fmt.Println(PQ[0])
	return PQ[0]
}

type KthSmallestEle []int

func (k KthSmallestEle) Less(i, j int) bool {
	return k[i] < k[j]
}
func (k KthSmallestEle) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
func (k KthSmallestEle) Len() int {
	return len(k)
}
func (k *KthSmallestEle) Push(x interface{}) {
	*k = append(*k, x.(int))
}
func (k *KthSmallestEle) Pop() interface{} {
	old := *k
	n := len(old)
	x := old[n-1]
	*k = old[:n-1]
	return x
}
func KSortedArray(arr []int, k int) []int {
	n := len(arr)
	var pq KthSmallestEle
	heap.Init(&pq)
	min := int(math.Min(float64(n), float64(k)))
	for i := 0; i <= min; i++ {
		heap.Push(&pq, arr[i])
	}

	ans := []int{}
	for i := k + 1; i < n; i++ {
		ans = append(ans, pq[0])
		heap.Pop(&pq)
		heap.Push(&pq, arr[i])
		fmt.Println(ans)
	}
	for len(pq) > 0 {
		ans = append(ans, heap.Pop(&pq).(int))
	}
	fmt.Println(ans)
	return ans
}

func MergeKLinkedLists(lists []*Node1) *Node1 {
	var pq KthSmallestEle
	heap.Init(&pq)
	for i := 0; i < len(lists); i++ {
		temp := lists[i]
		for temp != nil {
			heap.Push(&pq, temp.data)
			temp = temp.next
		}
	}
	head := &Node1{}
	temp := head
	for len(pq) > 0 {
		val := heap.Pop(&pq).(int)
		temp.next = &Node1{data: int64(val)}
		temp = temp.next
	}
	return head.next
}

func TaskScheduler(tasks []byte, n int) int {
	taskFreq := make([]int, 26)

	for _, task := range tasks {
		taskFreq[task-'A']++
	}

	pq := make(KthLargestEle, 0)
	heap.Init(&pq)

	for _, freq := range taskFreq {
		if freq > 0 {
			heap.Push(&pq, freq)
		}
	}

	time := 0

	for len(pq) > 0 {
		var remain []int
		cycle := n + 1

		for cycle > 0 && len(pq) > 0 {
			maxFreq := heap.Pop(&pq).(int)
			if maxFreq > 1 {
				remain = append(remain, maxFreq-1)
			}
			time++
			cycle--
		}

		for _, count := range remain {
			heap.Push(&pq, count)
		}

		if len(pq) == 0 {
			break
		}

		time += cycle
	}

	return time

}
func isNStraightHand(hand []int, groupSize int) bool {
	// mp := make(map[int]int)
	// for _, x := range hand {
	// 	mp[x]++
	// }

	// minHeap := &MinHeap2{}
	// heap.Init(minHeap)

	// for u := range mp {
	// 	heap.Push(minHeap, u)
	// }

	// for len(*minHeap) > 0 {
	// 	var temp []int
	// 	for i := 0; i < groupSize; i++ {
	// 		if minHeap.Len() > 0 {
	// 			temp = append(temp, heap.Pop(minHeap).(int))
	// 		}
	// 	}

	// 	for _, val := range temp {
	// 		if mp[val]--; mp[val] > 0 {
	// 			heap.Push(minHeap, val)
	// 		}
	// 	}

	// 	if len(temp) == groupSize {
	// 		if !consecutive(temp) {
	// 			return false
	// 		}
	// 	} else {
	// 		return false
	// 	}
	// }

	// return true
	minHeap := &MinHeap2{}
	heap.Init(minHeap)
	mp := make(map[int]int)
	for _, x := range hand {
		if mp[x] == 0 {
			heap.Push(minHeap, x)
		}
		mp[x]++
	}

	for len(*minHeap) > 0 {
		var temp []int
		for i := 0; i < groupSize; i++ {
			if minHeap.Len() > 0 {
				curr := heap.Pop(minHeap).(int)
				if len(temp) > 0 && curr != temp[len(temp)-1]+1 {
					return false
				}
				temp = append(temp, curr)
			} else {
				return false
			}
		}

		for _, val := range temp {
			if mp[val]--; mp[val] > 0 {
				heap.Push(minHeap, val)
			}
		}
	}

	return true
}

func consecutive(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1]+1 {
			return false
		}
	}
	return true
}

type MinHeap2 []int

func (k MinHeap2) Less(i, j int) bool {
	return k[i] < k[j]
}
func (k MinHeap2) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
func (k MinHeap2) Len() int {
	return len(k)
}
func (k *MinHeap2) Push(x interface{}) {
	*k = append(*k, x.(int))
}
func (k *MinHeap2) Pop() interface{} {
	old := *k
	n := len(old)
	x := old[n-1]
	*k = old[:n-1]
	return x
}

type Twitter struct {
	Tweets []Tweet
}

type Tweet struct {
	UserId    int
	TweetId   []int
	Follower  []int
	Following []int
}

func Constructor() Twitter {
	return Twitter{}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if len(this.Tweets) > 0 {
		tweets := this.Tweets
		found := true
		for _, v := range tweets {
			if v.UserId == userId {
				v.TweetId = append(v.TweetId, tweetId)
				found = true
			}
		}
		if !found {
			tweet := Tweet{
				UserId:  userId,
				TweetId: []int{tweetId},
			}
			tweets = append(tweets, tweet)
		}
	} else {
		tweet := Tweet{
			UserId:  userId,
			TweetId: []int{tweetId},
		}
		this.Tweets = append(this.Tweets, tweet)
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	if len(this.Tweets) > 0 {
		tweets := this.Tweets
		found := true
		tweetIds := []int{}
		following := []int{}
		for _, v := range tweets {
			if v.UserId == userId {
				tweetIds = append(tweetIds, v.TweetId...)
				following = append(following, v.Following...)
			}
		}
		for _, v := range tweets {
			if v.UserId == userId {
				tweetIds = append(tweetIds, v.TweetId...)
				following = append(following, v.Following...)
			}
		}
		if !found {
			return nil
		}
	} else {
		return nil
	}
	return nil
}

func (this *Twitter) Follow(followerId int, followeeId int) {

}

func (this *Twitter) Unfollow(followerId int, followeeId int) {

}

type KthLargestN struct {
	pq KthLargestEle
	k  int
}

func Construct(k int, nums []int) KthLargestN {
	kth := KthLargestN{
		pq: []int{},
		k:  k,
	}
	heap.Init(&kth.pq)
	for i := 0; i < len(nums); i++ {
		heap.Push(&kth.pq, nums[i])
	}
	return kth
}

func (this *KthLargestN) Add(val int) int {
	heap.Push(&this.pq, val)
	k := this.k
	ans := []int{}
	for i := 0; i < k-1; i++ {
		res := heap.Pop(&this.pq).(int)
		ans = append(ans, res)
	}
	kthEle := heap.Pop(&this.pq).(int)
	for i := 0; i < k-1; i++ {
		heap.Push(&this.pq, ans[i])
	}
	return kthEle
}

type KthLargestPQ struct {
	K    int
	Heap *MinHeapk
}

type MinHeapk []int

func (h MinHeapk) Len() int {
	return len(h)
}
func (h MinHeapk) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeapk) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeapk) Push(i any) {
	*h = append(*h, i.(int))
}
func (h *MinHeapk) Pop() any {
	old := *h
	length := h.Len()
	x := old[length-1]
	*h = old[0 : length-1]
	return x
}
func Constructor1(k int, nums []int) KthLargestPQ {
	h := &MinHeapk{}
	heap.Init(h)

	obj := KthLargestPQ{
		K:    k,
		Heap: h,
	}
	for _, num := range nums {
		obj.Add(num)
	}
	return obj
}

func (this *KthLargestPQ) Add(val int) int {
	heap.Push(this.Heap, val)
	if this.Heap.Len() > this.K {
		heap.Pop(this.Heap)
	}
	return (*this.Heap)[0]
}
