package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort_Sort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		in,res []int
	}{
		{"", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 11}, []int{1,2,3,4,5,6,7,8,9,10, 11}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MergeSort{}
			m.Sort(tt.in)
			assert.Equal(t, tt.res, tt.in)
		})
	}
}
