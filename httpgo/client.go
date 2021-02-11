package httpgo

import (
	"net/http"
	"sync"
)

// Body type is generic where conversion to required type occurs internally
type Body interface{}

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

// Client is implemented by httpClient
type Client interface {
	// methods
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body Body) (*Response, error)
	Put(url string, headers http.Header, body Body) (*Response, error)
	Patch(url string, headers http.Header, body Body) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}

// methods
func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body Body) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}
func (c *httpClient) Put(url string, headers http.Header, body Body) (*Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}
func (c *httpClient) Patch(url string, headers http.Header, body Body) (*Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}
func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

// New instantiates an httpClient
func New() Client {
	client := &httpClient{}
	return client
}
