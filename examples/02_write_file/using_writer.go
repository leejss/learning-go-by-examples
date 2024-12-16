package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Key aspects

// 1. Write operation using bufio.NewWriter

// 2. Create file using os.Create
// 3. And create writer using bufio.NewWriter
// 4. Write to buffer
// 5. Flush buffer to file

func usingWriter(filePath string, s string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	// Write to buffer
	for i := 0; i < 100; i++ {
		_, err := writer.WriteString(s)

		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	err = writer.Flush()

	if err != nil {
		log.Fatalf("Error flushing buffer: %v", err)
	}

	fmt.Println("File written successfully")
}
