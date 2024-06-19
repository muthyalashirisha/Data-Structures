package problems

import (
	"container/heap"
)

// HuffmanTreeNode represents a node in the Huffman tree
type HuffmanTreeNode struct {
	char      rune
	frequency int
	left      *HuffmanTreeNode
	right     *HuffmanTreeNode
}

// HuffmanHeap is a priority queue that implements heap.Interface
type HuffmanHeap []*HuffmanTreeNode

func (h HuffmanHeap) Len() int           { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h HuffmanHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanTreeNode))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// BuildHuffmanTree builds the Huffman tree for given frequencies
func BuildHuffmanTree(freq map[rune]int) *HuffmanTreeNode {
	h := &HuffmanHeap{}
	heap.Init(h)

	// Create a leaf node for each character and add it to the priority queue
	for char, frequency := range freq {
		heap.Push(h, &HuffmanTreeNode{char: char, frequency: frequency})
	}

	// Iterate until the heap contains only one node
	for h.Len() > 1 {
		// Extract the two nodes of highest priority (lowest frequency)
		left := heap.Pop(h).(*HuffmanTreeNode)
		right := heap.Pop(h).(*HuffmanTreeNode)

		// Create a new internal node with these two nodes as children and
		// with frequency equal to the sum of the two nodes' frequencies
		merged := &HuffmanTreeNode{
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}

		// Add the new node to the heap
		heap.Push(h, merged)
	}

	// The remaining node is the root of the Huffman tree
	return heap.Pop(h).(*HuffmanTreeNode)
}

// GenerateHuffmanCodes generates the Huffman codes for each character
func GenerateHuffmanCodes(node *HuffmanTreeNode, prefix string, codes map[rune]string) {
	if node == nil {
		return
	}

	// If this is a leaf node, then it contains one of the input characters
	if node.left == nil && node.right == nil {
		codes[node.char] = prefix
	}

	// Traverse the left and right subtrees
	GenerateHuffmanCodes(node.left, prefix+"0", codes)
	GenerateHuffmanCodes(node.right, prefix+"1", codes)
}
