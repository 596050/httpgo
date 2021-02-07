package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/596050/httpgo/httpgo"
)

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	time.Sleep(10 * time.Second)
// 	fmt.Fprint(w, "hello world!")
// }

// func serveRootHandler() {
// 	http.HandleFunc("/", rootHandler)
// 	http.ListenAndServe(":8080", nil)
// }
var (
	httpClient = getGithubClient()
)

func getGithubClient() httpgo.HTTPClient {
	client := httpgo.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)
	return client
}

func getUrls() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	resp, err := httpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func main() {
	getUrls()
}
