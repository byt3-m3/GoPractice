package main

import (
	"fmt"
	"log"
	"net/http"
)

type httpGetter func(url string) (resp *http.Response, err error)

var HTTPGetter = http.Get

func ExecuteReqeust() {
	resp, err := HTTPGetter("google.com")
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Got 200 Response %v", resp)
	}

}

func main() {
	ExecuteReqeust()
}
