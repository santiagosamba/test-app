package utils

import "net/http"

type HTTPClientMock struct {
	Response *http.Response
	Err      error
}

func (c *HTTPClientMock) Get(url string) (*http.Response, error) {
	return c.Response, c.Err
}
