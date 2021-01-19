package search

import (
	"fmt"
	"testing"
)

func TestLeftLeft(t *testing.T) {
	a := &AVLNode{}
	a = a.Add(X(50))
	a = a.Add(X(30))
	a= a.Add(X(10))

	fmt.Println(a)
}

func TestLeftRight(t *testing.T) {
	a := &AVLNode{}
	a = a.Add(X(50))
	a = a.Add(X(30))
	a= a.Add(X(40))

	fmt.Println(a)
}
