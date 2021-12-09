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
		fmt.Println("Could not read file:", err)
		os.Exit(1)
	}
	// redo part 1 as I removed it on accident
	part2(input)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	bitLength := len(lines[0])
	fmt.Println("Bit length:", bitLength)

	oxygenGen := make([]string, len(lines))
	copy(oxygenGen, lines)

	bitPosition := 0
	for len(oxygenGen) != 1 {
		oneBitCount, zeroBitCount := 0, 0
		for _, line := range oxygenGen {
			if line == "" {
				continue
			}
			if string(line[bitPosition]) == "1" {
				oneBitCount++
			}
			if string(line[bitPosition]) == "0" {
				zeroBitCount++
			}

		}
		if oneBitCount >= zeroBitCount {
			oxygenGen = getValuesPerBitColumn(oxygenGen, bitPosition, "1")
		} else {
			oxygenGen = getValuesPerBitColumn(oxygenGen, bitPosition, "0")
		}
		bitPosition++
	}

	cO2Scrub := make([]string, len(lines))
	copy(cO2Scrub, lines)

	bitPosition = 0
	for len(cO2Scrub) != 1 {
		oneBitCount, zeroBitCount := 0, 0
		for _, line := range cO2Scrub {
			if line == "" {
				continue
			}
			if string(line[bitPosition]) == "1" {
				oneBitCount++
			}
			if string(line[bitPosition]) == "0" {
				zeroBitCount++
			}

		}
		if zeroBitCount <= oneBitCount {
			cO2Scrub = getValuesPerBitColumn(cO2Scrub, bitPosition, "0")
		} else {
			cO2Scrub = getValuesPerBitColumn(cO2Scrub, bitPosition, "1")
		}
		bitPosition++
	}

	fmt.Println("oxygen:", oxygenGen)
	fmt.Println("cO2:", cO2Scrub)

	oxygen, _ := strconv.ParseInt(oxygenGen[0], 2, 64)
	cO2, _ := strconv.ParseInt(cO2Scrub[0], 2, 64)

	fmt.Printf("oxygen:\t%d\n", oxygen)
	fmt.Printf("cO2:   \t%d\n", cO2)
	fmt.Println("Life Support:", oxygen*cO2)
}

func getValuesPerBitColumn(values []string, colIndex int, bitValue string) []string {
	newValues := []string{}
	for _, value := range values {
		if value == "" {
			continue
		}
		if string(value[colIndex]) == bitValue {
			newValues = append(newValues, value)
		}
	}
	return newValues
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
