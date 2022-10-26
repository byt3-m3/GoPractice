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

func mockHttpGetterFailure(url string) (resp *http.Response, err error) {
	fmt.Println("Invoking Mock")
	return &http.Response{StatusCode: http.StatusBadRequest}, err
}

func TestBasicTestSuccess(t *testing.T) {

	ExecuteRequest(mockHttpGetter)

}

func TestBasicTestFailure(t *testing.T) {

	ExecuteRequest(mockHttpGetterFailure)

}
