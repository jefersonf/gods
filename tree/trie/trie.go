// simple Trie implementation
// limitations: does not work with non-english alphabet
// this means that it won't work for words with more than one byte unicode size.
package trie

type Trie struct {
	levels []Level
}

type Level struct {
	encodedAlphabet rune
	counter         int
	output          bool
}

func (t *Trie) AddString(s string) {
	var (
		i int
		c rune
	)
	for i, c = range s {
		if len(t.levels) < i+1 {
			t.levels = append(t.levels, Level{
				encodedAlphabet: 1 << (c - 'a'),
				output:          false,
			})
		}
		t.levels[i].encodedAlphabet |= (1 << (c - 'a'))
	}
	t.levels[i].counter += 1
	t.levels[i].output = true
}

func (t *Trie) CheckString(s string) bool {
	if len(s) > len(t.levels) {
		return false
	}
	var (
		i int
		c rune
	)
	for i, c = range s {
		pos := rune(1 << (c - 'a'))
		if t.levels[i].encodedAlphabet&pos == 0 {
			return false
		}
	}
	return t.levels[i].output
}

func (t *Trie) Height() int {
	return len(t.levels)
}

func New() *Trie {
	return &Trie{
		levels: make([]Level, 0),
	}
}