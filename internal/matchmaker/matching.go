package matchmaker

import (
	"fmt"

	"github.com/majorbruteforce/hi-five/pkg/graph"
)

func matchCandidatesBatch(candidates []*Candidate) [][2]*Candidate {
	matches := make([][2]*Candidate, 0)

	edges := Build(candidates)
	results, err := graph.MaxWeightedMatching(edges)

	if err != nil {
		fmt.Printf("error matching candidates: %v", err)
		return nil
	}

	for _, r := range results {
		match := [2]*Candidate{candidates[r[0]], candidates[r[1]]}
		matches = append(matches, match)
	}
	return matches
}
