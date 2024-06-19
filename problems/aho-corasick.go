package problems

type TrieNode1 struct {
	children map[rune]*TrieNode1
	fail     *TrieNode1
	output   []string
}

type AhoCorasick struct {
	root *TrieNode1
}

func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root: &TrieNode1{
			children: make(map[rune]*TrieNode1),
			output:   []string{},
		},
	}
}

func (ac *AhoCorasick) Add(pattern string) {
	current := ac.root
	for _, char := range pattern {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &TrieNode1{
				children: make(map[rune]*TrieNode1),
				output:   []string{},
			}
		}
		current = current.children[char]
	}
	current.output = append(current.output, pattern)
}

func (ac *AhoCorasick) Build() {
	queue := []*TrieNode1{}
	for _, node := range ac.root.children {
		node.fail = ac.root
		queue = append(queue, node)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for char, child := range current.children {
			failNode := current.fail
			for failNode != nil {
				if failChild, exists := failNode.children[char]; exists {
					child.fail = failChild
					break
				}
				failNode = failNode.fail
			}
			if failNode == nil {
				child.fail = ac.root
			}
			child.output = append(child.output, child.fail.output...)
			queue = append(queue, child)
		}
	}
}

func (ac *AhoCorasick) Search(text string) map[string][]int {
	current := ac.root
	matches := make(map[string][]int)

	for i, char := range text {
		for current != ac.root && current.children[char] == nil {
			current = current.fail
		}
		if nextNode, exists := current.children[char]; exists {
			current = nextNode
		} else {
			current = ac.root
		}
		for _, pattern := range current.output {
			if _, exists := matches[pattern]; !exists {
				matches[pattern] = []int{}
			}
			matches[pattern] = append(matches[pattern], i-len(pattern)+1)
		}
	}

	return matches
}
