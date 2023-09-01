// simple Trie implementation
// limitations: does not work with non-english alphabet
// this means that it won't work for words with more than one byte unicode size.
package trie

type Trie struct {
	rootNode *Node
}

type Node struct {
	char rune
	next []*Node
}

func (t *Trie) AddString(word string) {
	curNode := t.rootNode
	for _, c := range word {
		index := c - 'a'
		if curNode.next[index] == nil {
			curNode.next[index] = NewNode(c)
		}
		curNode = curNode.next[index]
	}
}

func (t *Trie) SearchWord(word string) bool {
	curNode := t.rootNode
	for _, c := range word {
		index := c - 'a'
		if curNode.next[index] == nil {
			return false
		}
		curNode = curNode.next[index]
	}
	return true
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
