package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/596050/httpgo/httpgo"
)

func getClientGithub() httpgo.HTTPClient {
	client := httpgo.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)
	return client
}

var (
	httpClientGithub = getClientGithub()
)

// get
func getGithubUrls() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	resp, err := httpClientGithub.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

// post
type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user user) {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	resp, err := httpClientGithub.Post("https://api.github.com", headers, user)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func main() {
	getGithubUrls()
}
