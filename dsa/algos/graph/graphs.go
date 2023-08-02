package graph

import (
	"Gapp/dsa/ds/queue"
	"fmt"
)

type Vertex struct {
	Key           int
	AdjacencyList []*Vertex
}

type Graph struct {
	Vertices []*Vertex
	Size     int
}

func NewGraph(noOfVertices int) *Graph {
	v := make([]*Vertex, noOfVertices)
	return &Graph{Vertices: v, Size: noOfVertices}
}

func (graph *Graph) BFS() {
	var visited []bool
	visited = make([]bool, graph.Size)
	queue := queue.NewQueue()
	queue.Push(graph.Vertices[0])
	visited[0] = true

	for !queue.IsEmpty() {
		vertex, _ := queue.Pop().(*Vertex)
		fmt.Println(vertex.Key)
		for _, v := range vertex.AdjacencyList {
			if !visited[v.Key] {
				queue.Push(v)
			}
		}
	}
}

func (graph *Graph) DFS(vertex *Vertex, visited []bool) {
	fmt.Println(vertex.Key)
	for _, adj := range vertex.AdjacencyList {
		if !visited[adj.Key] {
			graph.DFS(adj, visited)
		}
	}
}
