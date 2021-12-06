package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"adventofcode/day3"
	"adventofcode/day4"
	"adventofcode/day5"
	"adventofcode/day6"
	"fmt"
	"log"
	"time"
)

func main() {
	// Day 1
	fmt.Printf("The increase in depth for Day 1 is : %d\n", day1.GetIncreasedDepthCountFromFile("./day1/input.txt"))
	fmt.Printf("The increase in sliding depth for Day 1 is : %d\n", day1.GetIncreasedSlidingDepthCountFromFile("./day1/input.txt"))

	// Day 2
	fmt.Printf("The position of the submarine in Day 2 is : %d\n", day2.GetPositionOfSubmarine("./day2/input.txt"))
	fmt.Printf("The position of the submarine in Day 2 after aiming is : %d\n", day2.GetPositionAndAimOfSubmarine("./day2/input.txt"))

	// Day 3
	fmt.Printf("The power consumption of the submarine in Day 3 is : %d\n", day3.GetSubmarinePowerConsumption("./day3/input.txt"))
	fmt.Printf("The life suport rating of the submarine in Day 3 is : %d\n", day3.GetSubmarineLifeSupportRating("./day3/input.txt"))

	// Day 4
	fmt.Printf("The final score for Day 4 is : %d\n", day4.GetFinalBingoScoreFromFile("./day4/input.txt"))
	fmt.Printf("The final score of the last winning board in Day 4 is : %d\n", day4.GetFinalBingoScoreFromLastToWinFromFile("./day4/input.txt"))

	// Day 4
	start := time.Now()

	fmt.Printf("The intersections for Day 5 is : %d\n", day5.GetOverlapsFromFile("./day5/input.txt"))
	//fmt.Printf("The intersections for Day 5 is : %d\n", day5.GetOverlapsFromFile("./day5/sample_input.txt"))
	//fmt.Printf("The final score of the last winning board in Day 4 is : %d\n", day4.GetFinalBingoScoreFromLastToWinFromFile("./day4/input.txt"))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	start = time.Now()
	fmt.Printf("The total lantern fish after day 80 for Day 6 is : %d\n", day6.GetTotalLanternfishFromFile("./day6/input.txt", 80))
	elapsed = time.Since(start)
	log.Printf("took %s", elapsed)

	start = time.Now()
	fmt.Printf("The total lantern fish after day 256 for Day 6 is : %d\n", day6.GetTotalLanternfishFromFile("./day6/input.txt", 256))
	elapsed = time.Since(start)
	log.Printf("took %s", elapsed)

}
