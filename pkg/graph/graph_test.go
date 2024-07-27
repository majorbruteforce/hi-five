package graph

import (
	"testing"
)

func TestBuild(t *testing.T) {
	g := NewGraph()
	candidates := []*Candidate{
		{ID: "A", Keywords: []string{"Go", "Graph"}},
		{ID: "B", Keywords: []string{"Go", "Programming"}},
		{ID: "C", Keywords: []string{"Graph", "Algorithms"}},
	}
	g.Build(candidates)
	if len(g.Vertices) != 3 {
		t.Errorf("Build() should add 3 vertices, got %d", len(g.Vertices))
	}
	if g.EdgeCount() != 4 {
		t.Errorf("Build() should add 4 edges, got %d", len(g.Edges))
	}
	if g.Edges[g.Vertices[1]][g.Vertices[2]] != nil {
		t.Error("Build() should not create edge between 'B' and 'C', edge was created.")
	}
	if weight := g.Edges[g.Vertices[0]][g.Vertices[1]].Weight; weight != 1 {
		t.Errorf("Build() should set edge weight between 'A' and 'B' to 1, got %d", weight)
	}
	if weight := g.Edges[g.Vertices[0]][g.Vertices[2]].Weight; weight != 1 {
		t.Errorf("Build() should set edge weight between 'A' and 'C' to 1, got %d", weight)
	}

}
