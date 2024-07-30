// pkg/graph/matching.go

package graph

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type Edge struct {
	Node1  int     `json:"node1"`
	Node2  int     `json:"node2"`
	Weight float64 `json:"weight"`
}

func MaxWeightedMatching(edges []Edge) ([][2]int, error) {
	edgeArg := formatEdges(edges)

	scriptPath, err := filepath.Abs(filepath.Join("scripts", "max_weighted_matching.py"))
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("/usr/bin/python3", scriptPath, edgeArg)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result [][2]int
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func formatEdges(edges []Edge) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, edge := range edges {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("[%d, %d, %.1f]", edge.Node1, edge.Node2, edge.Weight))
	}
	sb.WriteString("]")
	return sb.String()
}
