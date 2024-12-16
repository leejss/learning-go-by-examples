package main

import (
	"fmt"
	"log"
	"os"
)

// Key aspects

// 1. Create file using os.Create
// 2. Write string to file using file.WriteString
// 3. Check error
// 4. Close file using defer

func simpleWriteFile(filePath string, s string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	defer file.Close()

	_, err = file.WriteString(s)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Println("File written successfully")
}
