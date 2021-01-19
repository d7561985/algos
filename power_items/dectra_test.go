package power_items

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeykstraGet(t *testing.T) {
	type args struct {
		in     Edge
		target Vertex
	}

	L := make(map[string]*Edge)
	for i := 'A'; i < 'Z'; i++ {
		L[string(i)] = &Edge{Root: Vertex(i), Weigh: int(i)}
	}

	L["E"].Children = []Edge{{Root: "Z", Weigh: 100}}
	L["F"].Children = []Edge{{Root: "Z", Weigh: 999}}

	L["C"].Children = []Edge{
		{
			Root:     "D",
			Children: L["D"].Children,
			Weigh:    1,
		},
		{
			Root:     "E",
			Children: L["E"].Children,
			Weigh:    5,
		},
		{
			Root:     "F",
			Children: L["F"].Children,
			Weigh:    10,
		},
	}

	L["A"].Children = []Edge{{
		Root:     "C",
		Children: L["C"].Children,
		Weigh:    100,
	}}

	L["B"].Children = []Edge{{
		Root:     "C",
		Children: L["C"].Children,
		Weigh:    1,
	}}
	fmt.Println(L)

	tests := []struct {
		name    string
		args    args
		wantRes Result
	}{
		{
			name: "",
			args: args{
				in: Edge{
					Root: "root",
					Children: []Edge{
						*L["A"],
						*L["B"],
					},
					Weigh: 0,
				},
				target: "Z",
			},
			wantRes: Result{
				Cost:  172,
				Route: []Vertex{"B", "C", "E", "Z"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := DeykstraGet(tt.args.in, tt.args.target); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("DeykstraGet() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestA(t *testing.T) {
	root := map[Vertex][]E{
		"1": {{"2", 7}, {"3", 9}, {"6", 14}},
		"2": {{"1", 7}, {"3", 10}, {"4", 15}},
		"3": {{"1", 9}, {"2", 10}, {"4", 11}, {"6", 2}},
		"4": {{"2", 15}, {"3", 11}, {"5", 6}},
		"6": {{"1", 14}, {"3", 2}, {"5", 9}},
	}

	d := NewD(root, "5")
	d.Do("1")
	fmt.Printf("%#v", d.Mark)
}

func TestФ(t *testing.T) {
	fmt.Println(float32(1)/3, float64(1)/3)

}
