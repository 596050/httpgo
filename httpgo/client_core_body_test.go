package httpgo

import (
	"testing"
)

func TestGetRequestBody(t *testing.T) {
	// Initialisation
	client := httpClient{}
	t.Run("WithNil", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)
		// Validation
		if err != nil {
			t.Error("no error expected")
		}
		if body != nil {
			t.Error("no body expected")
		}
	})
	t.Run("WithJSON", func(t *testing.T) {
		// Execution
		requestBody := []string{"1", "2"}
		body, err := client.getRequestBody("application/json", requestBody)
		// Validation
		if err != nil {
			t.Error("no error expected")
		}
		if string(body) != `["1","2"]` {
			t.Errorf("invalid body: %s", string(body))
		}
	})
	t.Run("WithXML", func(t *testing.T) {
		// Execution
		requestBody := []string{"1", "2"}
		body, err := client.getRequestBody("application/xml", requestBody)
		// Validation
		if err != nil {
			t.Error("no error expected")
		}
		if string(body) != `<string>1</string><string>2</string>` {
			t.Errorf("invalid body: %s", string(body))
		}
	})
	t.Run("WithDefault", func(t *testing.T) {
		// Execution
		requestBody := []string{"1", "2"}
		body, err := client.getRequestBody("application/json", requestBody)
		// Validation
		if err != nil {
			t.Error("no error expected")
		}
		if string(body) != `["1","2"]` {
			t.Errorf("invalid body: %s", string(body))
		}
	})
}
