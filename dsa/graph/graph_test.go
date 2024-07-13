package graph

import "testing"

func TestGraphBFS(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	GraphWithEdges(3, false, edges)
}

func TestGraphDFS(t *testing.T) {

}
