package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}
