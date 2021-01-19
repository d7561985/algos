package ConvexHull

import (
	"algo/types"
	"reflect"
	"testing"
)

func TestScan_Do(t *testing.T) {
	type fields struct {
		points types.IPoints
	}
	tests := []struct {
		name   string
		fields fields
		want   types.IPoints
	}{
		{
			"",
			fields{
				points: []types.IPoint{
					{1, 1}, {0, 0}, {2, 2}, {5, 3}, {5, 1},
					{6, 5}, {7, 4}, {8, 3}, {10, 1},
					{5, 0}, {1, 3}, {2, 5}, {6, 7}, {11, 6},
					{13, 2}, {7, -1},
				},
			},
			[]types.IPoint{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scan{
				points: tt.fields.points,
			}
			if got := s.Do(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
