package ai

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {

	Matrix := [][]int{{2, 5, 8}, {4, 0, -1}}

	/*Matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}

	/* Matrix2 := [][]int{{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9}}

	{13, 14, 15, 16}
	*/
	fmt.Printf("Output = %v", SpiralForm(Matrix))
}

func TestMatrixRotation(t *testing.T) {

	Matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	RotateMatrixBy90(Matrix)
}
