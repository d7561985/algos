package ConvexHull

import (
	"algo/types"
	"reflect"
	"testing"
)

func TestNativeGrahamScan(t *testing.T) {
	type args struct {
		pts types.IPoints
	}
	tests := []struct {
		name string
		args args
		want types.IPoints
	}{
		{
			"",
			args{pts: types.IPoints{
				{1, 1}, {0, 0}, {2, 2}, {5, 3}, {5, 1},
				{6, 5}, {7, 4}, {8, 3}, {10, 1},
				{5, 0}, {1, 3}, {2, 5}, {6, 7}, {11, 6},
				{13, 2}, {7, -1},
			}}, []types.IPoint{},

		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NativeGrahamScan(tt.args.pts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NativeGrahamScan() = %v, want %v", got, tt.want)
			}
		})
	}
}
