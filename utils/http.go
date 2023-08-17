package utils

import "net/http"

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type defaultHTTPClient struct{}

func NewHTTPClient() HTTPClient {
	return &defaultHTTPClient{}
}

func (c *defaultHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
