// pkg/graph/matching.go

package graph

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
)

type Edge struct {
	Node1  int     `json:"node1"`
	Node2  int     `json:"node2"`
	Weight float64 `json:"weight"`
}

func MaxWeightedMatching(edges []Edge) ([][2]int, error) {
	edgesJSON, err := json.Marshal(edges)
	if err != nil {
		return nil, err
	}

	scriptPath, err := filepath.Abs(filepath.Join("scripts", "max_weighted_matching.py"))
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("python3", scriptPath, string(edgesJSON))
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
