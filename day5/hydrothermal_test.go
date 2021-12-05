package day5

import "testing"

func TestGetOverlapsFromFile(t *testing.T) {
	actual := GetOverlapsFromFile("./sample_input.txt")
	expected := 5

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}
