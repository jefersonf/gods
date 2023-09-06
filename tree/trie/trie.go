// simple Trie implementation
// limitations: does not work with non-english alphabet
// this means that it won't work for words with more than one byte unicode size.
package trie

import (
	"strings"
)

type Trie struct {
	rootNode *Node
	height   int
}

type Node struct {
	char           rune
	next           []*Node
	endStringCount int
}

func (t *Trie) AddString(s string) {
	currentNode := t.rootNode
	normWord := normalizeStr(s)
	for _, c := range normWord {
		index := c - 'a'
		if currentNode.next[index] == nil {
			currentNode.next[index] = NewNode(c)
		}
		currentNode = currentNode.next[index]
	}
	if len(s) > 0 {
		currentNode.endStringCount += 1
	}
	t.height = max(t.height, len(normWord))
}

func (t *Trie) Search(s string) bool {
	currentNode := t.rootNode
	normWord := normalizeStr(s)
	for _, c := range normWord {
		index := c - 'a'
		if currentNode.next[index] == nil {
			return false
		}
		currentNode = currentNode.next[index]
	}
	return true
}

func (t *Trie) Count(s string) (count int) {
	currentNode := t.rootNode
	normWord := normalizeStr(s)
	for _, c := range normWord {
		index := c - 'a'
		if currentNode.next[index] == nil {
			return
		}
		currentNode = currentNode.next[index]
	}
	count = currentNode.endStringCount
	return
}

func (t *Trie) Size() (wordCount int) {
	queue := make([]*Node, 26)
	copy(queue, t.rootNode.next)
	for len(queue) > 0 {
		topNode := queue[0]
		newQueue := make([]*Node, len(queue)-1)
		copy(newQueue, queue[1:])
		queue = newQueue
		if topNode == nil {
			continue
		}
		if topNode.endStringCount > 0 {
			wordCount += 1
		}
		if topNode.next != nil {
			queue = append(queue, topNode.next...)
		}
	}
	return
}

func (t *Trie) Height() int {
	return t.height
}

func normalizeStr(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", ""))
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
