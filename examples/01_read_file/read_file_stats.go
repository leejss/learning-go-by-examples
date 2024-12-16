package main

import (
	"fmt"
	"log"
	"os"
)

func readFileStats(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	stats, err := file.Stat()

	if err != nil {
		log.Fatalf("Error getting file stats: %v", err)
	}

	bytesSize := stats.Size()
	kbSize := bytesSize / 1024
	mbSize := kbSize / 1024

	fmt.Printf("File name: %s\n", stats.Name())
	fmt.Printf("File size: %d bytes (%d KB, %d MB)\n", bytesSize, kbSize, mbSize)
}
