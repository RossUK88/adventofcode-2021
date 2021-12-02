package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"fmt"
)

func main() {
	// Day 1
	fmt.Printf("The increase in depth for Day 1 is : %d\n", day1.GetIncreasedDepthCountFromFile("./day1/input.txt"))
	fmt.Printf("The increase in sliding depth for Day 1 is : %d\n", day1.GetIncreasedSlidingDepthCountFromFile("./day1/input.txt"))

	// Day 2
	fmt.Printf("The position of the submarine in Day 2 is : %d\n", day2.GetPositionOfSubmarine("./day2/input.txt"))
}
