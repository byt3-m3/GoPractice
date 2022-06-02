package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg1 sync.WaitGroup
	ctx := context.Background()

	wg1.Add(3)

	chan1 := make(chan string, 5)

	go makePhoneCall(ctx, chan1, &wg1)
	go answerPhoneCall(ctx, chan1, &wg1)

	select {

	case <-ctx.Done():
		log.Println("Handled Cancelled context")
		log.Fatalln(ctx.Err())

	}
	//
	//case data1 := <-chan1:
	//	fmt.Println("Received Data1")
	//	fmt.Println(data1)
	//
	//case data2 := <-chan2:
	//	fmt.Println("Received Datas2")
	//	fmt.Println(data2)
	//
	//case <-chan1:
	//	data := <-chan1
	//	fmt.Println("Default Case")
	//	fmt.Println(data)
	//}
	wg1.Wait()
	fmt.Println("Finished")

}

func makePhoneCall(ctx context.Context, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	ctx, cancelCtx := context.WithCancel(ctx)
	count := 0
	for count <= 2 {
		fmt.Println(count)
		ch <- "Dialing Mom"
		fmt.Println("Made Call")
		//time.Sleep(2 * time.Second)
		count += 1

	}

	cancelCtx()

	fmt.Println(count)

}

func answerPhoneCall(ctx context.Context, ch <-chan string, wg *sync.WaitGroup) {
	ctx, cancelCtx := context.WithCancel(ctx)
	fmt.Println(ctx)
	for {
		select {
		case msg := <-ch:
			fmt.Println(len(msg))
			if len(msg) == 0 {
				cancelCtx()
				wg.Done()

			}
			fmt.Println(fmt.Sprintf("New Message: %s", msg))
			time.Sleep(3 * time.Second)

			//case <-ctx.Done():
			//	log.Println("Handled Failed Context")
			//	log.Fatalln(ctx.Err())

			//default:
			//	fmt.Println("Awaiting")

		}
	}

}
