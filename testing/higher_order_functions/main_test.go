package main

import (
	"fmt"
	"net/http"
	"testing"
)

func mockHttpGetter(url string) (resp *http.Response, err error) {
	fmt.Println("Invoking Mock")
	return &http.Response{StatusCode: http.StatusOK}, err
}

func TestBasicTest(t *testing.T) {

	ExecuteRequest(mockHttpGetter)
}
