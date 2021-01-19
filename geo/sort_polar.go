package geo

import (
	"algo/types"
	"math"
	"sort"
)

/**
 * Compares two 2-dimensional points in the Cartesian plane by polar coordinate with
 * respect to base point.
 *
 * Function uses the {@link Math#atan2} function to perform computation.
 *
 * Returns Descending order, rather than ascending order.
 *
 * @param one first point to be compared
 * @param two second point to be compared
 * @return 0 if equal, +1 if one has smaller polar angle, -1 if one has larger polar angle.
 */
func ReversePolarSort(in types.IPoints, base types.IPoint) {
	sort.SliceStable(in, func(i, j int) bool {
		oneAngel := math.Atan2(in[i].Y-base.Y, in[i].X-base.X)
		twoAngel := math.Atan2(in[j].Y-base.Y, in[j].X-base.X)

		if oneAngel > twoAngel {
			return true
		} else if oneAngel < twoAngel {
			return false
		}

		// if same angle, then must order by decreasing magnitude
		// to ensure that the convex hull algorithm is correct
		if in[i].Y < in[j].Y {
			return false
		}

		return true
	})
}
