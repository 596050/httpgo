package httpgo

import (
	"errors"
	"io"
	"net/http"
)

// do is a private method taking which makes the http request
func (c *httpClient) do(method string, url string, headers http.Header, body io.Reader) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	fullHeaders := c.getRequestHeaders(headers)
	req.Header = fullHeaders

	return client.Do(req)
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// add common headers to the request
	for header, value := range c.Headers {
		// headers should have value of length one
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	// add custom headers to the request
	for header, value := range requestHeaders {
		// headers should have value of length one
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}
