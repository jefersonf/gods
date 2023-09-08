package dsu

type DSU struct {
	parent []uint64
	size   []uint64
}

func (t *DSU) MakeSet(v uint64) {
	t.parent[v] = v
	t.size[v] = 1
}

func (t *DSU) UnionSets(a, b uint64) {
	a = t.FindSet(a)
	b = t.FindSet(b)
	if a != b {
		if t.size[a] < t.size[b] {
			a, b = b, a
		}
		t.parent[b] = a
		t.size[a] += t.size[b]
	}
}

func (t *DSU) FindSet(v uint64) uint64 {
	if v == t.parent[v] {
		return v
	}
	t.parent[v] = t.FindSet(t.parent[v])
	return t.parent[v]
}

func (t *DSU) Size() uint64 {
	return uint64(len(t.parent) - 1)
}

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
