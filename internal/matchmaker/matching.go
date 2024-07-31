package matchmaker

import (
	"fmt"

	"github.com/majorbruteforce/hi-five/pkg/graph"
)

func MatchCandidatesBatch(candidates []Candidate) [][2]string {
	matches := make([][2]string, 0)

	edges := Build(candidates)
	results, err := graph.MaxWeightedMatching(edges)

	if err != nil {
		fmt.Printf("error matching candidates: %v", err)
		return nil
	}
	fmt.Println(results)
	for _, r := range results {
		match := [2]string{candidates[r[0]-1].ID, candidates[r[1]-1].ID}
		matches = append(matches, match)
	}
	return matches
}
