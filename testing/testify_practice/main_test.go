package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"testing"
)

type TestMock struct {
	mock.Mock
}

func (m *TestMock) Get(url string) (*http.Response, error) {
	//m.Called(url)
	args := m.Called(url)
	fmt.Println("Mock Invoked", url)
	m.Called()
	return args.Get(0).(*http.Response), args.Error(0)
}

func TestGetRequest(t *testing.T) {
	mockObj := new(TestMock)
	//HTTPGetter = mockObj.Get
	mockObj.On("Get", "https://google.com").Return(&http.Response{StatusCode: http.StatusOK})
	resp, err := mockObj.Get("https://google.com")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(resp)
	fmt.Println(mockObj.Calls)
}
