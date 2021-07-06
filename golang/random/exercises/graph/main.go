package main

import (
	"fmt"
)

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int) error {
	if contains(g.vertices, k) {
		return fmt.Errorf("Vertex %v not added because it already exists", k)
	}
	g.vertices = append(g.vertices, &Vertex{key: k})

	return nil
}

// AddEdge adds an endge to the graph
func (g *Graph) AddEdge(from, to int) error {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		return fmt.Errorf("Invalid edge: (%v->%v)", from, to)
	}
	// check if edge already exists
	if contains(fromVertex.adjacent, to) {
		return fmt.Errorf("Edge already exists: (%v->%v)", from, to)
	}
	// add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	return nil
}

// getVertex returns a pointer to the Vertex wsith a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex on the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v:", v.key)
		for _, y := range v.adjacent {
			fmt.Printf(" %v", y.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	err := test.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = test.AddEdge(6, 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = test.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	test.Print()
}
