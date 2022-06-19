package main

import (
	"log"
	"net/http"
)

type HttpGetter interface {
	GetRequest(url string) (*http.Response, error)
}

var HTTPGetter = http.Get

func GetRequest(url string) (*http.Response, error) {
	resp, err := HTTPGetter(url)
	if err != nil {
		log.Fatal(err)

	}
	return resp, nil
}

func main() {

}
