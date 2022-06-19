package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"testing"
)

type TestMock struct {
	mock.Mock
}

func (m *TestMock) Get(url string) (*http.Response, error) {
	// This Mocks using type assertions,
	args := m.Called(url)
	fmt.Println("Mock Invoked", url)
	return args.Get(0).(*http.Response), nil
}

func TestGetRequest(t *testing.T) {

	type (
		TestCase struct {
			name         string
			url          string
			mockResponse *http.Response
		}
	)

	testCases := []TestCase{
		{"Basic Google TestCase",
			"https://google.com",
			&http.Response{StatusCode: http.StatusOK},
		},
		{"Bad Domain TestCase",
			"https://baddomain.com",
			&http.Response{StatusCode: http.StatusBadRequest},
		},
	}

	for _, testCase := range testCases {
		log.Printf("Running %s", testCase.name)

		// Set Up TestEnv
		mockObj := new(TestMock)
		HTTPGetter = mockObj.Get // Monkey-Patches HttpGetter

		// SetUp Expectations
		mockObj.On("Get", testCase.url).Return(testCase.mockResponse)

		// Invocation
		resp, err := GetRequest(testCase.url)
		if err != nil {
			log.Println(err)

		}

		// Validation
		assert.Equal(t, resp, testCase.mockResponse)
		assert.Equal(t, 1, len(mockObj.Calls))
	}

}
