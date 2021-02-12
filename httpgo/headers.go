package httpgo

import "net/http"

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

// allows for both common and custom headers
func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// add common headers to the request
	for header, value := range c.builder.headers {
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
