package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"testing"
)

const (
	TestUrl = "https://google.com"
)

type (
	MockHTTPClient struct {
		mock.Mock
	}
)

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	log.Println("Mock Hit")
	return &http.Response{StatusCode: 200}, nil
}

func TestDefaultHttpClientDo(t *testing.T) {
	mClient := MockHTTPClient{}
	t.Run("test when making request to google.com", func(t *testing.T) {
		dfBuilder := DataFetcherBuilder{}

		dataFetcher := dfBuilder.WithClient(&mClient).Build()

		mClient.On("Do", TestUrl)

		resp, err := dataFetcher.GetData(TestUrl)
		if err != nil {
			assert.NoError(t, err, "unexpected error occurred")
		}
		assert.Equal(t, resp.StatusCode, 200)

	})

	t.Run("test when making request to google.com", func(t *testing.T) {

		dataFetcher := NewDataFetcher(&DataFetcherConfig{}, WithClientOpt(&mClient))

		mClient.On("Do", TestUrl)

		resp, err := dataFetcher.GetData(TestUrl)
		if err != nil {
			assert.NoError(t, err, "unexpected error occurred")
		}
		assert.Equal(t, resp.StatusCode, 200)

	})

	t.Run("test when making request to google.com", func(t *testing.T) {
		var clientGenFunc = func() HTTPClient {
			return &mClient
		}
		dataFetcher := NewDataFetcherFromGenerator(clientGenFunc)

		mClient.On("Do", TestUrl)

		resp, err := dataFetcher.GetData(TestUrl)
		if err != nil {
			assert.NoError(t, err, "unexpected error occurred")
		}
		assert.Equal(t, resp.StatusCode, 200)

	})
}
