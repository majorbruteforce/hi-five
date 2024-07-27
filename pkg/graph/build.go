package graph

// Candidate represents a candidate with an ID and a list of keywords.
type Candidate struct {
	ID       string
	Keywords []string
}

// Build constructs the graph from a list of candidates.
// Each candidate is added as a vertex, and edges are added between vertices
// based on the number of common keywords. The weight of each edge is the count
// of common keywords between the connected candidates.
func (g *Graph) Build(candidateList []*Candidate) {
	var addedVertices []*Vertex

	for _, c := range candidateList {
		addedVertices = append(addedVertices, g.AddVertex(c.ID))
	}
	for i := 0; i < len(candidateList); i++ {
		for j := i + 1; j < len(candidateList); j++ {
			weight := commonKeywords(candidateList[i].Keywords, candidateList[j].Keywords)
			if weight > 0 {
				g.AddEdge(addedVertices[i], addedVertices[j], weight)
			}
		}
	}
}

// commonKeywords calculates the number of common keywords between two lists of keywords.
// It returns the count of common keywords.
func commonKeywords(k1, k2 []string) int {
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
