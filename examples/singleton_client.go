package examples

import (
	"time"

	"github.com/596050/httpgo/httpgo"
)

var (
	httpClient = getHTTPClient()
)

func getHTTPClient() httpgo.Client {
	client := httpgo.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}
