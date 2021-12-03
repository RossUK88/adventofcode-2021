package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetSubmarineLifeSupportRating(file string) int64 {
	allInputs, inputsFromFile, err := readPositionsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateLifeSupportRating(allInputs, inputsFromFile)
}

func CalculateLifeSupportRating(allInputs []string, diagnostics [][]int) int64 {
	oxygenRating := getOxygenRating(allInputs, diagnostics)
	co2Rating := getCo2Rating(allInputs, diagnostics)

	return oxygenRating * co2Rating
}

func getOxygenRating(allInputs []string, initialValues [][]int) int64 {
	diagnostics := allInputs
	values := initialValues

	// TODO - We are using this wrong
	valueLookingFor := mostCommonBinaryInSlice(values[0], true)
	currentIndex := 0
	// Keep looping until only one is left
	for len(diagnostics) != 1 {
		var tempValues []string
		for _, diag := range diagnostics {
			char := rune(diag[currentIndex])
			intRep, _ := strconv.Atoi(string(char))
			if intRep == valueLookingFor {
				tempValues = append(tempValues, diag)
			}
		}

		currentIndex++
		// Check new thing looking for
		diagnostics = tempValues
		//fmt.Println(diagnostics)
		values = flipStringsToIntSlice(diagnostics)
		valueLookingFor = mostCommonBinaryInSlice(values[0], true)
		fmt.Println(values, valueLookingFor)
	}

	// Do something with values
	oxygenRatingString := diagnostics[0]

	oxygenRating, _ := strconv.ParseInt(oxygenRatingString, 2, 16)

	return oxygenRating
}

func getCo2Rating(allInputs []string, initialValues [][]int) int64 {
	diagnostics := allInputs
	values := initialValues

	mostCommon := mostCommonBinaryInSlice(values[0], false)
	valueLookingFor := 1
	if mostCommon == 1 {
		valueLookingFor = 0
	}

	currentIndex := 0
	// Keep looping until only one is left
	for len(diagnostics) != 1 {
		var tempValues []string
		for _, diag := range diagnostics {
			char := rune(diag[currentIndex])
			intRep, _ := strconv.Atoi(string(char))
			if intRep == valueLookingFor {
				tempValues = append(tempValues, diag)
			}
		}

		currentIndex++
		// Check new thing looking for
		diagnostics = tempValues

		values = flipStringsToIntSlice(diagnostics)
		mostCommon = mostCommonBinaryInSlice(values[currentIndex-1], false)
		if mostCommon == 1 {
			valueLookingFor = 0
		} else {
			valueLookingFor = 1
		}

	}

	// Do something with values
	oxygenRatingString := diagnostics[0]

	oxygenRating, _ := strconv.ParseInt(oxygenRatingString, 2, 16)

	return oxygenRating
}

func flipStringsToIntSlice(inputs []string) [][]int {
	var results [][]int
	row := 0

	// 2d Array Inverter
	for _, input := range inputs {

		for charIdx, char := range input {
			// This is the first Character in each row, so we need a new row in our inverted
			if len(results) < charIdx+1 {
				results = append(results, []int{})
			}

			// Grab the current set of Diagnostics from our invert
			currentDiagnostics := results[charIdx]

			// Convert the Rune to an Int
			runeToInt, _ := strconv.Atoi(string(char))

			// Append the binary number and add back in at the index
			currentDiagnostics = append(currentDiagnostics, runeToInt)
			results[charIdx] = currentDiagnostics
		}

		row++
	}

	return results
}

func GetSubmarinePowerConsumption(file string) int64 {
	_, inputsFromFile, err := readPositionsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateSubmarinePowerConsumption(inputsFromFile)
}

func CalculateSubmarinePowerConsumption(diagnostics [][]int) int64 {
	var gamma []int
	var epislon []int

	for _, diagnostic := range diagnostics {
		gamma = append(gamma, mostCommonBinaryInSlice(diagnostic, false))
		epsilonValue := 1
		if mostCommonBinaryInSlice(diagnostic, false) == 1 {
			epsilonValue = 0
		}

		epislon = append(epislon, epsilonValue)

	}

	var gammaString string
	var epsilonString string
	for _, bit := range gamma {
		gammaString = fmt.Sprintf("%s%d", gammaString, bit)
	}
	for _, bit := range epislon {
		epsilonString = fmt.Sprintf("%s%d", epsilonString, bit)
	}

	gammaDecimal, _ := strconv.ParseInt(gammaString, 2, 16)
	epislonDecimal, _ := strconv.ParseInt(epsilonString, 2, 16)

	return gammaDecimal * epislonDecimal
}

func mostCommonBinaryInSlice(binarySlice []int, roundDown bool) int {
	on := 0
	off := 0

	for _, binary := range binarySlice {
		if binary == 1 {
			on += 1
		} else {
			off += 1
		}
	}

	if roundDown {
		if on > off {
			return 1
		}
	} else {
		if on >= off {
			return 1
		}
	}

	return 0
}

func readPositionsFromFile(fileLocation string) ([]string, [][]int, error) {
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

		allInputs = append(allInputs, input)

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

	return allInputs, diagnostics, nil
}
