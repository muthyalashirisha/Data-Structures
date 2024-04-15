package problems

import "fmt"

type Node1 struct {
	data int64
	next *Node1
}

func detectLoop(node *Node1) bool {
	slow := node
	fast := node
	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		} else {
			slow = slow.next
			fast = fast.next.next
		}
	}
	return false
}
func leftView(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	queue := []Node{}
	queue = append(queue, *root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if i == 0 {
				ans = append(ans, curr.key)
			}
			if curr.left != nil {
				queue = append(queue, *curr.left)
			}
			if curr.right != nil {
				queue = append(queue, *curr.right)
			}
		}
	}
	fmt.Println(ans)

	res := []int{}
	recursiveLeftView(root, 0, res)
	fmt.Println(res)
	return ans
}

func recursiveLeftView(node *Node, level int, ans []int) {
	if node == nil {
		return
	}
	if len(ans) == level {
		ans = append(ans, node.key)
	}
	recursiveLeftView(node.left, level+1, ans)
	recursiveLeftView(node.right, level+1, ans)
}
