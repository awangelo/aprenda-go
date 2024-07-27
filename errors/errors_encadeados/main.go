package main

import (
	"fmt"
	"os"
)

func main() {
	err := processFile()
	if err != nil {
		fmt.Println(err)
	}
}

func readFile() error {
	file, err := os.Open("file.txt")
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()
	return nil
}

func processFile() error {
	err := readFile()
	if err != nil {
		return fmt.Errorf("Cannot obtain file: %w", err)
	}
	return nil
}
