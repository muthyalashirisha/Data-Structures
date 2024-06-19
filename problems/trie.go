package problems

// TrieNode represents a node in the trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the trie structure
type Trie struct {
	root *TrieNode
}

// NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = NewTrieNode()
		}
		current = current.children[char]
	}
	current.isEnd = true
}

// Search checks if a word is in the trie
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}
	return current.isEnd
}

// StartsWith checks if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}
	return true
}
