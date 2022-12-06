package main

import (
	"log"
	"net/http"
)

type (
	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	DataFetcher interface {
		GetData(url string) (*http.Response, error)
	}
	DataFetcherConfig struct {
		client HTTPClient
	}
	DataFetcherConfigOpt func(config *DataFetcherConfig)

	defaultDetaFetcher struct {
		client HTTPClient
	}
)

func NewDataFetcherFromGenerator(clientGeneratorFunc func() HTTPClient) DataFetcher {
	client := clientGeneratorFunc()
	return &defaultDetaFetcher{client: client}
}

func NewDataFetcher(config *DataFetcherConfig, options ...DataFetcherConfigOpt) DataFetcher {
	for _, opt := range options {
		opt(config)
	}

	df := defaultDetaFetcher{client: config.client}
	return df
}
func WithClientOpt(client HTTPClient) DataFetcherConfigOpt {
	return func(config *DataFetcherConfig) {
		config.client = client
	}
}

type DataFetcherBuilder struct {
	client HTTPClient
}

func (b *DataFetcherBuilder) WithClient(client HTTPClient) *DataFetcherBuilder {
	b.client = client
	return b
}

func (b DataFetcherBuilder) Build() DataFetcher {
	if b.client == nil {
		b.client = &http.Client{}
	}

	df := defaultDetaFetcher{client: b.client}

	return df
}

func (c defaultDetaFetcher) GetData(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("unable to build request")
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println("unable to print request")
		return nil, err

	}

	if resp.StatusCode == 200 {
		log.Println("request successfully completed")

	}
	return resp, nil
}

func main() {

}
