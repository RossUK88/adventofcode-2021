package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetOverlapsFromFile(file string) int {
	vents, err := readVentsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateOverLaps(vents)
}

func CalculateOverLaps(vents []string) int {

	fmt.Println(vents)

	return 1
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
