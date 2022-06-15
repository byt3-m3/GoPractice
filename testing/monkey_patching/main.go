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
		fmt.Printf("Got 200 Response %v", resp)
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
