package day3

import "testing"

func TestCalculateSubmarinePowerConsumption(t *testing.T) {
	// 1010 0000 1100
	// 0111 1110 0111
	// 1111 0000 1110
	// 1100 0001 1001
	// 0010 0100 1011

	// this should turn into
	// 10110 - 1
	// 01110 - 1
	// 11100 - 1
	// 01100 - 0
	// 01000 - 0
	// 01001 - 0
	// 01000 - 0
	// 00010 - 0
	// 10111 - 1
	// 11100 - 1
	// 01101 - 1
	// 01011 - 1

	// Gamma - 11100001111 - 3599
	// Epsilon - 00011110000 - 496
	// Power = 1785104

	actual := GetSubmarinePowerConsumption("./sample_input.txt")
	expected := int64(1785104)

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}

func TestGetSubmarineLifeSupportRating(t *testing.T) {
	// 1010 0000 1100
	// 0111 1110 0111
	// 1111 0000 1110
	// 1100 0001 1001
	// 0010 0100 1011

	// Oxygen Rating
	// - 1st pass
	// 1010 0000 1100
	// 1111 0000 1110
	// 1100 0001 1001
	// - 2nd Pass
	// 1111 0000 1110
	// 1100 0001 1001
	// - 3rd Pass
	// 1111 0000 1110
	// Value = 3854

	// C02 Rating
	// - 1st pass
	// 0111 1110 0111
	// 0010 0100 1011
	// - 2nd pass
	// 0010 0100 1011
	// Value = 587

	// Life Support = 2103808

	actual := GetSubmarineLifeSupportRating("./sample_input.txt")
	expected := int64(2103808)

	if actual != expected {
		t.Errorf("Actual %d, Expected: %d", actual, expected)
	}
}
