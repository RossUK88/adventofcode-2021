package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetPositionOfSubmarine(file string) int {
	inputsFromFile, err := readPositionsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateSubmarinePosition(inputsFromFile)
}

func CalculateSubmarinePosition(positions []string) int {
	grid := map[string]int{
		"x": 0,
		"y": 0,
	}

	for _, position := range positions {
		if isHorizontal(position) {
			grid["x"] += getMovementAmount(position, false)
		} else {
			grid["y"] += getMovementAmount(position, false)
		}
	}

	return grid["x"] * grid["y"]
}

func GetPositionAndAimOfSubmarine(file string) int {
	inputsFromFile, err := readPositionsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateSubmarinePositionAndAim(inputsFromFile)
}

func CalculateSubmarinePositionAndAim(positions []string) int {
	grid := map[string]int{
		"x":   0,
		"y":   0,
		"aim": 0,
	}

	for _, position := range positions {
		if isHorizontal(position) {
			//forward X does two things:
			//It increases your horizontal position by X units.
			//	It increases your depth by your aim multiplied by X.
			movement := getMovementAmount(position, true)
			grid["x"] += movement
			grid["y"] += grid["aim"] * movement
		} else {
			grid["aim"] += getMovementAmount(position, false)
		}
	}

	return grid["x"] * grid["y"]
}

func isHorizontal(positionalData string) bool {
	return strings.HasPrefix(positionalData, "forward")
}

func getMovementAmount(positionalData string, forcePositive bool) int {
	// Split on spaces
	splits := strings.SplitAfter(positionalData, " ")

	// Get the end of the array which _should_ always be the int we need
	movement, err := strconv.Atoi(splits[len(splits)-1])
	if err != nil {
		log.Fatal("Unable to get positional movement data")
	}

	// If we are going up then we need to get the negative
	if !forcePositive {
		if strings.Contains(positionalData, "up") {
			movement *= -1
		}
	}

	return movement
}

func readPositionsFromFile(fileLocation string) ([]string, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Unable to close file, %s", err)
		}
	}(file)

	positions := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		positions = append(positions, input)
	}

	return positions, nil
}
