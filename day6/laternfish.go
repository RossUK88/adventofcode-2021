package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetTotalLanternfishFromFile(file string, days int) int {
	fish, err := readFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	return CalculateLanternFish(fish, days)
}

func CalculateLanternFish(fish map[int]int, days int) int {
	// TODO we need to keep the next count rather than the old
	// next[8] += fish[0]
	// next[6] += fish[0]
	//	 next[j] += fish[j+1]
	// fish = next

	oldFish := fish
	for i := 1; i <= days; i++ {
		if i <= 4 {
			fmt.Println(fish, oldFish)
		}

		for j := 0; j <= 8; j++ {
			if j == 0 {
				newFish := fish[j]

				// TODO WHY YOU NO PUT FISHY IN 6?!?!
				if _, ok := fish[6]; ok {
					fmt.Println(fish[6])
					fish[6] += newFish
				} else {
					fish[6] = newFish
				}

				fish[8] = newFish
			} else {
				fish[j-1] = oldFish[j]
			}
		}

		oldFish = fish
	}

	totalFish := 0
	for j := 8; j >= 0; j-- {
		totalFish += fish[j]
	}

	fmt.Println(fish)
	return totalFish
}

func readFromFile(filename string) (map[int]int, error) {
	fish := make(map[int]int)

	file, err := os.Open(filename)
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
		for _, fishy := range strings.Split(input, ",") {
			life, err := strconv.Atoi(fishy)
			if err != nil {
				return fish, err
			}

			if _, ok := fish[life]; ok {
				fish[life] += 1
			} else {
				fish[life] = 1
			}
		}
	}

	return fish, nil
}
