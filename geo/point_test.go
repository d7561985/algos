package geo

import (
	"algo/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPoints_Sort(t *testing.T) {
	tests := []struct {
		name string
		p    types.IPoints
		res  types.IPoints
	}{
		{
			"",
			types.IPoints{{10, -1}, {10, 0}, {1, 1}, {0, 0}},
			types.IPoints{{0, 0}, {1, 1}, {10, 0}, {10, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IHeapSort(tt.p)
			assert.Equal(t, tt.res, tt.p)
		})
	}
}
