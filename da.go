package ponda

// Node is Double Array Node
type Node struct {
	Base  int32
	Check int32
}

// DoubleArray is data structure of Double-Array trie.
type DoubleArray struct {
	Nodes []Node
	Dict  RuneDic
}

// Builder has data for buinding Double Array.
type Builder struct {
	nodeCap int
	wordCap int
	words   [][]rune
}

// NewBuilder inits Double Array.
func NewBuilder(options ...func(*Builder)) *Builder {
	b := &Builder{}
	for _, option := range options {
		option(b)
	}
	b.words = make([][]rune, 0, b.wordCap)

	return b
}

// NodeCap set capacity of nodes.
func NodeCap(cap int) func(*Builder) {
	return func(da *Builder) {
		da.nodeCap = cap
	}
}

// WordCap set capacity of words.
func WordCap(cap int) func(*Builder) {
	return func(da *Builder) {
		da.wordCap = cap
	}
}

// Add add str to words list.
func (b *Builder) Add(s string) *Builder {
	b.words = append(b.words, []rune(s))
	return b
}

// Build builds Double Array from Builder.
func (b *Builder) Build() *DoubleArray {
	return &DoubleArray{
		Nodes: make([]Node, 0, b.nodeCap),
		Dict:  buildRuneDict(b.words),
	}
}

// Lookup lookups key in Double Array.
func (da *DoubleArray) Lookup(key []rune) (int32, bool) {
	var index int32
	for i := len(key) - 1; i >= 0; i-- {
		branch, ok := da.Dict[key[i]]
		if !ok {
			return 0, false
		}

		base := da.Nodes[index].Base
		if base < 0 {
			base = -base
		}
		next := base + branch
		if next >= int32(len(da.Nodes)) || da.Nodes[next].Check != index {
			return 0, false
		}
		index = next
	}

	if da.Nodes[index].Base >= 0 {
		return 0, false
	}
	return index, true
}
