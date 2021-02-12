package httpgo

import (
	"net/http"
	"time"
)

// implements ClientBuilder interface
type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	headers            http.Header
	baseUrl            string
	client             *http.Client
}

// ClientBuilder uses builder pattern for configuration
type ClientBuilder interface {
	// headers
	SetHeaders(headers http.Header) ClientBuilder
	// timeouts
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(numConnections int) ClientBuilder

	// create http client
	Build() Client
	SetHttpClient(c *http.Client) ClientBuilder
}

// NewBuilder instantiates an httpClient
func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

// headers
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

// timeouts
func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(numConnections int) ClientBuilder {
	c.maxIdleConnections = numConnections
	return c
}

// create http client
func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}
