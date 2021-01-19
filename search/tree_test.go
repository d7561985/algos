package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type X int

func (x X) Comp(t interface{}) int {
	switch {
	case x == t.(X):
		return 0
	case x > t.(X):
		return 1
	default:
		return -1
	}
}

func TestTreeNode_ForEach(t1 *testing.T) {
	tests := []struct {
		name   string
		in     []X
		target []X
	}{
		{
			"",
			[]X{7, 1, 2, 4, 5, 6, 8, 3},
			[]X{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			v := BST{}

			for i := range tt.in {
				v.Add(tt.in[i])
			}

			var tar []X
			v.ForEach(func(i Comparable) {
				tar = append(tar, i.(X))
			})

			assert.Equal(t1, tt.target, tar)
		})
	}
}
