package trees

// Integer represents a generic interface for integer values.
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Float represents a generic interface for floating-point values.
type Float interface {
	float32 | float64
}

// Number represents a Numeric value.
type Number interface {
	Integer | Float
}

// BinaryIndexed represents a slice of numbers with support for efficient
// prefix sum computation. The Zero value is an empty slice.
type BinaryIndexed[T Number] struct {
	data []T
}

// New creates a new BinaryIndexed with the given elements.
func New[T Number](a ...T) *BinaryIndexed[T] {
	n := len(a)
	t := make([]T, n)
	copy(t, a)
	for i := range t {
		if j := i | (i + 1); j < n {
			t[j] += t[i]
		}
	}
	return &BinaryIndexed[T]{
		data: t,
	}
}

// Get returns the element at index i.
func (t *BinaryIndexed[T]) Get(i int) T {
	sum := t.data[i]
	j := i + 1
	j -= j & -j
	for i > j {
		sum -= t.data[i-1]
		i -= i & -i
	}
	return sum
}

// Set sets the element at index i to val.
func (t *BinaryIndexed[T]) Set(i int, val T) {
	val -= t.Get(i)
	for l := len(t.data); i < l; i |= i + 1 {
		t.data[i] += val
	}
}

// Add adds val to the element at index i.
func (t *BinaryIndexed[T]) Add(i int, val T) {
	for l := len(t.data); i < l; i |= i + 1 {
		t.data[i] += val
	}
}

// Sum returns the sum of the first i elements from 0 to i-1.
func (t *BinaryIndexed[T]) Sum(i int) T {
	var sum T
	for i > 0 {
		sum += t.data[i-1]
		i -= i & -i
	}
	return sum
}

// Len returns the number of elements in the BinaryIndexed tree.
func (t *BinaryIndexed[T]) Len() int {
	return len(t.data)
}

// SumRange returns the sum of elements from i to j-1.
func (t *BinaryIndexed[T]) SumRange(i, j int) T {
	sum := t.Sum(j) - t.Sum(i)
	return sum
}

// Append appends a new element into BinaryIndexed.
func (t *BinaryIndexed[T]) Append(val T) {
	i := len(t.data)
	t.data = append(t.data, 0)
	t.data[i] = val - t.Get(i)
}
