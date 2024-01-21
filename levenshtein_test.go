package diff

import (
	"reflect"
	"testing"
)

func TestLevenshteinEditOps(t *testing.T) {
	result := LevenshteinEditOps("Come on", "Came on")
	expected := []EditOp{
		{
			Action: 'R',
			Src:    "o",
			Dst:    "a",
			SrcPos: 1,
			DstPos: 1,
		},
	}

	if len(result) != 1 || !reflect.DeepEqual(result[0], expected[0]) {
		t.Errorf("Expected Levenshtein edit ops %+v, but got %+v", expected, result)
	}
}

func TestLevenshteinDistance(t *testing.T) {
	result := LevenshteinDistance("Come on", "Came on")
	expected := 1

	if result != expected {
		t.Errorf("Expected Levenshtein distance of %d, but got %d", expected, result)
	}
}
