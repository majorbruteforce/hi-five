package main

import (
	"github.com/majorbruteforce/hi-five/pkg/graph"
)

func main() {
	g := graph.NewGraph()
	candidates := []*graph.Candidate{
		{ID: "1", Keywords: []string{"go", "python", "java"}},
		{ID: "2", Keywords: []string{"java", "c++", "python"}},
		{ID: "3", Keywords: []string{"go", "rust", "c"}},
		{ID: "4", Keywords: []string{"rust", "python", "java"}},
	}
	g.Build(candidates)
	g.PrintGraph()
}
