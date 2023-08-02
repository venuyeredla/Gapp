package linear

import (
	"fmt"
)

func PrintShape(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j < i {
				fmt.Print("   ")
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Println()
	}
}

func PrintTriangle(size int) {
	position := (2*size - 1) / 2
	for i := 1; i <= size; i++ {
		for k := 1; k <= position; k++ {
			fmt.Print("  ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println()
		position--
	}
}
