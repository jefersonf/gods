package ahocorasick

const K = 26

type Vertex struct {
	next   [K]int
	output bool
	p      int
	pch    rune
	link   int
	goArr  [K]int
}

func NewVertex(p int, ch rune) Vertex {
	v := new(Vertex)
	v.p = p
	v.pch = ch
	for i := 0; i < K; i++ {
		v.next[i] = -1
		v.goArr[i] = -1
	}
	return *v
}

var tree []Vertex

func addString(s string) {
	v := 0
	for _, ch := range s {
		c := ch - 'a'
		if tree[v].next[c] == -1 {
			tree[v].next[c] = len(tree)
			tree = append(tree, NewVertex(v, ch))
		}
		v = tree[v].next[c]
	}
	tree[v].output = true
}

func getLink(v int) int {
	if tree[v].link == -1 {
		if v == 0 || tree[v].p == 0 {
			tree[v].link = 0
		} else {
			tree[v].link = goFunc(getLink(tree[v].p), tree[v].pch)
		}
	}
	return tree[v].link
}

func goFunc(v int, ch rune) int {
	c := ch - 'a'
	if tree[v].goArr[c] == -1 {
		if tree[v].next[c] != -1 {
			tree[v].goArr[c] = tree[v].next[c]
		} else {
			tree[v].goArr[c] = v
			if v != 0 {
				tree[v].goArr[c] = goFunc(getLink(v), ch)
			}
		}
	}
	return tree[v].goArr[c]
}
