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
}

// NewBuilder instantiates an httpClient
func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
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
