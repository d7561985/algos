package power_items

import (
	"testing"
)

func TestRS(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RS(tt.args.in); got != tt.want {
				t.Errorf("RS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRS(b *testing.B) {
	benchmarks := []struct {
		name string
		in []int
	}{
		{"x", []int{1, 2, 3, 4, 5}},
		{"x2", []int{100, 2, 3, 4, 5, 5555}},
		{"x3}", []int{}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RS(bm.in)
			}
		})
	}
}

func BenchmarkS(b *testing.B) {
	benchmarks := []struct {
		name string
		in []int
	}{
		{"x", []int{1, 2, 3, 4, 5}},
		{"x2", []int{100, 2, 3, 4, 5, 5555}},
		{"x3}", []int{}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				S(bm.in)
			}
		})
	}
}
