package main

import "fmt"

func main() {

	fmt.Println("Simple Write")
	simpleWriteFile("test.txt", "Hello, World!!!")

	fmt.Println("Simple Append")
	simpleAppendFile("test.txt")

	fmt.Println("Using Writer")
	usingWriter("buffered_test.txt", "Hello, World!!!")
}
