package httpgo

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialisation
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeadersTestCases := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion",
	}
	customHeadersTestCases := map[string]string{
		"X-Request-Id": "ABC-123",
	}
	// add common headers
	for key, value := range commonHeadersTestCases {
		commonHeaders.Set(key, value)
	}
	client.Headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	// add custom headers
	for key, value := range customHeadersTestCases {
		requestHeaders.Set(key, value)
	}
	finalHeaders := client.getRequestHeaders(requestHeaders)
	// Validation
	if len(finalHeaders) != 3 {
		t.Error("expected 3 headers")
	}

	// has correct values for common and custom headers
	for key, value := range commonHeadersTestCases {
		if finalHeaders.Get(key) != value {
			t.Errorf("unexpected request header\n key: %s\n value: %s", key, value)
		}
	}

	for key, value := range customHeadersTestCases {
		if finalHeaders.Get(key) != value {
			t.Errorf("unexpected request header\n key: %s\n value: %s", key, value)
		}
	}
}
