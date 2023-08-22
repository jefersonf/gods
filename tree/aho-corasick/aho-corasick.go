package aho_corasick

type Vertex struct {
   Next []int
   Output bool
}

type Trie struct {
   Vertices []Vertex
}

func (v *Vertex) newVertex() *Vertex {
   s := make([]int, 26)
   for i := range s {
     s[i] = -1
  } 
  return &Vertex{
   Next: s,
   Output: false,
}
}

func New(size int) *Trie {
   return &Trie{
    Vertices: make([]Vertex, 1),
}
}
