package power_items

type MergeSort struct {
	in []int
}

// only for sorted elements
func Merge(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))

	for !(len(a) == 0 && len(b) == 0) {
		switch {
		case len(a) == 0:
			res = append(res, b[0])
			b = b[1:]
		case len(b) == 0:
			res = append(res, a[0])
			a = a[1:]
		case a[0] > b[0]:
			res = append(res, b[0])
			b = b[1:]
		default:
			res = append(res, a[0])
			a = a[1:]
		}
	}

	return res
}

func MS(in []int) []int {
	switch len(in) {
	case 0:
		panic("unpossible")
	case 1:
		return in
	case 2:
		if in[0] > in[1] {
			in[1] = in[0]
		}

		return in
	}

	left := MS(in[0 : len(in)/2])
	right := MS(in[len(in)/2:])

	return Merge(left, right)
}

func LineSort(in []int) {
	for i := 1; i < len(in); i++ {
		if in[i-1] > in[i] {
			in[i], in[i-1] = in[i-1], in[i]

			if len(in) <= 1000 {
				LineSort(in)
			}

			i = 1
			continue
		}
	}

	return
}

func InsertSort(in []int) {
	for i := 1; i < len(in); i++ {
		j, v := i-1, in[i]

		for j >= 0 && in[j] > v {
			in[j+1] = in[j]
			j--
		}

		in[j+1] = v
	}

	return
}

// How perform swap
func InsertSort2(in []int) []int {
	for i := 1; i < len(in); i++ {
		j, v := i-1, in[i]

		for j >= 0 && in[j] > v {
			j--
		}

		copy(in[j+2:i+1], in[j+1:i])
		in[j+1] = v
	}

	return in
}

// O(n^2)
func SelectSort(in []int) []int {
	for i := len(in) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if in[j] < in[i] {
				continue
			}

			in[j], in[i] = in[i], in[j]
		}
	}

	return in
}
