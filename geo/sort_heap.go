package geo

import (
	"algo/types"
)

type heap types.IPoints

func IHeapSort(p types.IPoints) {
	heap(p).Sort()
}

func (p heap) Sort() {
	p.buildHeap()

	for i := len(p) - 1; i > 0; i-- {
		p[0], p[i] = p[i], p[0]
		p.heapify(0, i)
	}
}

// half of
func (p heap) buildHeap() {
	for i := len(p)/2 - 1; i >= 0; i-- {
		p.heapify(i, len(p))
	}
}

// compare start point with all tree
func (p heap) heapify(start int, end int) {
	largest := start
	left := 2*start + 1
	right := 2*start + 2

	if left < end && !p[left].Less(p[start]) {
		largest = left
	}

	if right < end && !p[right].Less(p[largest]) {
		largest = right
	}

	if largest == start {
		return
	}

	p[largest], p[start] = p[start], p[largest]
	p.heapify(largest, end)
}
