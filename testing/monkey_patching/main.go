package main

import (
	"fmt"
	"log"
	"net/http"
)

var HTTPGetter = http.Get

func ExecuteReqeust() (*http.Response, error) {
	resp, err := HTTPGetter("google.com")
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Completed Business Logic")
	}

	if resp.StatusCode == http.StatusBadRequest {
		fmt.Println("Failed Business Logic")
	}
	return resp, nil

}

func main() {
	resp, err := ExecuteReqeust()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v", resp)
}
