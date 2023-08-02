package maths

import (
	"fmt"
	"math"
)

// pow(2,n) -- Exponential
func Subset(arr []int) {
	caridnality := len(arr)
	var powerSetSize int = int(math.Pow(2, float64(caridnality)))
	for counter := 0; counter < powerSetSize; counter++ {
		for j := 0; j < caridnality; j++ {
			if (counter & (1 << j)) > 0 {
				fmt.Print(arr[j])
			}
		}
		fmt.Println(" ")
	}
}

// Subset sum problem is to find subset of elements that are selected from a given set whose sum adds up to a given number K.
// S,  s={a  a S  }  Sum(s) =K
func SubsetSum() {

}
