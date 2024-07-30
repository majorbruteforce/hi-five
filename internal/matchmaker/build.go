package matchmaker

import "github.com/majorbruteforce/hi-five/pkg/graph"

type Edge = graph.Edge

// Build constructs the graph from a list of candidates.
// Each candidate is added as a vertex, and edges are added between vertices
// based on the number of common keywords. The weight of each edge is the count
// of common keywords between the connected candidates.
func Build(candidateList []*Candidate) []Edge {

	// edges stores the matched edges of the graph
	var edges []Edge

	for i := 0; i < len(candidateList); i++ {
		for j := i + 1; j < len(candidateList); j++ {
			weight := commonKeywordsCount(candidateList[i].Keywords, candidateList[j].Keywords)
			if weight > 0 {

				e := Edge{
					Node1:  i + 1,
					Node2:  j + 1,
					Weight: float64(weight),
				}
				edges = append(edges, e)
			}
		}
	}

	return edges
}

// commonKeywordsCount calculates the number of common keywords between two lists of keywords.
// It returns the count of common keywords.
func commonKeywordsCount(k1, k2 []string) int {
	keywordSet := make(map[string]struct{})
	var count int
	for _, k := range k1 {
		keywordSet[k] = struct{}{}
	}
	for _, k := range k2 {
		if _, exists := keywordSet[k]; exists {
			count++
		}
	}
	return count
}
