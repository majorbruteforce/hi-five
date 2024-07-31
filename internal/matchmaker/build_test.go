package matchmaker

import (
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
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

	expected := []Edge{
		{Node1: 1, Node2: 2, Weight: 2},
		{Node1: 1, Node2: 3, Weight: 3},
		{Node1: 1, Node2: 4, Weight: 3},
		{Node1: 1, Node2: 5, Weight: 2},
		{Node1: 2, Node2: 4, Weight: 4},
		{Node1: 2, Node2: 5, Weight: 2},
		{Node1: 3, Node2: 4, Weight: 1},
		{Node1: 3, Node2: 5, Weight: 3},
		{Node1: 4, Node2: 5, Weight: 1},
	}

	result := Build(candidates)

	if len(result) != len(expected) {
		t.Errorf("expected length to be %d but got %d", len(expected), len(result))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("slices dont match:\n%v\n%v", result, expected)
	}

}
