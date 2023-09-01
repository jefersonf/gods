// simple Trie implementation
// limitations: does not work with non-english alphabet
// this means that it won't work for words with more than one byte unicode size.
package trie

import "strings"

type Trie struct {
	rootNode *Node
	height   int
}

type Node struct {
	char rune
	next []*Node
}

func (t *Trie) AddString(word string) {
	curNode := t.rootNode
	normWord := normalizeWord(word)
	for _, c := range normWord {
		index := c - 'a'
		if curNode.next[index] == nil {
			curNode.next[index] = NewNode(c)
		}
		curNode = curNode.next[index]
	}
	t.height = max(t.height, len(normWord))
}

func (t *Trie) SearchWord(word string) bool {
	curNode := t.rootNode
	normWord := normalizeWord(word)
	for _, c := range normWord {
		index := c - 'a'
		if curNode.next[index] == nil {
			return false
		}
		curNode = curNode.next[index]
	}
	return true
}

func (t *Trie) Height() int {
	return t.height
}

func normalizeWord(word string) string {
	return strings.ToLower(strings.ReplaceAll(word, " ", ""))
}

func NewNode(char rune) *Node {
	return &Node{
		char: char,
		next: make([]*Node, 26),
	}
}

func New() *Trie {
	return &Trie{rootNode: NewNode('$')}
}
