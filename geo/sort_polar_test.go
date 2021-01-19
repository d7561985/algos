package geo

import (
	"algo/types"
	"fmt"
	"testing"
)

func TestReversePolarSort(t *testing.T) {
	type args struct {
		in   types.IPoints
		base types.IPoint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"",
			args{
				in: []types.IPoint{
					{1, 1}, {0, 0}, {2, 2}, {5, 3}, {5, 1},
					{6, 5}, {7, 4}, {8, 3}, {10, 1},
					{5, 0}, {1, 3}, {2, 5}, {6, 7}, {11, 6},
					{13, 2}, {7, -1},
				},
				base: types.IPoint{7, -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReversePolarSort(tt.args.in, tt.args.base)
			fmt.Println(tt.args.in)
		})
	}
}
