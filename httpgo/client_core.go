package httpgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

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

func (c *httpClient) getHTTPClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}

		c.client = &http.Client{
			Timeout: c.builder.connectionTimeout + c.builder.responseTimeout,
			// should allow for configuring according to traffic patterns
			Transport: &http.Transport{
				// maximum idle (keep-alive) connections to keep per-host
				MaxIdleConnsPerHost: c.builder.maxIdleConnections,
				// amount of time to wait for a server's response headers after fully writing the request (including its body, if any). This time does not include the time to read the response body.
				ResponseHeaderTimeout: c.builder.responseTimeout,
				// maximum amount of time to wait for a given connection
				DialContext: (&net.Dialer{
					Timeout: c.builder.connectionTimeout,
				}).DialContext,
			},
		}
	})

	return c.client
}

// do is a private method which makes the http request
func (c *httpClient) do(method string, url string, headers http.Header, body Body) (*Response, error) {
	// handle headers
	fullHeaders := c.getRequestHeaders(headers)
	// handle body
	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	if mock := mockupServer.getMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}

	// create request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	client := c.getHTTPClient()

	// set headers on request
	req.Header = fullHeaders
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		header:     response.Header,
		body:       responseBody,
	}

	return &finalResponse, nil
}
