package treap

type Treap struct {
	rootNode *node
}

type node struct {
	key      int
	priority int
	left     *node
	right    *node
}

func Split(key int, t, l, r *node) {
	if t == nil {
		l = nil
		r = nil
	} else if t.key <= key {
		Split(key, t.right, t.right, r)
		l = t
	} else {
		Split(key, t.left, l, t.left)
		r = t
	}
}

func NewNode(key int, priority int) *node {
	return &node{
		key:      key,
		priority: priority,
		left:     nil,
		right:    nil,
	}
}
