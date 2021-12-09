package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input, err := readFile("input1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

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
