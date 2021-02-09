package httpgo

import (
	"net/http"
)

// Body type is generic where conversion to required type occurs internally
type Body interface{}

type httpClient struct {
	Headers http.Header
}

// New instantiates an httpClient and returns it
func New() HTTPClient {
	client := &httpClient{}
	return client
}

// HTTPClient is implemented by httpClient
type HTTPClient interface {
	SetHeaders(headers http.Header)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body Body) (*http.Response, error)
	Put(url string, headers http.Header, body Body) (*http.Response, error)
	Patch(url string, headers http.Header, body Body) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body Body) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}
func (c *httpClient) Put(url string, headers http.Header, body Body) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}
func (c *httpClient) Patch(url string, headers http.Header, body Body) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}
func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
