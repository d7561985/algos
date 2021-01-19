package sort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuick_Sort(t *testing.T) {
	type args struct {
		in *[]int
	}

	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"", []int{9, 5, 2, 7, 0}, []int{0, 2, 5, 7, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quick{}
			q.Sort(tt.in)
			assert.Equal(t, tt.want, tt.in)
		})
	}
}

func TestA(t *testing.T) {
	Q := &[]int{1}
	x := *Q
	x[0] = 2
	fmt.Println(*Q)
}

func TestQuickSortUseMem(t *testing.T) {
	v := genRandIntArray()

	res := QuickSortUseMem(v)

	sort.SliceStable(v, func(i, j int) bool {
		return v[i] < v[j]
	})

	assert.Equal(t, v, res)
}

func BenchmarkQuick_Sort(b *testing.B) {
	v := genRandIntArray()
	q := &Quick{}

	for i := 0; i < b.N; i++ {
		l := make([]int, len(v))
		copy(l, v)

		q.Sort(l)
	}
}
