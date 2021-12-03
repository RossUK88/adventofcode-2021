package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetSubmarinePowerConsumption(file string) int {
	inputsFromFile, err := readPositionsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateSubmarinePowerConsumption(inputsFromFile)
}

func CalculateSubmarinePowerConsumption([][]int) int {

	return 1
}

func readPositionsFromFile(fileLocation string) ([][]int, error) {
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

	// 0101
	// 1011
	// 0001

	// Position 1 goes to [0][]int{0,1,0}
	// Position 2 goes to [1][]int{1,0,0}
	// Position 3 goes to [2][]int{0,1,0}
	// Position 4 goes to [3][]int{1,1,1}

	var diagnostics [][]int
	scanner := bufio.NewScanner(file)
	row := 0

	// 2d Array Inverter
	for scanner.Scan() {
		input := scanner.Text()

		for charIdx, char := range input {
			// This is the first Character in each row, so we need a new row in our inverted
			if len(diagnostics) < charIdx+1 {
				diagnostics = append(diagnostics, []int{})
			}

			// Grab the current set of Diagnostics from our invert
			currentDiagnostics := diagnostics[charIdx]

			// Convert the Rune to an Int
			runeToInt, _ := strconv.Atoi(string(char))

			// Append the binary number and add back in at the index
			currentDiagnostics = append(currentDiagnostics, runeToInt)
			diagnostics[charIdx] = currentDiagnostics
		}

		row++
	}

	fmt.Println(diagnostics)

	return diagnostics, nil
}
