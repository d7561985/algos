package power_items

import (
	"reflect"
)

func RecPowerSetX(in [][]uint) [][]uint {
	res := make([][]uint, 0, len(in))
	copy(res, in)

	for i := range in {
		for j, re := range res {
			if reflect.DeepEqual(re, in[i]) {
				res = append(res[:j], res[j+1:]...)
				break
			}
		}

		res = append(res, RecPowerSetX(res)...)
		res = append(res, in[i])
	}

	return res
}

func PowerSetX(in [][]uint) [][]uint {
	res := make([][]uint, 0, len(in))

	for i := range in {
		for j := range res {
			res = append(res, append(res[j], in[i]...))
		}

		res = append(res, in[i])
	}

	return res
}

func mergeSort(a, b []byte) []byte {
	res := make([]byte, 0, len(a)+len(b))
	for !(len(a) == 0 && len(b) == 0) {
		switch {
		case len(a) == 0 || len(b) > 0 && a[0] >= b[0]:
			res = append(res, b[0])
			b = b[1:]
		case len(b) == 0 || len(a) > 0 && b[0] >= a[0]:
			res = append(res, a[0])
			a = a[1:]
		}
	}

	return res
}

func CommandAndConqure(in []byte) (res []byte) {
	if len(in) == 1 {
		return in
	}

	return mergeSort(CommandAndConqure(in[:len(in)/2]), CommandAndConqure(in[len(in)/2:]))
}
