package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoCard struct {
	numbers []int
	board   [5][]Position
}
type Position struct {
	value  int
	marked bool
}

func (b *BingoCard) buildBoard() {
	fmt.Println("Building Board")
	row := 0
	cursor := 0
	for cursor+5 <= len(b.numbers) {
		numbers := b.numbers[cursor : cursor+5]

		var positions []Position

		for _, number := range numbers {
			positions = append(positions, Position{
				value:  number,
				marked: false,
			})
		}

		b.board[row] = positions

		row++
		cursor += 5
	}
}

func GetFinalBingoScoreFromFile(file string) int {
	bingoNumbers, cards, err := readBingoCardsFromFile(file)
	if err != nil {
		log.Fatalf("Unable to read inputs from file, %s", err)
	}

	fmt.Println(cards[0].board)

	return CalculateFinalScore(bingoNumbers, cards)
}

func CalculateFinalScore(numbers []int, cards []BingoCard) int {

	return 1
}

func readBingoCardsFromFile(fileLocation string) ([]int, []BingoCard, error) {
	var bingoNumbers []int
	var cards []BingoCard

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
	row := 0
	var allNumbers []int

	for scanner.Scan() {
		input := scanner.Text()
		// Bingo Numbers
		if row == 0 {
			numbers := strings.Split(input, ",")
			for _, number := range numbers {
				conv, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal("Unable to convert bingo number form string to int")
				}
				bingoNumbers = append(bingoNumbers, conv)
			}
			row++
			continue
		}

		if input == "" {
			row++
			continue
		}

		numbers := strings.Split(input, " ")
		for _, number := range numbers {
			if number == "" {
				continue
			}
			number = strings.TrimSpace(number)
			conv, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println(err)
				log.Fatal("Unable to convert bingo number form string to int")
			}
			allNumbers = append(allNumbers, conv)
		}

		row++
	}

	if len(allNumbers)%25 != 0 {
		log.Fatal("We have Rogue numbers in the imput")
	}

	// We have no read all the numbers into a buffer now we can create our boards
	cursor := 0
	i := 0
	for cursor+25 <= len(allNumbers) {
		fmt.Printf("Creating board %d \n", i+1)
		card := BingoCard{
			numbers: allNumbers[cursor : cursor+25],
			board:   [5][]Position{},
		}

		card.buildBoard()

		cards = append(cards, card)

		i++
		cursor += 25
	}

	return bingoNumbers, cards, nil
}
