package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFileWithSearch(filepath string, search string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains the search string
		if strings.Contains(line, search) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
