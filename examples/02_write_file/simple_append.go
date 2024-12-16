package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Key aspects

// 1. Open file using os.OpenFile with append mode

func simpleAppendFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Return current local time
	t := time.Now()

	// Format time to RFC3339
	_, err = file.WriteString(t.Format(time.RFC3339))
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Println("File appended successfully")
}
