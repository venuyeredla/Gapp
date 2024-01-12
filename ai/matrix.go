package ai

import (
	"strconv"
	"strings"
)

/*
Input:

Output: 1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10

	Outer: 0 -> m/2
*/
func SpiralForm(M [][]int) string {

	R := len(M)
	C := len(M[0])
	collector := make([]int, 0)
	for i := 0; i <= (R / 2); i++ {
		row := i
		col := i
		for ; col < C-i; col++ {
			collector = append(collector, M[row][col])
		}
		col--
		row++
		// Column
		for ; row < R-i; row++ {
			collector = append(collector, M[row][col])
		}
		row--
		if row == i {
			continue
		}
		col--
		//Row
		for ; col >= i; col-- {
			collector = append(collector, M[row][col])
		}
		row--
		col++
		//Column
		if col == (C - i - 1) {
			continue
		}
		for ; row > i; row-- {
			collector = append(collector, M[row][col])
		}
	}
	var sb strings.Builder
	for _, val := range collector {
		sb.WriteString(strconv.Itoa(val))
		sb.WriteString(", ")
	}
	return sb.String()
}

/*
Input:

	1  2  3
	4  5  6
	7  8  9

Output:
	Left
	3  6  9
	2  5  8
	1  4  7

	right
*/

/*
Approach :
 1. New output matrix.
 2. In place.
*/
func RotateMatrixBy90(M [][]int) {

	R := len(M)
	C := len(M[0])
	collector := make([]int, 0)
	for i := 0; i < (R / 2); i++ {
		row := i
		col := i
		for col < C-i {
			collector = append(collector, M[row][col])
			col++
		}
		col--
		row++
		// Column
		for row < R-i {
			collector = append(collector, M[row][col])
			row++
		}
		row--
		col--
		//Row
		for col >= i {
			collector = append(collector, M[row][col])
			col--
		}
		row--
		col++
		//Column

		for row > i {
			collector = append(collector, M[row][col])
			row--
		}
		row++
		col--
	}

}
