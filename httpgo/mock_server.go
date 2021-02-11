package httpgo

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enabled      bool
	serverMutext sync.Mutex
	mocks        map[string]*Mock
}

// allows
func StartMockServer() {
	// a single routine can enable at a time
	mockupServer.serverMutext.Lock()
	defer mockupServer.serverMutext.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	// a single routine can disable at a time
	mockupServer.serverMutext.Lock()
	defer mockupServer.serverMutext.Unlock()

	mockupServer.enabled = false
}

func FlushMocks() {
	// a single routine can flush at a time
	mockupServer.serverMutext.Lock()
	defer mockupServer.serverMutext.Unlock()

	mockupServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	// a single routine can add at a time
	mockupServer.serverMutext.Lock()
	defer mockupServer.serverMutext.Unlock()

	// var headersString strings.Builder
	// for key, val := range mock.Header {
	// 	headersString.WriteString(key + strings.Join(val, ","))
	// }
	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	if !m.enabled {
		return nil
	}

	if mock := m.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}
	return &Mock{
		Error: errors.New(fmt.Sprintf("no mock matching %s from '%s' with given body", method, url)),
	}

}
