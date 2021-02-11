package httpgo

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status     string
	statusCode int
	header     http.Header
	body       []byte
}

func (r *Response) Status() string {
	return r.status
}
func (r *Response) StatusCode() int {
	return r.statusCode
}
func (r *Response) Header() http.Header {
	return r.header
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) UnmarshalJSON(target interface{}) error {
	return json.Unmarshal(r.body, target)
}
