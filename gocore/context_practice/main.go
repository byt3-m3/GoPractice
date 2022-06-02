package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type (
	Person struct {
		FirstName string
		LastName  string
	}
)

func operation1(ctx context.Context, p Person) {
	fmt.Println(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println(p)
	dtime, status := ctx.Deadline()
	if time.Now().After(dtime) {
		log.Fatalln(status)
	}
	select {

	case <-ctx.Done():
		log.Println("Timeout Exceeded")
		log.Fatalln(ctx.Err())

	}
}

func operation2(ctx context.Context, p Person) {
	fmt.Println(p)
}

func main() {

	ctx := context.Background()

	ctx, _ = context.WithTimeout(ctx, time.Second*2)

	p := Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	operation1(ctx, p)
	operation2(ctx, p)
	select {

	case <-ctx.Done():
		log.Println("Timeout Exceeded")
		log.Fatalln(ctx.Err())

	}

}
