package main

import (
	"fmt"
)

func main() {
	filepath := "example.txt"

	fmt.Println("=== File Stats ===")
	readFileStats(filepath)

	fmt.Println("\n=== File Contents ===")
	readFile(filepath)

	fmt.Println("\n=== File with Search ===")
	readFileWithSearch(filepath, "Globetrotter")
}
