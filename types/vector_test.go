package types

import (
	"fmt"
	"testing"
)

func TestCreateVector(t *testing.T) {
	v := CreateVector(0)
	for i := 100; i > 0; i-- {
		v.Put(i, i)
	}

	for itr := v.Begin(); itr != nil; itr = itr.Next() {
		fmt.Println(itr.v)
	}
}
