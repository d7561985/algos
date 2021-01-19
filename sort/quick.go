package sort

import (
	"algo/power_items"
)

type Quick struct {
}

func (q *Quick) Sort(in []int) {
	q.quickSort(in, 0, len(in)-1)
}

func (q *Quick) quickSort(in []int, left, right int) {
	if left >= right {
		return
	}

	if diff := right - left + 1; diff < 7{
		power_items.LineSort(in[left:right])
		return
	}

/*	switch diff := right - left + 1; diff {
	case 1:
		return
	case 2:
		if in[left] > in[right] {
			in[left], in[right] = in[right], in[left]
		}

		return
	}*/

	p := q.partition(in, left, right, left+(right-left)/2)
	q.quickSort(in, left, p-1)
	q.quickSort(in, p+1, right)
}

func (q *Quick) partition(ar []int, left, right, pivotIndex int) (store int) {
	pivot := ar[pivotIndex]
	store = left

	// move pivot value to the right
	ar[right], ar[pivotIndex] = pivot, ar[right]

	// everything not exceed pivot move to the left
	for idx := left; idx < right; idx++ {
		if ar[idx] <= pivot {
			ar[idx], ar[store] = ar[store], ar[idx]
			store++
		}
	}

	ar[right], ar[store] = ar[store], ar[right]
	return store
}

// another implementation but with mem exposition
func QuickSortUseMem(in []int) []int {
	if len(in) < 7 {
		power_items.InsertSort(in)
		return in
	}

	idx := len(in) / 2
	left, right := make([]int, 0, len(in)), make([]int, 0, len(in))
	pivot := in[idx]

	for i, v := range in {
		if i == idx {
			continue
		}

		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	idx = len(left)
	in[len(left)] = pivot

	copy(in[:idx], QuickSortUseMem(left))
	copy(in[idx+1:], QuickSortUseMem(right))

	return in
}
