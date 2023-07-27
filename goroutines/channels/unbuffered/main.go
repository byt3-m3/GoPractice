package main

import (
	"fmt"
	"time"
)

func main() {

	//var unBufChan = make(chan string, 5)
	var unBufChan = make(chan string)

	defer close(unBufChan)

	go func(ch chan string) {
		time.Sleep(3 * time.Second)

		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)

	}(unBufChan)

	unBufChan <- "test"
	unBufChan <- "test-2"
	unBufChan <- "test-3"
	unBufChan <- "test-4"
	time.Sleep(5 * time.Second)

}
