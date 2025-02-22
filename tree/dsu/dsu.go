package dsu

type DSU struct {
	parent []uint64
	size   []uint64
}

// MakeSet assigns the initial set ID for a node
func (t *DSU) MakeSet(v uint64) {
	t.parent[v] = v
	t.size[v] = 1
}

// UnionSets joins the sets belonging to two different nodes
func (t *DSU) UnionSets(a, b uint64) {
	a = t.FindSet(a)
	b = t.FindSet(b)
	x := min(a, b)
	y := max(a, b)
	if x != y {
		if t.size[x] < t.size[y] {
			x, y = y, x
		}
		t.parent[y] = x
		t.size[x] += t.size[y]
	}
}

// FindSet returns the set ID of a node
func (t *DSU) FindSet(v uint64) uint64 {
	if v == t.parent[v] {
		return v
	}
	t.parent[v] = t.FindSet(t.parent[v])
	return t.parent[v]
}

// IsSameSet returns a true when u and v are in the same set, otherwise false
func (t *DSU) IsSameSet(u, v uint64) bool {
	return t.FindSet(u) == t.FindSet(v)
}

// Size returns the size of the DSU
func (t *DSU) Size() uint64 {
	return uint64(len(t.parent) - 1)
}

// New creates a new DSU instance by passing a size
func New(n uint64) *DSU {
	dsu := &DSU{
		parent: make([]uint64, n+1),
		size:   make([]uint64, n+1),
	}
	for i := uint64(1); i <= n; i++ {
		dsu.MakeSet(i)
	}
	return dsu
}
