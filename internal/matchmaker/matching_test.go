package matchmaker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatchCandidatesBatch(t *testing.T) {
	candidates := []Candidate{
		{
			ID:       "candidate-1",
			Keywords: []string{"golang", "python", "java", "javascript", "c++"},
		},
		{
			ID:       "candidate-2",
			Keywords: []string{"ruby", "swift", "kotlin", "java", "golang"},
		},
		{
			ID:       "candidate-3",
			Keywords: []string{"php", "c++", "scala", "python", "javascript"},
		},
		{
			ID:       "candidate-4",
			Keywords: []string{"swift", "java", "python", "golang", "ruby"},
		},
		{
			ID:       "candidate-5",
			Keywords: []string{"kotlin", "javascript", "php", "c++", "ruby"},
		},
	}
	expected := [][2]string{
		{"candidate-1", "candidate-3"},
		{"candidate-4", "candidate-2"},
	}

	result := MatchCandidatesBatch(candidates)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v\n got: %v\n", expected, result)
	}

	fmt.Println(result)
}
