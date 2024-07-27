package graph

import "fmt"

// Vertex represents a single vertex in the graph.
type Vertex struct {
	ID string
}

// Edge represents an edge in the graph with a weight.
type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight int
}

// Graph represents a graph with vertices and edges.
type Graph struct {
	Vertices []*Vertex
	Edges    map[*Vertex]map[*Vertex]*Edge
}

// NewGraph creates and returns a new instance of Graph.
func NewGraph() *Graph {
	return &Graph{
		Vertices: make([]*Vertex, 0),
		Edges:    make(map[*Vertex]map[*Vertex]*Edge),
	}
}

// AddVertex adds a vertex with the given ID to the graph.
// It returns the added vertex.
func (g *Graph) AddVertex(id string) *Vertex {
	v := &Vertex{ID: id}
	g.Vertices = append(g.Vertices, v)
	return v
}

// AddEdge adds an edge between two vertices with a specified weight.
// The edge is added in both directions since the graph is undirected.
func (g *Graph) AddEdge(from, to *Vertex, weight int) {
	if g.Edges[from] == nil {
		g.Edges[from] = make(map[*Vertex]*Edge)
	}
	g.Edges[from][to] = &Edge{From: from, To: to, Weight: weight}

	if g.Edges[to] == nil {
		g.Edges[to] = make(map[*Vertex]*Edge)
	}
	g.Edges[to][from] = &Edge{From: to, To: from, Weight: weight}
}

// PrintGraph prints the vertices and edges of the graph.
// Vertices are listed first, followed by the edges with their weights.
func (g *Graph) PrintGraph() {
	fmt.Println("Vertices:")
	for _, v := range g.Vertices {
		fmt.Println(" -", v.ID)
	}

	fmt.Println("Edges:")
	for from, edges := range g.Edges {
		for to, edge := range edges {
			fmt.Printf(" - %s -> %s (weight: %d)\n", from.ID, to.ID, edge.Weight)
		}
	}
}
