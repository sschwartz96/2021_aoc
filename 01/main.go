package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input, err := readFile("input1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	values := strings.Split(input, "\n")
	prevVal := 0
	counter := 0
	for i, val := range values {
		if val == "" {
			continue
		}
		intVal, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			fmt.Println("Error parsing string value:", err)
			os.Exit(2)
		}
		if i == 0 {
			prevVal = intVal
			continue
		}

		if intVal > prevVal {
			counter++
		}
		prevVal = intVal
	}

	fmt.Println("Increments:", counter)
}

func part2() {
	input, err := readFile("input1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	values := strings.Split(input, "\n")
	prevVals := make([]int, 4)
	counter := 0
	for i, val := range values {
		if val == "" {
			continue
		}

		intVal, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			fmt.Println("Error parsing string value:", err)
			os.Exit(2)
		}

		modVal := i % 4
		switch modVal {
		case 0:
			prevVals[0] = intVal

			prevVals[2] += intVal
			prevVals[3] += intVal
		case 1:
			prevVals[1] = intVal

			prevVals[3] += intVal
			prevVals[0] += intVal
		case 2:
			prevVals[2] = intVal

			prevVals[0] += intVal
			prevVals[1] += intVal
		case 3:
			prevVals[3] = intVal

			prevVals[1] += intVal
			prevVals[2] += intVal
		}

		if i >= 3 && prevVals[(modVal+2)%4] > prevVals[(modVal+1)%4] {
			counter++
			prevVals[(modVal+1)%4] = 0
		}
	}

	fmt.Println("Increments:", counter)
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
