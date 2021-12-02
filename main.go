package main

import (
	"adventofcode/day1"
	"fmt"
)

func main() {
	fmt.Printf("The increase in depth for Day 1 is : %d\n", day1.GetIncreasedDepthCountFromFile("./day1/input.txt"))
	fmt.Printf("The increase in sliding depth for Day 1 is : %d\n", day1.GetIncreasedSlidingDepthCountFromFile("./day1/input.txt"))
}
