package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Store board in board struct
	input, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	inputParts := strings.Split(input, "\n\n")

	numbersDrawn := strings.Split(inputParts[0], ",")

	boards := []*Board{}
	for i, boardData := range inputParts {
		if i == 0 || boardData == "" {
			continue
		}
		boards = append(boards, parseBoard(boardData))
	}

	fmt.Println("Part 1:")
	part1(numbersDrawn, boards)

	fmt.Println("\n\nPart 2:")
	part2(numbersDrawn, boards)
}

func part1(numbersDrawn []string, boards []*Board) {
	for _, numberStr := range numbersDrawn {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Printf("Error converting drawn number %s: %v", numberStr, err)
			os.Exit(2)
		}
		for _, board := range boards {
			win := board.drawn(number)
			if win {
				fmt.Println("Board won: ", board)
				score := board.getUnmarkedSum() * number
				fmt.Println("Score: ", score)
				return
			}
		}
	}
}

func part2(numbersDrawn []string, boards []*Board) {
	var lastWinBoard *Board
	var lastWinScore int
	for _, numberStr := range numbersDrawn {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Printf("Error converting drawn number %s: %v", numberStr, err)
			os.Exit(2)
		}
		for _, board := range boards {
			if board.won {
				continue
			}
			win := board.drawn(number)
			if win {
				board.won = true
				lastWinBoard = board
				lastWinScore = board.getUnmarkedSum() * number
			}
		}
	}
	fmt.Println("Last board won: ", lastWinBoard)
	fmt.Println("Last score: ", lastWinScore)
}

type Board struct {
	numbers [5][5]int
	marked  [5][5]bool
	won     bool
}

func parseBoard(input string) *Board {
	board := &Board{
		numbers: [5][5]int{},
		marked:  [5][5]bool{},
		won:     false,
	}
	rows := strings.Split(input, "\n")
	for y, row := range rows {
		if y == 5 {
			continue
		}
		x := 0
		numbers := make([]string, 5)
		for _, c := range row {
			if string(c) == " " {
				if len(numbers[x]) > 0 {
					number, err := strconv.ParseInt(numbers[x], 10, 32)
					if err != nil {
						fmt.Println("Error parsing int:", err, "\nNumber:", numbers[x])
					}
					board.numbers[y][x] = int(number)
					x++
				}
				continue
			}
			numbers[x] += string(c)
		}
		if x < 5 {
			number, err := strconv.ParseInt(numbers[x], 10, 32)
			if err != nil {
				fmt.Println("Error parsing int:", err, "\nNumber:", numbers[x])
			}
			board.numbers[y][x] = int(number)
		}
	}
	return board
}

func (b *Board) drawn(number int) bool {
	for y, col := range b.numbers {
		for x, value := range col {
			if value == number {
				b.marked[y][x] = true
				return b.checkRow(y) || b.checkCol(x)
			}
		}
	}
	return false
}

func (b *Board) checkRow(y int) bool {
	for i := 0; i < len(b.numbers); i++ {
		if !b.marked[y][i] {
			return false
		}
	}
	return true
}

func (b *Board) checkCol(x int) bool {
	for i := 0; i < len(b.numbers[x]); i++ {
		if !b.marked[i][x] {
			return false
		}
	}
	return true
}

func (b *Board) getUnmarkedSum() int {
	sum := 0
	for y, col := range b.marked {
		for x, m := range col {
			if !m {
				sum += b.numbers[y][x]
			}
		}
	}
	return sum
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
