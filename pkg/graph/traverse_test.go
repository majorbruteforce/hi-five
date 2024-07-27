package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	g := New()
	one := g.AddVertex("1")
	two := g.AddVertex("2")
	three := g.AddVertex("3")
	four := g.AddVertex("4")
	five := g.AddVertex("5")
	six := g.AddVertex("6")

	g.AddEdge(one, two, 1)
	g.AddEdge(one, three, 1)
	g.AddEdge(two, four, 1)
	g.AddEdge(two, five, 1)
	g.AddEdge(three, six, 1)

	var result string

	g.BFS(one, func(v *Vertex) {
		result += v.ID
	})

	if result != "123456" {
		t.Errorf("Expected traversal result to be 123456; instead got %s\n", result)
	}

}

func TestDFS(t *testing.T) {
	g := New()
	one := g.AddVertex("1")
	two := g.AddVertex("2")
	three := g.AddVertex("3")
	four := g.AddVertex("4")
	five := g.AddVertex("5")
	six := g.AddVertex("6")

	g.AddEdge(one, two, 1)
	g.AddEdge(one, three, 1)
	g.AddEdge(two, four, 1)
	g.AddEdge(two, five, 1)
	g.AddEdge(three, six, 1)

	var result string

	g.DFS(one, func(v *Vertex) {
		result += v.ID
	})

	if result != "136254" {
		t.Errorf("Expected traversal result to be 124536; instead got %s\n", result)
	}

}
