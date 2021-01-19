package sort

// approach from book

type Entry struct {
	element int
	next    *Entry
}

type Bucket struct {
	size int
	head *Entry
}

type Buckets []*Bucket

func NewBucket2(size int) Buckets {
	return make([]*Bucket, size)
}

func (b Buckets) SortPointers(in []int) {
	for i, v := range in {
		k := b.hash(v)
		e := &Entry{element: in[i]}

		if b[k] == nil {
			b[k] = &Bucket{}
		}

		if b[k].head != nil {
			e.next = b[k].head
		}

		b[k].head = e
		b[k].size++
	}

	b.extract(in)
}

func (b Buckets) extract(in []int) {
	var idx, low int

	for i := range b {
		if b[i] == nil {
			continue
		}

		var ptr *Entry
		ptr = b[i].head

		if b[i].size == 1 {
			in[idx] = ptr.element
			b[i].size = 0
			idx++
			continue
		}

		low = idx
		in[idx] = ptr.element
		ptr = ptr.next
		idx++

		for ptr != nil {
			j := idx - 1

			for j >= low && in[j] > ptr.element {
				in[j+1] = in[j]
				j--
			}

			in[j+1] = ptr.element
			ptr = ptr.next
			idx++
		}

		b[i].size = 0
	}
}

func (b Buckets) hash(v int) int {
	return v / len(b)
}
