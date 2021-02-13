package httpgo

import (
	"fmt"
	"net/http"
)

// Mock structure provides a way to configure HTTP mocks using the combined
// request method, URL and request body.
type Mock struct {
	Method             string
	Url                string
	RequestBody        string
	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

// GetResponse provides a response based on a mock configuration
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}
	return &response, nil
}
