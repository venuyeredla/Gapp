package backtracking

import "fmt"

var N int = 9

/*
Takes a partially filled-in grid and attempts

	to assign values to all unassigned locations in
	such a way to meet the requirements for
	Sudoku solution (non-duplication across rows,
	columns, and boxes)
*/
func solveSudoku(grid [][]int, row, col int) bool {

	/*if we have reached the 8th
	row and 9th column (0
	indexed matrix) ,
	we are returning true to avoid further
	backtracking       */
	if row == N-1 && col == N {
		return true
	}

	// Check if column value  becomes 9 ,
	// we move to next row
	// and column start from 0
	if col == N {
		row++
		col = 0
	}

	// Check if the current position
	// of the grid already
	// contains value >0, we iterate
	// for next column
	if grid[row][col] != 0 {
		return solveSudoku(grid, row, col+1)
	}

	for num := 1; num < 10; num++ {

		// Check if it is safe to place
		// the num (1-9)  in the
		// given row ,col ->we move to next column
		if isSafe(grid, row, col, num) {

			/*
			   	assigning the num in the current

			   (row,col)  position of the grid and
			   assuming our assigned num in the position
			   is correct
			*/
			grid[row][col] = num

			// Checking for next
			// possibility with next column
			if solveSudoku(grid, row, col+1) {
				return true
			}

		}
		/*
		   	removing the assigned num , since our

		   assumption was wrong , and we go for next
		   assumption with diff num value
		*/
		grid[row][col] = 0
	}
	return false
}

/* A utility function to print grid */
func print(grid [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%v", grid[i][j])
			fmt.Println()
		}
	}
}

// Check whether it will be legal
// to assign num to the
// given row, col
func isSafe(grid [][]int, row, col, num int) bool {
	// Check if we find the same num
	// in the similar row , we
	// return false
	for x := 0; x <= 8; x++ {
		if grid[row][x] == num {
			return false
		}

	}

	// Check if we find the same num
	// in the similar column ,
	// we return false
	for x := 0; x <= 8; x++ {
		if grid[x][col] == num {
			return false
		}
	}

	// Check if we find the same num
	// in the particular 3*3
	// matrix, we return false
	var startRow int = row - row%3
	var startCol = col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}

		}

	}

	return true
}
