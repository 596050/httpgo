package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Errorf("%s is invalid", HeaderContentType)
	}
	if HeaderUserAgent != "User-Agent" {
		t.Errorf("%s is invalid", HeaderUserAgent)
	}
	if ContentTypeJSON != "application/json" {
		t.Errorf("%s is invalid", ContentTypeJSON)
	}
	if ContentTypeXML != "application/xml" {
		t.Errorf("%s is invalid", ContentTypeXML)
	}
	if ContentTypeOctetStream != "application/octet-stream" {
		t.Errorf("%s is invalid", ContentTypeOctetStream)
	}
}
