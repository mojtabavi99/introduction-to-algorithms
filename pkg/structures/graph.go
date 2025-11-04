package structures

import (
	"fmt"
	"slices"
)

// Graph represents a simple directed or undirected graph
// using an adjacency list representation.
type Graph struct {
	Adjacency map[string][]string // adjacency list: vertex → connected vertices
	Directed  bool                // whether the graph is directed
}

// initializes and returns a new graph.
func GraphInit(directed bool) *Graph {
	return &Graph{
		Adjacency: make(map[string][]string),
		Directed:  directed,
	}
}

// AddVertex adds a new vertex to the graph if it doesn't already exist.
func (graph *Graph) AddVertex(vertex string) {
	// map lookup
	if _, exists := graph.Adjacency[vertex]; !exists {
		graph.Adjacency[vertex] = []string{}
	}
}

// AddEdge connects from → to (and optionally to → from if undirected).
func (graph *Graph) AddEdge(from, to string) {
	graph.AddVertex(from)
	graph.AddVertex(to)

	// Avoid duplicate edge
	if !slices.Contains(graph.Adjacency[from], to) {
		graph.Adjacency[from] = append(graph.Adjacency[from], to)
	}

	// If undirected, add reverse edge
	if !graph.Directed && !slices.Contains(graph.Adjacency[to], from) {
		graph.Adjacency[to] = append(graph.Adjacency[to], from)
	}
}

// RemoveEdge removes an edge between two vertices.
func (graph *Graph) RemoveEdge(from, to string) {
	if graph.Adjacency == nil {
		return
	}

	graph.Adjacency[from] = removeFromSlice(graph.Adjacency[from], to)
	if !graph.Directed {
		graph.Adjacency[to] = removeFromSlice(graph.Adjacency[to], from)
	}
}

// RemoveVertex deletes a vertex and all edges associated with it.
func (graph *Graph) RemoveVertex(v string) {
	if graph.Adjacency == nil {
		return
	}

	delete(graph.Adjacency, v)
	for key, neighbors := range graph.Adjacency {
		graph.Adjacency[key] = removeFromSlice(neighbors, v)
	}
}

// HasEdge checks whether an edge exists between two vertices.
func (graph *Graph) HasEdge(from, to string) bool {
	if graph.Adjacency == nil {
		return false
	}

	neighbors, exists := graph.Adjacency[from]
	if !exists {
		return false
	}

	return slices.Contains(neighbors, to)
}

// PrettyPrint prints the adjacency list of the graph in a readable format.
func (graph *Graph) PrettyPrint() {
	if len(graph.Adjacency) == 0 {
		fmt.Println("Graph is empty.")
		return
	}

	fmt.Println("Graph Adjacency List:")
	for vertex, neighbors := range graph.Adjacency {
		fmt.Printf("%s → %v\n", vertex, neighbors)
	}
}

// removeFromSlice removes a target string from a slice (if it exists).
func removeFromSlice(slice []string, target string) []string {
	for i, v := range slice {
		if v == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
