package problems

import (
	"fmt"
)

func mergeResult(node1 *Node1, node2 *Node1) *Node1 {
	head := &Node1{}
	//merge lls
	temp := head
	temp1 := node1
	temp2 := node2
	for temp1 != nil && temp2 != nil {
		if temp1.data < temp2.data {
			temp.next = &Node1{data: temp1.data}
			temp1 = temp1.next
		} else {
			temp.next = &Node1{data: temp2.data}
			temp2 = temp2.next
		}
		temp = temp.next
	}
	for temp1 != nil {
		temp.next = &Node1{data: temp1.data}
		temp1 = temp1.next
		temp = temp.next
	}
	for temp2 != nil {
		temp.next = &Node1{data: temp2.data}
		temp2 = temp2.next
		temp = temp.next
	}
	head = head.next
	t := head.next
	for t.next != nil {
		fmt.Println(t.data)
		t = t.next
	}

	//reverse ll
	var prev *Node1
	var next *Node1
	var current *Node1
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
