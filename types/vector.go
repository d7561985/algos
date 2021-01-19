package types

type Vector struct {
	left, right *Vector

	Idx int
	v   []int
}

func CreateVector(idx int) *Vector {
	return &Vector{Idx: idx}
}

func (v *Vector) Data() []int {
	return v.v
}

func (v *Vector) Begin() *Vector {
	if v.left != nil {
		return v.left.Begin()
	}

	return v
}

func (v *Vector) Next() *Vector {
	return v.right
}

func (v *Vector) lookup(idx int) *Vector {
	if v.Idx == idx {
		return v
	}

	// go left
	if v.Idx > idx {
		if v.left != nil && v.left.Idx >= idx {
			return v.left.lookup(idx)
		}

		vv := CreateVector(idx)

		vv.left, v.left, vv.right = v.left, vv, v

		if vv.left != nil {
			vv.left.right = vv
		}

		return vv
	}

	// go right
	if v.right != nil && v.right.Idx <= idx {
		return v.right.lookup(idx)
	}

	vv := CreateVector(idx)
	vv.right, v.right, vv.left = v.right, vv, v

	if vv.right != nil{
		vv.right.left = vv
	}

	return vv
}

func (v *Vector) Put(idx int, val int) {
	vv := v.lookup(idx)
	vv.v = append(vv.v, val)
}
