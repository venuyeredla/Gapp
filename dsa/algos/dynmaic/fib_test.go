package dynmaic

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	input := 8
	fmt.Printf("%v = %v ", input, fibC(8))
}
