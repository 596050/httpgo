package httpgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

// allows for both common and custom headers
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

// marshals body to encoding based on the content type header
func (c *httpClient) getRequestBody(contentType string, body Body) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	// custom content type management
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

// do is a private method which makes the http request
func (c *httpClient) do(method string, url string, headers http.Header, body Body) (*http.Response, error) {
	client := http.Client{}
	// handle headers
	fullHeaders := c.getRequestHeaders(headers)
	// handle body
	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}
	// create request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	// set headers on request
	req.Header = fullHeaders
	return client.Do(req)
}
