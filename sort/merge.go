package sort

type MergeSort struct{}

func NewMergeSort() MergeSort {
	return MergeSort{}
}

func (m *MergeSort) Sort(in []int) {
	cop := make([]int, len(in))
	copy(cop, in)

	m.sort(cop, in, 0, len(cop))
}

func (m *MergeSort) sort(a, result []int, start, end int) {
	switch end - start {
	case 0, 1:
		return
	case 2:
		if result[start] > result[start+1] {
			result[start+1], result[start] = result[start], result[start+1]
		}

		return
	}

	mid := (end + start) / 2

	m.sort(result, a, start, mid)
	m.sort(result, a, mid, end)

	i, idx, j := start, start, mid

	for idx < end {
		if j >= end || (i < mid && a[i] < a[j]) {
			result[idx] = a[i]
			i++
		} else {
			result[idx] = a[j]
			j++
		}

		idx++
	}
}
