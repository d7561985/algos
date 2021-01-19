package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBucketSort_Sort(t *testing.T) {
	type fields struct {
		in    []int
		block int
	}
	type args struct {
		in []int
	}
	tests := []struct {
		name   string
		fields fields
		res    []int
	}{
		{
			"",
			fields{in: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 11}, block: 4},
			[]int{1,2,3,4,5,6,7,8,9,10, 11}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBucketSort(tt.fields.block)
			b.Sort(tt.fields.in)
			assert.Equal(t, tt.res, tt.fields.in)
		})
	}
}
