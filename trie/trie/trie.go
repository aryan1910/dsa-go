// https://leetcode.com/problems/implement-trie-prefix-tree/
//
// Implement insert, search and startsWith
package trie

type Node struct {
	links map[rune]*Node
	isEnd bool
}

func NewNode() *Node {
	return &Node{
		links: make(map[rune]*Node),
		isEnd: false,
	}
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: NewNode(),
	}
}

func (t *Trie) Insert(word string) {
	temp := t.root

	for _, c := range word {
		_, ok := temp.links[c]
		if !ok {
			temp.links[c] = NewNode()
		}
		temp = temp.links[c]
	}
	temp.isEnd = true
}

func (t *Trie) Search(word string) bool {
	temp := t.root
	for _, c := range word {
		_, ok := temp.links[c]
		if !ok {
			return false
		}
		temp = temp.links[c]
	}
	return temp.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	temp := t.root
	for _, c := range prefix {
		_, ok := temp.links[c]
		if !ok {
			return false
		}
		temp = temp.links[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
