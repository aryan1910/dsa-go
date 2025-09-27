// https://www.naukri.com/code360/problems/implement-trie_1387095
// insert, count words equal to, count words starting with, erase
package trie2

type Node struct {
	links       map[rune]*Node
	countEnd    int
	countPrefix int
}

func NewNode() *Node {
	return &Node{
		links:       make(map[rune]*Node),
		countEnd:    0,
		countPrefix: 0,
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
		temp.countPrefix++
	}
	temp.countEnd++
}

func (t *Trie) countWordsEqualTo(word string) int {
	temp := t.root
	for _, c := range word {
		_, ok := temp.links[c]
		if !ok {
			return 0
		}
		temp = temp.links[c]
	}
	return temp.countEnd
}

func (t *Trie) countWordsStartingWith(prefix string) int {
	temp := t.root
	for _, c := range prefix {
		_, ok := temp.links[c]
		if !ok {
			return 0
		}
		temp = temp.links[c]
	}
	return temp.countPrefix
}

func (t *Trie) erase(prefix string) {
	temp := t.root
	for _, c := range prefix {
		_, ok := temp.links[c]
		if !ok {
			return
		}
		temp = temp.links[c]
		temp.countPrefix--
	}
	temp.countEnd--
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
