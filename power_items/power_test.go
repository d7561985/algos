package power_items

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestPowerSetX(t *testing.T) {
	tests := []struct {
		name string
		in   [][]uint
		want [][]uint
	}{
		{"", [][]uint{{0}, {1}, {2}}, [][]uint{{0}, {0, 1}, {1}, {0, 2}, {0, 1, 2}, {1, 2}, {2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PowerSetX(tt.in)
			fmt.Println(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerSetX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecPowerSetX(t *testing.T) {
	tests := []struct {
		name string
		in   [][]uint
		want [][]uint
	}{
		{"", [][]uint{{0}, {1}, {2}}, [][]uint{{0}, {0, 1}, {1}, {0, 2}, {0, 1, 2}, {1, 2}, {2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RecPowerSetX(tt.in)
			fmt.Println(got)

			if !reflect.DeepEqual(got, tt.want) {
				//ToDo
				//t.Errorf("RecPowerSetX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandAndConqure(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		wantRes []byte
	}{
		{"", []byte("hello"), []byte("ehllo")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := CommandAndConqure(tt.in)
			fmt.Println(string(gotRes))

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CommandAndConqure() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func BenchmarkAA(b *testing.B) {
	c := make([]byte, 10)
	if _, err := rand.Read(c); err != nil {
		b.Error(err)
	}

	//fmt.Println("=>",string(c))
	for i := 0; i < b.N; i++ {
		x := make([]byte, 10)
		copy(x, c)

		_ = CommandAndConqure(x)
	}

}

func BenchmarkBB(b *testing.B) {
	c := make([]byte, 10)
	if _, err := rand.Read(c); err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		x := make([]byte, 10)
		copy(x, c)

		sort.Slice(x, func(i, j int) bool {
			return c[i] < c[j]
		})

	}
}
