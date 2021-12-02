package day1

import "testing"

func TestIncreasedDepthCount(t *testing.T) {
	actual := IncreasedDepthCount([]int{1, 2, 40, 2, 5, 2, 1, 60})
	expected := 4

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}

func TestIncreasedSlidingDepthCount(t *testing.T) {
	// [1]
	// [1, 2]
	// [1, 2, 40]
	// [2, 40, 2] // 1 - Inc
	// [40, 2, 5] // 2 - Inc
	// [2, 5, 2]  // 2 - Dec
	// [5, 2, 1] // 2 - Dec
	// [2, 1, 60] // 3 - Inc
	// [1, 60, 40] // 4 - Inc
	// [60, 40, 60] // 5 - Inc
	// [40, 60, 12] // 5 - Dec
	// [60, 12, 32] // 5 - Dec
	actual := IncreasedSlidingDepthCount([]int{1, 2, 40, 2, 5, 2, 1, 60, 40, 60, 12, 32})
	expected := 5

	if actual != expected {
		t.Errorf("Actual: %d, Expected: %d", actual, expected)
	}

}
