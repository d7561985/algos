package power_items

import (
	"math"
)

type Vertex string

type Edge struct {
	Root     Vertex
	Children []Edge
	Weigh    int
}

type Result struct {
	Cost  int
	Route []Vertex
}

func DeykstraGet(root Edge, target Vertex) (res Result) {
	vals := dLook(root, target, res)
	res.Cost = math.MaxInt64

	for _, val := range vals {
		if val.Cost < res.Cost {
			res = val
		}
	}

	return res
}

func dLook(root Edge, target Vertex, input Result) (res []Result) {
	if root.Root == target {
		res = append(res, input)
		return
	}

	for idx := range root.Children {
		child := root.Children[idx]

		v := dLook(child, target, Result{Cost: input.Cost + child.Weigh, Route: append(input.Route, child.Root)})
		if v == nil {
			continue
		}

		res = append(res, v...)
	}

	return
}
