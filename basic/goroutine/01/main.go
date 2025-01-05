package main

import (
	"fmt"
	"time"
)

func printHello(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func calc(a int, b int, ch chan int) {
	// Do some calculation
	ch <- a + b
}

func main() {
	go printHello("Hello from goroutine1")
	go printHello("Hello from goroutine2")

	ch := make(chan int)
	go calc(1, 2, ch)
	result := <-ch

	fmt.Println("Hello from main")
	fmt.Println("Result:", result)
	time.Sleep(1 * time.Second)
}
