package httpgo

import (
	"errors"
	"io"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body io.Reader) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}
	return client.Do(req)
}
