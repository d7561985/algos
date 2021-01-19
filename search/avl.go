package search

type AVLNode struct {
	V           Comparable
	Left, Right *AVLNode
	height      int
}

func (a *AVLNode) computeHeight() {
	height := -1

	if a.Left != nil && a.Left.height > height {
		height = a.Left.height
	}

	if a.Right != nil && a.Right.height > height {
		height = a.Right.height
	}

	a.height = height + 1
}

func (a *AVLNode) heightDifference() int {
	leftTarget, rightTarget := 0, 0

	if a.Left != nil {
		leftTarget = 1 + a.Left.height
	}

	if a.Right != nil {
		rightTarget = 1 + a.Right.height
	}

	return leftTarget - rightTarget
}

func (a *AVLNode) Add(v Comparable) *AVLNode {
	var newRoot = a

	switch {
	case a.V == nil:
		a.V = v
	case a.V.Comp(v) > 0:
		a.Left = a.addToSubTree(a.Left, v)

		if a.heightDifference() != 2 {
			break
		}

		if v.Comp(a.Left.V) <= 0 {
			newRoot = a.rotateRight()
		} else {
			newRoot = a.rotateLeftRight()
		}
	default:
		a.Right = a.addToSubTree(a.Right, v)

		if a.heightDifference() != -2 {
			break
		}

		if v.Comp(a.Right) > 0 {
			newRoot = a.rotateLeft()
		} else {
			newRoot = a.rotateRightLeft()
		}

	}

	newRoot.computeHeight()

	return newRoot
}

func (a *AVLNode) addToSubTree(parent *AVLNode, v Comparable) *AVLNode {
	if parent == nil {
		return &AVLNode{V: v}
	}

	return parent.Add(v)
}

func (a *AVLNode) rotateRight() *AVLNode {
	newRoot := a.Left
	newRoot.Right, a.Left = a, newRoot.Right

	return newRoot
}

func (a *AVLNode) rotateRightLeft() *AVLNode {
	child := a.Right
	newRoot := child.Left
	grand1, grand2 := newRoot.Left, newRoot.Right

	child.Left, a.Right = grand2, grand1
	newRoot.Left = a
	newRoot.Right = child

	child.computeHeight()
	a.computeHeight()

	return newRoot
}

func (a *AVLNode) rotateLeft() *AVLNode {
	newRoot := a.Right
	newRoot.Left, a.Right = a, newRoot.Left

	return newRoot
}

func (a *AVLNode) rotateLeftRight() *AVLNode {
	child := a.Left
	newRoot := child.Right
	grand1, grand2 := newRoot.Left, newRoot.Right
	child.Right, a.Left = grand1, grand2

	newRoot.Left, newRoot.Right = child, a

	child.computeHeight()
	a.computeHeight()

	return newRoot
}
