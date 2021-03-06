package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetIncreasedDepthCountFromFile(file string) int {
	inputsFromFile, err := readFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return IncreasedDepthCount(inputsFromFile)
}

func GetIncreasedSlidingDepthCountFromFile(file string) int {
	inputsFromFile, err := readFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return IncreasedSlidingDepthCount(inputsFromFile)
}

func IncreasedDepthCount(inputs []int) int {
	curr := 0
	incCount := 0
	for index, i := range inputs {
		if index == 0 {
			curr = i
			continue
		}

		if i > curr {
			incCount++
		}
		curr = i
	}

	return incCount
}

func IncreasedSlidingDepthCount(inputs []int) int {
	curr := 0
	incCount := 0

	// We need to start on atleast the 3rd index to get our first count
	prev := 0
	for i := 2; i < len(inputs); i++ {
		if i == 2 {
			prev = inputs[i-2] + inputs[i-1] + inputs[i]
			continue
		}

		curr = inputs[i-2] + inputs[i-1] + inputs[i]
		if curr > prev {
			incCount++
		}
		prev = curr
	}

	return incCount
}

func readFromFile(fileLocation string) ([]int, error) {
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
	inputs := []int{}
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return inputs, err
		}
		inputs = append(inputs, input)
	}

	return inputs, nil
}
