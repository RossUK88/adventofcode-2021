package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetOverlapsFromFile(file string) int {
	vents, err := readVentsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateOverLaps(vents)
}

func CalculateOverLaps(vents []string) int {
	mappedVents := make(map[string]string)
	intersections := 0

	for _, vent := range vents {
		positions := strings.Split(vent, " -> ")
		from := positions[0]
		to := positions[1]

		linesFrom := strings.Split(from, ",")
		fromX := linesFrom[0]
		fromY := linesFrom[1]

		linesTo := strings.Split(to, ",")
		toX := linesTo[0]
		toY := linesTo[1]

		// We are going down

		if fromX == toX {
			fromYConv, _ := strconv.Atoi(fromY)
			toYConv, _ := strconv.Atoi(toY)

			f := fromYConv
			t := toYConv
			if f > t {
				f = t
				t = fromYConv
			}

			for i := f; i <= t; i++ {
				keyFromPosition := getMapKeyFromPosition(fromX, fmt.Sprintf("%d", i))

				if _, ok := mappedVents[keyFromPosition]; ok {
					strToInt, _ := strconv.Atoi(mappedVents[keyFromPosition])
					if strToInt == 1 {
						intersections++
					}
					ammIntersections := strToInt + 1

					mappedVents[keyFromPosition] = fmt.Sprintf("%d", ammIntersections)
				} else {
					mappedVents[keyFromPosition] = fmt.Sprintf("%d", 1)
				}
			}

		} else if fromY == toY {
			fromXConv, _ := strconv.Atoi(fromX)
			toXConv, _ := strconv.Atoi(toX)

			f := fromXConv
			t := toXConv
			if f > t {
				f = t
				t = fromXConv
			}

			for i := f; i <= t; i++ {
				keyFromPosition := getMapKeyFromPosition(fmt.Sprintf("%d", i), fromY)

				if _, ok := mappedVents[keyFromPosition]; ok {
					strToInt, _ := strconv.Atoi(mappedVents[keyFromPosition])
					if strToInt == 1 {
						intersections++
					}
					ammIntersections := strToInt + 1

					mappedVents[keyFromPosition] = fmt.Sprintf("%d", ammIntersections)

				} else {
					mappedVents[keyFromPosition] = fmt.Sprintf("%d", 1)
				}
			}
		} else {
			fmt.Println("Diag")
		}
	}

	return intersections
}

func readVentsFromFile(fileLocation string) ([]string, error) {

	var allInputs []string
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

	scanner := bufio.NewScanner(file)

	// 2d Array Inverter
	for scanner.Scan() {
		input := scanner.Text()

		allInputs = append(allInputs, input)
	}

	return allInputs, nil
}

func getMapKeyFromPosition(x string, y string) string {
	var positions []string
	positions = append(positions, x)
	positions = append(positions, y)
	return strings.Join(positions, "")
}
