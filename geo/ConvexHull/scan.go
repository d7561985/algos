package ConvexHull

import (
	"algo/geo"
	"algo/types"
	"fmt"
)

type Scan struct {
	points types.IPoints
}

func New(in types.IPoints) *Scan {
	return &Scan{in}
}

func (s *Scan) Do() types.IPoints {
	geo.IHeapSort(s.points)
	if len(s.points) < 3 {
		return s.points
	}

	fmt.Println(s.points)
	n := len(s.points)

	upper := geo.NewPartialHull(s.points[0], s.points[1])
	for i := 2; i < len(s.points); i++ {
		upper = append(upper, s.points[i])

		for upper.HasThree() && upper.AreLastThreeNotRight() {
			upper.RemoveMiddleOfLastThree()
		}
	}

	lower := geo.NewPartialHull(s.points[n-1], s.points[n-2])
	for i := n - 3; i >= 0; i-- {
		lower = append(lower, s.points[i])

		for lower.HasThree() && lower.AreLastThreeNotRight() {
			lower.RemoveMiddleOfLastThree()
		}
	}

	hull := make(types.IPoints, len(upper)+len(lower)-2)

	copy(hull, upper)
	copy(hull[len(upper):], lower[1:len(lower)-1])

	fmt.Println(upper)
	fmt.Println(lower)
	return hull
}
