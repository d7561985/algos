package power_items

import (
	"fmt"
	"math"
)

type E struct {
	Vertex Vertex
	Weigh  int
}

type D struct {
	Target  Vertex
	List    map[Vertex][]E
	Mark    map[Vertex]int
	Visited map[Vertex]struct{}
}

func NewD(in map[Vertex][]E, target Vertex) *D {
	return &D{Target: target, List: in, Mark: make(map[Vertex]int), Visited: make(map[Vertex]struct{})}
}

func (d *D) Do(root Vertex) {
	if len(d.Visited) == len(d.List) {
		return
	}

	var small = E{Weigh: math.MaxInt64}

	fmt.Println("root:", root, " ", small.Vertex)

	for _, item := range d.List[root] {
		if _, ok := d.Visited[item.Vertex]; ok {
			continue
		}

		if item.Weigh < small.Weigh {
			small = item
		}

		d.mark(item.Vertex, d.Mark[root] + item.Weigh)
	}

	d.Visited[root] = struct{}{}

	if small.Weigh == math.MaxInt64 {
		return
	}


	d.visit(small.Vertex, d.Mark[root] + small.Weigh)
	d.Do(small.Vertex)
}

func (d *D) mark(child Vertex, weight int) {
	if v, ok := d.Mark[child]; ok && v > weight || !ok {
		d.Mark[child] = weight
	}
}

func (d *D) visit(root Vertex, weight int) {
	for _, child := range d.List[root] {
		d.mark(child.Vertex, weight + child.Weigh)
	}
}
