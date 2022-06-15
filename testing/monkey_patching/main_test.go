package main

import (
	"fmt"
	"net/http"
	"testing"
)

func mockHttpGetterSuccess(url string) (resp *http.Response, err error) {
	fmt.Println("Invoking Mock")
	// Do Custom Things....
	return &http.Response{StatusCode: http.StatusOK}, err
}

func mockHttpGetterBadRequest(url string) (resp *http.Response, err error) {
	fmt.Println("Invoking Mock")
	// Do Custom Things....
	return &http.Response{StatusCode: http.StatusBadRequest}, err
}

func TestBasicTest(t *testing.T) {
	// Overshadows the package variable with Mock HTTP Getter.
	HTTPGetter = mockHttpGetterSuccess

	resp, err := ExecuteReqeust()
	if err != nil {
		t.Fail()
	}

	if resp.StatusCode != http.StatusOK {
		t.Fail()
	}
}

func TestBasicTestStatusBadRequest(t *testing.T) {
	HTTPGetter = mockHttpGetterBadRequest

	resp, err := ExecuteReqeust()
	if err != nil {
		t.Fail()
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fail()
	}

}
