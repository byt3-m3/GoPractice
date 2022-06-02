package main

import (
	"fmt"
	"sync"
	"time"
)

func startWorker(workerID int) {
	fmt.Printf("Working.. Worker=%d\n", workerID)
	time.Sleep(3 * time.Second)
	fmt.Printf("Completed Worker=%d\n", workerID)

}

func main() {
	wg := sync.WaitGroup{}

	workerCount := 10
	wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go func(number int) {
			startWorker(number)
			wg.Done()
		}(i)
	}

	fmt.Println("Starting")

	wg.Wait()
	//time.Sleep(6 * time.Second)

}
