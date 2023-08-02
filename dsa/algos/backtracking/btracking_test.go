package backtracking

import "testing"

func TestRatMaze(t *testing.T) {
	iniput := [][]int{{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1}}
	RatMaze(iniput)
}
