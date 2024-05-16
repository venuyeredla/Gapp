package graph

import (
	"Gapp/dsa/ds/stackqueue"
	"fmt"
)

type Vertex struct {
	Key int
	Adj []int
}

type Graph struct {
	Vertices []*Vertex
	Size     int
}

func NewGraph(noOfVertices int) *Graph {
	v := make([]*Vertex, noOfVertices)
	return &Graph{Vertices: v, Size: noOfVertices}
}

func GraphWithEdges(noOfVertices int, directed bool, edges [][]int) *Graph {
	vertices := make([]*Vertex, noOfVertices)
	for i := range vertices {
		adj := make([]int, 0)
		vertices[i] = &Vertex{Key: i, Adj: adj}
	}
	for _, edge := range edges {
		vertices[edge[0]].Adj = append(vertices[edge[0]].Adj, edge[1])
		if !directed {
			vertices[edge[1]].Adj = append(vertices[edge[1]].Adj, edge[0])
		}
	}
	return &Graph{Vertices: vertices, Size: noOfVertices}
}

func (graph *Graph) BFS() {
	visited := make([]bool, graph.Size)
	queue := stackqueue.NewQueue()
	queue.Push(graph.Vertices[0])
	visited[0] = true
	for !queue.IsEmpty() {
		vertex, _ := queue.Pop().(*Vertex)
		fmt.Println(vertex.Key)
		for _, v := range vertex.Adj {
			if !visited[v] {
				queue.Push(v)
			}
		}
	}
}

func (graph *Graph) DFS(vertex *Vertex, visited []bool) {
	fmt.Println(vertex.Key)
	for _, adj := range vertex.Adj {
		if !visited[adj] {
			graph.DFS(graph.Vertices[adj], visited)
		}
	}
}

type MatrixGraph struct {
	MG   [][]int
	Size int
}

func MGraphWithEdges(noOfVertices int, directed bool, edges [][]int) *MatrixGraph {
	graph := make([][]int, noOfVertices)
	for i := 0; i < noOfVertices; i++ {
		graph[i] = make([]int, noOfVertices)
	}
	for _, edge := range edges {
		graph[edge[0]][edge[1]] = 1

		if !directed {
			graph[edge[1]][edge[0]] = 1
		}
	}
	return &MatrixGraph{MG: graph, Size: noOfVertices}
}

func (mgraph *MatrixGraph) DFS() {
	visited := make([]bool, mgraph.Size)
	MDFS(0, visited, mgraph.MG)
}

func MDFS(vertex int, visited []bool, graph [][]int) {
	fmt.Printf("Start =%v", vertex)
	// Set current node as visited
	visited[vertex] = true

	// For every node of the graph
	for i := 0; i < len(graph[vertex]); i++ {
		// If some node is adjacent to the current node
		// and it has not already been visited
		if graph[vertex][i] == 1 && !visited[i] {
			MDFS(i, visited, graph)
		}
	}
}

func MBFS(start int, graph [][]int) {
	visited := make([]bool, len(graph))
	queue := stackqueue.NewQueue()
	queue.Push(start)
	visited[start] = true
	for !queue.IsEmpty() {
		vertex, _ := queue.Pop().(int)
		fmt.Println(vertex)
		for i := 0; i < len(graph); i++ {
			if graph[vertex][i] == 1 && !visited[vertex] {
				queue.Push(i)
				visited[i] = true
			}
		}
	}
}
