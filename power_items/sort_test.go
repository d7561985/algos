package power_items_test

import (
	"algo/power_items"
	s2 "algo/sort"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMS(t *testing.T) {
	type args struct {
		in []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"a", args{in: []int{5, 2, 4, 6, 0}}, []int{0, 2, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := power_items.MS(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"a", args{in: []int{10, 5, 2, 4, 6, 0, 9}}, []int{0, 2, 4, 5, 6, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s2.QuickSortUseMem(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortUseMem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"a", args{in: []int{5, 2, 4, 6, 0}}, []int{0, 2, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := power_items.InsertSort2(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"a", args{in: []int{5, 2, 4, 6, 0}}, []int{0, 2, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := power_items.SelectSort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	s := genRandIntArray(100, 10000)

	b.Run("ms", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			power_items.MS(v)
		}
	})

	b.Run("ms-book", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			m := s2.MergeSort{}
			m.Sort(v)
		}
	})
	b.Run("qs QuickSortUseMem", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			s2.QuickSortUseMem(v)
		}
	})

	b.Run("qs 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			q := s2.Quick{}
			q.Sort(v)
		}
	})

	b.Run("ls", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			power_items.LineSort(v)
		}
	})

	b.Run("standart", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			sort.SliceStable(v, func(i, j int) bool {
				return v[i] < v[j]
			})
		}
	})

	b.Run("insert-sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			power_items.InsertSort(v)
		}
	})

	b.Run("insert-sort2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			power_items.InsertSort2(v)
		}
	})

	b.Run("select-sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			power_items.SelectSort(v)
		}
	})

	b.Run("heap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			h := new(s2.Heap)
			h.Sort(&v)
		}
	})

	b.Run("bucket", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			buck := s2.NewBucketSort(512)
			buck.Sort(v)
		}
	})

	b.Run("bucket-book", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := make([]int, len(s))
			copy(v, s)

			h := s2.NewBucket2(512)
			h.SortPointers(v)
		}
	})
}

func genRandIntArray(size int, max int64) []int {
	rand.Seed(time.Now().Unix())
	sort := make([]int, size)
	for i := range sort {
		sort[i] = int(rand.Int63n(max))
	}

	return sort
}

func TestAa(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := []int{8, 8, 8, 8, 8, 8, 8}

	copy(c[:len(a)], a)
	copy(c[3+1:], b)
	fmt.Println(c)
}

func TestLineSort(t *testing.T) {
	v := genRandIntArray(1000, math.MaxInt64)
	e := make([]int, len(v))
	copy(e, v)

	power_items.LineSort(v)

	sort.SliceStable(e, func(i, j int) bool {
		return e[i] < e[j]
	})

	assert.Equal(t, e, v)
}

func TestAA(t *testing.T) {
	var i **int
	x := 6
	y := &x
	i = &y

	fmt.Println(**i)
}
