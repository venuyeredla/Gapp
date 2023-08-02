package ai

import (
	"fmt"
	"testing"
)

func TestVector(t *testing.T) {
	t.Skip()
	vector := VectorR(5)
	vector.Print()
	cVector := VectorC(5)
	cVector.Print()
	GenMatrix(2, 3).Print()
}

func TestDotProduct(t *testing.T) {
	a := NewVector(4)
	a.InitVector(2.0)
	b := NewVector(4)
	b.InitVector(2.0)
	scalar := DotProduct(a, b)
	fmt.Println(scalar)
}
