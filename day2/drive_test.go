package day2

import "testing"

func TestIncreasedDepthCount(t *testing.T) {
	// Horizontal 20
	// Vertical 14
	// Total 280
	actual := GetPositionOfSubmarine("./sample_input.txt")
	expected := 280

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}
