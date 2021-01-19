package search

type Comparable interface {
	// -1 - cur is less
	// 0 - cur is equal
	// 1 - cur is higher
	Comp(target interface{}) int
}

type TreeNode struct {
	V           Comparable
	Left, Right *TreeNode
}

func NewNode(v Comparable) *TreeNode {
	return &TreeNode{V: v}
}

func (t *TreeNode) Add(v Comparable) {
	switch {
	case t.V == nil:
		t.V = v
	case t.V.Comp(v) > 0:
		if t.Left == nil {
			t.Left = &TreeNode{}
		}

		t.Left.Add(v)
	default:
		if t.Right == nil {
			t.Right = &TreeNode{}
		}
		t.Right.Add(v)
	}
}

func (t *TreeNode) ForEach(fn func(Comparable)) {
	if t.Left != nil {
		t.Left.ForEach(fn)
	}

	fn(t.V)

	if t.Right != nil {
		t.Right.ForEach(fn)
	}
}

type BST struct {
	Root *TreeNode
}

func (b *BST) Add(v Comparable) {
	if b.Root == nil {
		b.Root = NewNode(v)
		return
	}

	b.Root.Add(v)
}

func (b *BST) Contains(v Comparable) bool {
	node := b.Root

	for node != nil {
		switch node.V.Comp(v) {
		case 0:
			return true
		case 1:
			node = node.Right
		default:
			node = node.Left
		}
	}

	return false
}

func (b *BST) ForEach(fn func(v Comparable)) {
	if b.Root == nil {
		return
	}

	b.Root.ForEach(fn)
}
