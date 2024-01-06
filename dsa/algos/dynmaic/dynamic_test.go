package dynmaic

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	input := 8
	fmt.Printf("%v = %v ", input, fibC(8))
}

func TestFactors(t *testing.T) {
	fmt.Printf("\nFibNth(%v) = %v ", 10, FibNth(10))
}

func TestCoingChange(t *testing.T) {
	A := []int{1, 2, 3}
	count := CoinChange(A, 4)
	t.Logf("Output is : %v", count)
}
