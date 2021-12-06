package day6

import "testing"

func TestGetTotalLanterfishFromFile(t *testing.T) {
	actual := GetTotalLanternfishFromFile("./sample_input.txt", 80)
	expected := 5934

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}
