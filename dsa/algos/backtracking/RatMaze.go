package backtracking

var mazX, MazY int

func RatMaze(input [][]int) {
	mazX = len(input) - 1
	MazY = len(input[0]) - 1
}

func moveRat(input [][]int, output [][]int, x, y int) bool {
	//base condition
	if x == mazX && y == MazY {
		return true
	}
	if input[x][x] == 1 { //Checking for current postion

		output[x][y] = 1
		if (x + 1) <= MazY {
			xm := moveRat(input, output, x+1, y)
			if xm {
				return true
			}
		}

		if (y + 1) <= mazX {
			ym := moveRat(input, output, x, y+1)
			if ym {
				return true
			}
		}
		output[x][y] = 0 //o
		return false
	}

	return false
}
