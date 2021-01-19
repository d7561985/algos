package sort

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestHeap_Sort(t *testing.T) {
	h := new(Heap)
	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	h.Sort(&arr)
	fmt.Println(arr)
}

func genRandIntArray() []int {
	rand.Seed(time.Now().Unix())
	sort := make([]int, 10000)
	for i := range sort {
		sort[i] = int(rand.Int63n(math.MaxInt64))
	}

	return sort
}

func BenchmarkSort(b *testing.B) {
	s := genRandIntArray()

	b.Run("heap", func(b *testing.B) {
		v := make([]int, len(s))
		copy(v, s)
		h := new(Heap)

		for i := 0; i < b.N; i++ {
			h.Sort(&v)
		}
	})
}
