package problems

import (
	"container/list"
	"fmt"
)

func ListsDemo() {
	var newList list.List
	newList.PushBack(123)
	newList.PushBack(456)
	newList.PushBack(789)

	for element := newList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// var list1 list.List

}

type MaxHeap1 struct {
	arr []int
}

func (hp *MaxHeap1) Insert(key int) {
	hp.arr = append(hp.arr, key)
	hp.heapifyMaxUp(len(hp.arr) - 1)
}
func (hp *MaxHeap1) heapifyMaxUp(index int) {
	for hp.arr[parent(index)] < hp.arr[index] {
		hp.swap(parent(index), index)
		index = parent(index)
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*i + 2
}
func right(i int) int {
	return 2*i + 2
}
func (hp *MaxHeap1) swap(int1, int2 int) {
	hp.arr[int1], hp.arr[int2] = hp.arr[int2], hp.arr[int1]
}

func heaps() {
	m := &MaxHeap1{}
	fmt.Println(m)

	makeHeap := []int{1, 2, 3}
	for _, v := range makeHeap {
		m.Insert(v)
	}
}
