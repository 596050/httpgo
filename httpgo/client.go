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
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body Body, headers ...http.Header) (*Response, error)
	Put(url string, body Body, headers ...http.Header) (*Response, error)
	Patch(url string, body Body, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

// methods

func (c *httpClient) Get(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}
func (c *httpClient) Post(url string, body Body, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}
func (c *httpClient) Put(url string, body Body, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}
func (c *httpClient) Patch(url string, body Body, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), body)
}
func (c *httpClient) Delete(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}
func (c *httpClient) Options(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}

// New instantiates an httpClient
func New() Client {
	client := &httpClient{}
	return client
}
