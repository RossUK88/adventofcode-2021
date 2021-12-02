package day2

import "testing"

func TestIncreasedDepthCount(t *testing.T) {
	// Horizontal 25
	// Vertical 14
	// Total 280
	actual := GetPositionOfSubmarine("./sample_input.txt")
	expected := 350

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}

func TestCalculateSubmarinePositionAndAim(t *testing.T) {
	//forward 9 // x: 9 y: 0 aim: 0
	//forward 8 // x: 17 y: 0 aim: 0
	//down 6 // x: 17 y: 0 aim: 6
	//down 9 // x: 17 y: 0 aim: 15
	//down 1 // x: 17 y: 0 aim: 16
	//up 2 // x: 17 y: 0 aim: 14
	//forward 3 // x: 20 y: 42 aim: 14
	//down 4 // x: 20 y: 42 aim: 18
	//down 3 // x: 20 y: 42 aim: 21
	//up 7 // x: 20 y: 42 aim: 14
	// forward 5 // x: 25 y: 112 aim: 14

	// Horizontal 25
	// Vertical 112
	// Total 2800
	actual := GetPositionAndAimOfSubmarine("./sample_input.txt")
	expected := 2800

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}
