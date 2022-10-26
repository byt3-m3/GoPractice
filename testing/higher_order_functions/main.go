package main

import (
	"fmt"
	"log"
	"net/http"
)

type httpGetter func(url string) (resp *http.Response, err error)

func ExecuteRequest(getter httpGetter) {
	resp, err := getter("http://google.com")
	//resp, err := http.Get("google.com")
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Got 200 Response %v", resp)
	}

	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Got 400 Response %v", resp)
	}

}

func main() {
	ExecuteRequest(http.Get)
}
