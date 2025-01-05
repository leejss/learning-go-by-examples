package main

import (
	"fmt"
	"sync"
	"time"
)

// sync package를 이용하여 goroutine끼리 협업하기
// waitgroup 사용하여 모든 goroutine이 완료될 때 까지 기다리기. wg.wait()

func worker(id int, wg *sync.WaitGroup) {
	// Mark done
	defer wg.Done()

	// Do work
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 모의 작업(1초 대기)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// waitgroup 생성
	var wg sync.WaitGroup

	// 3개의 goroutine를 생성
	for i := 1; i <= 3; i++ {
		// count up
		wg.Add(1)

		// goroutine 생성
		go worker(i, &wg)
	}

	// main goroutine가 waitgroup의 Done()를 불러가고, main goroutine가 끝나기를 기다리기
	wg.Wait()
}
