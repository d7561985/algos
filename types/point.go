package types

type IPoint struct {
	X, Y float64
}

type IPoints []IPoint

func (i *IPoint) Less(p IPoint) bool {
	switch {
	case i.X == p.X:
		return i.Y < p.Y
	default:
		return i.X < p.X
	}
}

