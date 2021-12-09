package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	x, y := 0, 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		command := strings.Split(strings.TrimSpace(line), " ")
		if len(command) != 2 {
			fmt.Println("len(command) != 2")
			continue
		}

		direction := command[0]
		delta, err := strconv.Atoi(command[1])
		if err != nil {
			fmt.Println("Could not convert string to int:", err)
			os.Exit(2)
		}

		switch direction {
		case "forward":
			x += delta
		case "up":
			y -= delta
		case "down":
			y += delta
		default:
			fmt.Println("Unknown direction:", direction)
		}
	}
	fmt.Printf("Horizontal: %d\nVertical: %d\n", x, y)
	fmt.Println("Multiplied:", x*y)
}

func part2(input string) {
	x, y, aim := 0, 0, 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		command := strings.Split(strings.TrimSpace(line), " ")
		if len(command) != 2 {
			fmt.Println("len(command) != 2")
			continue
		}

		direction := command[0]
		delta, err := strconv.Atoi(command[1])
		if err != nil {
			fmt.Println("Could not convert string to int:", err)
			os.Exit(2)
		}

		switch direction {
		case "forward":
			x += delta
			y += -aim * delta
		case "up":
			aim += delta
		case "down":
			aim -= delta
		default:
			fmt.Println("Unknown direction:", direction)
		}
	}
	fmt.Printf("Horizontal: %d\nVertical: %d\n", x, y)
	fmt.Println("Multiplied:", x*y)
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
