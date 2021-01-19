package sort

import (
	"algo/power_items"
	"algo/types"
)

type BucketSort struct {
	store *types.Vector
	block int
	min   int
	max   int
}

func NewBucketSort(block int) *BucketSort {
	return &BucketSort{
		block: block,
	}
}

func (b *BucketSort) Sort(in []int) {
	b.scan(in)
	b.order(in)
}

func (b *BucketSort) hash(v int) int {
	return v / b.block
}

func (b *BucketSort) scan(in []int) {
	for i := range in {
		chain := b.hash(in[i])

		if b.store == nil {
			b.store = types.CreateVector(chain)
		}

		b.store.Put(chain, in[i])
	}
}

func (b *BucketSort) order(in []int) {
	w := 0

	for itr := b.store.Begin(); itr != nil; itr = itr.Next() {
		data := itr.Data()
		power_items.InsertSort(data)

		for j := range data {
			in[w] = data[j]
			w++
		}
	}
}
