package examples

import (
	"net/http"
	"time"

	"github.com/596050/httpgo/httpgo"
)

var (
	httpClient = getHTTPClient()
)

func getHTTPClient() httpgo.Client {
	currentClient := http.Client{}
	client := httpgo.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetHttpClient(&currentClient).
		Build()
	return client
}
