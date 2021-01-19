package ConvexHull

import (
	"algo/geo"
	"algo/types"
	"math"
)

func NativeGrahamScan(pts types.IPoints) types.IPoints {
	if len(pts) <= 3 {
		return pts
	}

	// lowest Y point
	var lowest int
	low := pts[0]

	for i := range pts {
		if pts[i].Y < pts[lowest].Y {
			lowest = i
			low = pts[i]
		}
	}

	// swap lowest point to the end of arr
	if len(pts)-1 != lowest {
		pts[lowest], pts[len(pts)-1] = pts[len(pts)-1], pts[lowest]
	}

	geo.ReversePolarSort(pts[:len(pts)-1], pts[len(pts)-1])

	// if all points in one line so we should handle them
	firstAngel := math.Atan2(pts[0].Y-low.Y, pts[0].X-low.X)
	lastAngel := math.Atan2(pts[len(pts)-2].Y-low.Y, pts[len(pts)-2].X-low.X)
	if firstAngel == lastAngel {
		return types.IPoints{low, pts[0]}
	}

	// 3 pts serchent in hull: n-1, n-2 and 0 idx
	list := types.IPoints{pts[len(pts)-2], pts[len(pts)-1]}
	for i := 0; i < len(pts)-1; i++ {
		for IsLeftTurn(list[len(list)-2], list[len(list)-1], pts[i]) {
			list = list[:len(list)-1]
		}

		list = append(list, pts[i])
	}

	return list[:len(list)-1]
}

func IsLeftTurn(p1, p2, p3 types.IPoint) bool {
	return (p2.X-p1.X)*(p3.Y-p1.Y)-(p2.Y-p1.Y)*(p3.X-p1.X) > 0
}
