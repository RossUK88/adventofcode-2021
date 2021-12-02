package day1

import "testing"

func TestIncreasedDepthCount(t *testing.T) {
	actual := IncreasedDepthCount([]int{1, 2, 40, 2, 5, 2, 1, 60})
	expected := 4

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}
