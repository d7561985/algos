package power_items

func S(in []int) (v int) {
	for _, val := range in {
		v += val
	}

	return
}

func S2(in []int) (res int) {
	for i := 0; i < len(in); i++ {
		res += in[i]
	}
	return
}

func RS(in []int) int {
	if len(in) == 0 {
		return 0
	}

	return in[0] + RS(in[1:])
}
