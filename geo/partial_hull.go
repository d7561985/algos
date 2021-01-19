package geo

import (
	"algo/types"
	"fmt"
)

type PartialHull types.IPoints

func NewPartialHull(a, b types.IPoint) PartialHull {
	return PartialHull{a, b}
}

func (p *PartialHull) RemoveMiddleOfLastThree() bool {
	if !p.HasThree() {
		return false
	}

	pos := len(*p) - 2
	*p = append((*p)[:pos], (*p)[pos+1:]...)
	return true
}

func (p PartialHull) HasThree() bool {
	return len(p) > 2
}

func (p PartialHull) AreLastThreeNotRight() bool {
	if !p.HasThree() {
		return false
	}

	pos := len(p) - 3

	x1, y1 := p[pos].X, p[pos].Y
	x2, y2 := p[pos+1].X, p[pos+1].Y
	x3, y3 := p[pos+2].X, p[pos+2].Y

	val1 := (x2 - x1) * (y3 - y1)
	val2 := (y2 - y1) * (x3 - x1)
	diff := val1 - val2

	fmt.Println(p[pos], p[pos+1], p[pos+2], val1, val2, diff)

	if diff >= 0 {
		return true
	}

	return false
}
