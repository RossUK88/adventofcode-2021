package day4

import "testing"

func TestGetFinalBingoScoreFromFile(t *testing.T) {

	actual := GetFinalBingoScoreFromFile("./sample_input.txt")
	expected := 4512

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}
